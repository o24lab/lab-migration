package migration

import (
	"domain/balance"
	"domain/betting"
	"domain/user"
	"domain/user/basicinfo"
	"domain/user/kyc"
	"domain/user/username"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"github.com/svc0a/worker"
	"go.mongodb.org/mongo-driver/bson"
	"migration/pkg"
	"migration/pkg/clientz"
	"migration/pkg/model"
	"migration/pkg/oldmain"
	"migration/pkg/oldsharding"
	"pkg/ddd"
	"pkg/errorx"
	"pkg/types"
	"time"
)

const workerNum = 1000

var (
	userWallets = map[int32]*model.WinUserWallet{}
	kycs        = map[int32][]model.WinUserKyc{}
	deposits    = map[int32]map[int64]*model.WinCoinDepositRecord{}
	withdraws   = map[int32]map[int64]*model.WinCoinWithdrawalRecord{}
)

func Migrate() errorx.Error {
	err := Init()
	if err != nil {
		logrus.Error(err)
		return err
	}
	data, err := func() ([]*model.WinUser, errorx.Error) {
		var err error
		db := oldmain.Use(clientz.MainDB()).WinUser
		data, err := db.Where(db.Mobile.Neq("9993334444")).Find()
		if err != nil {
			return nil, errorx.WithStack(err)
		}
		if len(data) == 0 {
			return nil, errorx.New("no record")
		}
		return data, nil
	}()
	if err != nil {
		logrus.Error(err)
		return err
	}
	wp := worker.New[model.WinUser](workerNum, worker.WithErrHandler[model.WinUser](func(err error) {
		//logrus.Error(err)
	}), worker.WithChanSize[model.WinUser](int64(len(data))))
	for _, item := range data {
		wp.Submit(*item)
	}
	wp.Start(func(winUser1 model.WinUser) error {
		u, err := User(&winUser1)
		if err != nil {
			logrus.Error(errorx.WithStack(err))
			return nil
		}
		go func() {
			err1 := Kyc(u)
			if err1 != nil {
				logrus.Error(errorx.WithStack(err1))
			}
		}()
		l := logrus.WithField("new user", u.ID).WithField("old user", u.User.ID)
		winCoinLogs, err := func() ([]*model.WinCoinLog, errorx.Error) {
			var err1 error
			db := oldsharding.Use(clientz.ShardingDB()).WinCoinLog
			q := db.Table(fmt.Sprintf("%s_%d", db.TableName(), pkg.CalculateShardingKey(int64(winUser1.ID))))
			winCoinLogs, err1 := q.Where(q.UID.Eq(winUser1.ID)).Order(q.ID.Asc()).Find()
			if err1 != nil {
				return nil, errorx.WithStack(err1)
			}
			return winCoinLogs, nil
		}()
		if err != nil {
			l.Error(err)
			return err
		}
		winBetSlips, err := func() (map[int64]*model.WinBetslips, errorx.Error) {
			winBetSlips := map[int64]*model.WinBetslips{}
			var err1 error
			db := oldsharding.Use(clientz.ShardingDB()).WinBetslips
			q := db.Table(fmt.Sprintf("%s_%d", db.TableName(), pkg.CalculateShardingKey(int64(winUser1.ID))))
			winBetSlips1, err1 := q.Where(q.XbUID.Eq(winUser1.ID)).Order(q.ID.Asc()).Find()
			if err1 != nil {
				return nil, errorx.WithStack(err1)
			}
			for _, betSlip := range winBetSlips1 {
				_, ok := winBetSlips[betSlip.ID]
				if ok {
					continue
				}
				winBetSlips[betSlip.ID] = betSlip
			}
			return winBetSlips, nil
		}()
		if err != nil {
			l.Error(err)
			return err
		}
		for _, winCoinLog := range winCoinLogs {
			coinLog1 := NewWinCoinLog(u, winCoinLog)
			err1 := coinLog1.Validate()
			if err1 != nil {
				l.Error(errorx.WithStack(err1))
			}
			//1-存款 2-提款 3-投注 4-派彩 5-返水 6-佣金 7-活动(奖励) 8-系统调账 9-退款 10-佣金钱包转主账户余额 11-小费" json:"category
			switch winCoinLog.Category {
			case 1:
				d := func() *model.WinCoinDepositRecord {
					d1, ok := deposits[winCoinLog.UID]
					if !ok {
						return nil
					}
					d2, ok := d1[winCoinLog.ReferID]
					if !ok {
						return nil
					}
					return d2
				}()
				if d != nil {
					err2 := coinLog1.WithDeposit(d).Deposit()
					if err2 != nil {
						l.Error(errorx.WithStack(err2))
					}
				}
			case 2:
				d := func() *model.WinCoinWithdrawalRecord {
					d1, ok := withdraws[winCoinLog.UID]
					if !ok {
						return nil
					}
					d2, ok := d1[winCoinLog.ReferID]
					if !ok {
						return nil
					}
					return d2
				}()
				if d != nil {
					err2 := coinLog1.WithWithdrawal(d).Withdraw()
					if err2 != nil {
						l.Error(errorx.WithStack(err2))
					}
				}
			case 3:
				d := func() *WinBetSlips {
					d1, ok := winBetSlips[winCoinLog.ReferID]
					if !ok {
						return nil
					}
					slip, err := getWinBetSlip(d1)
					if err != nil {
						return nil
					}
					return slip
				}()
				if d != nil {
					err2 := coinLog1.WithBetSlip(d).PlaceBet()
					if err2 != nil {
						l.Error(errorx.WithStack(err2))
					}
				}
			case 4:
				d := func() *WinBetSlips {
					d1, ok := winBetSlips[winCoinLog.ReferID]
					if !ok {
						return nil
					}
					slip, err := getWinBetSlip(d1)
					if err != nil {
						return nil
					}
					return slip
				}()
				if d != nil {
					err2 := coinLog1.WithBetSlip(d).Payout()
					if err2 != nil {
						l.Error(errorx.WithStack(err2))
					}
				}
			case 5, 6, 7, 8, 10, 11:
				err2 := coinLog1.Adjust()
				if err2 != nil {
					l.Error(errorx.WithStack(err2))
				}
			case 9:
				d := func() *WinBetSlips {
					d1, ok := winBetSlips[winCoinLog.ReferID]
					if !ok {
						return nil
					}
					slip, err2 := getWinBetSlip(d1)
					if err2 != nil {
						return nil
					}
					return slip
				}()
				if d != nil {
					err2 := coinLog1.WithBetSlip(d).Cancel()
					if err2 != nil {
						l.Error(errorx.WithStack(err2))
					}
				}
			default:
				err2 := coinLog1.Adjust()
				if err2 != nil {
					l.Error(errorx.WithStack(err2))
				}
			}
		}
		{
			b, err1 := balance.GetUserCurrentBalanceService.Request(newDDDCtx(u.ID), winUser1.ID)
			if err1 != nil {
				l.Error(errorx.WithStack(err1))
				return err1
			}
			wallet, ok := userWallets[u.User.ID]
			if ok {
				dif := decimal.NewFromFloat(wallet.Coin).Sub(b.Available)
				if !dif.IsZero() {
					_, err2 := balance.RecordService.Request(newDDDCtx(u.ID), balance.RecordCommand{
						Action: balance.ActionAdjust,
						Amount: dif,
						Force:  true,
					})
					if err2 != nil {
						l.Error(errorx.WithStack(err2))
					}
				}
			}
		}

		return nil
	})
	wp.Stop()
	return nil
}

func User(m *model.WinUser) (*WinUser, errorx.Error) {
	{
		winUser, err := userCollection.GetOne(bson.M{"user.id": m.ID})
		if err != nil && !errorx.ErrDataNotFound.Is(err) {
			return nil, err
		}
		if winUser != nil {
			return winUser, nil
		}
	}
	{
		user1, err := userQueryCollection.GetOne(bson.M{"$or": []bson.M{
			{"number": m.Mobile},
			{"email": m.Email},
			{"username": m.Username},
		}})
		if err != nil && !errorx.ErrDataNotFound.Is(err) {
			return nil, err
		}
		if user1 != nil {
			u := WinUser{
				Entity: types.Entity{
					ID: user1.ID,
				},
				User: *m,
			}
			err = userCollection.Create(&u)
			if err != nil {
				return nil, err
			}
			return &u, nil
		}
	}
	cmd := user.CreateUserCommand{}
	if m.Mobile != "" {
		cmd.PhoneNumber = &types.PhoneNumber{
			CountryCode: "63",
			Number:      m.Mobile,
		}
	}
	if m.Email != "" {
		cmd.Email = func() *types.Email {
			email := types.Email(m.Email)
			return &email
		}()
	}
	if m.Username != "" {
		cmd.Username = func() *username.Username {
			username1 := username.Username(m.Username)
			return &username1
		}()
	}
	cmd.BasicInfo = &basicinfo.UpdateCommand{}
	if m.Avatar != "" {
		cmd.BasicInfo.AvatarUrl = func() *basicinfo.AvatarUrl {
			n := basicinfo.AvatarUrl(m.Avatar)
			return &n
		}()
	}
	if m.Username != "" {
		cmd.BasicInfo.Nickname = func() *basicinfo.Nickname {
			n := basicinfo.Nickname(m.Username)
			return &n
		}()
	}
	cmd.BasicInfo.Gender = func() *types.Gender { g := types.GenderUnKnow; return &g }()
	id, err := user.CreateUserService.Request(newDDDCtx(types.NewID()), cmd)
	if err != nil {
		return nil, err
	}
	u := WinUser{
		Entity: types.Entity{
			ID: *id,
		},
		User: *m,
	}
	err = userCollection.Create(&u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func Kyc(user *WinUser) errorx.Error {
	_, ok := kycs[user.User.ID]
	if !ok {
		return nil
	}
	if len(kycs[user.User.ID]) == 0 {
		return nil
	}
	wp := worker.New[model.WinUserKyc](workerNum, worker.WithErrHandler[model.WinUserKyc](func(err error) {
		//logrus.Error(err)
	}))
	wp.Start(func(data model.WinUserKyc) error {
		l := logrus.WithField("new user", user.ID).WithField("old user", user.User.ID)
		cmd := kyc.ApplyCommand{}
		cmd.IDType = func() *kyc.IDType { i := kyc.IDTypeIDCard; return &i }()
		{
			//1-身份证 2-护照 3-驾照
			switch data.IDType {
			case 1:
				cmd.IDType = func() *kyc.IDType { i := kyc.IDTypeIDCard; return &i }()
			case 2:
				cmd.IDType = func() *kyc.IDType { i := kyc.IDTypePassport; return &i }()
			case 3:
				cmd.IDType = func() *kyc.IDType { i := kyc.IDTypeDrivingLicense; return &i }()
			default:
				err := errorx.New(fmt.Sprintf("invalid id type, data=%v", data.IDType))
				l.Error(errorx.WithStack(err))
			}
			cmd.IDNumber = &data.IDNumber
			//RealName *types.RealName `json:"realName,omitempty"`
			cmd.RealName = &types.RealName{}
			if data.FirstName != "" {
				cmd.RealName.FirstName = data.FirstName
			}
			if data.LastName != "" {
				cmd.RealName.LastName = data.LastName
			}
			if data.MiddleName != "" {
				cmd.RealName.MiddleName = data.MiddleName
			}
			//：1.Male 0.Female 2.Prefer not to say"
			//Gender   *types.Gender   `json:"gender,omitempty"`
			cmd.Gender = func() *types.Gender { gender := types.GenderUnKnow; return &gender }()
			switch data.Gender {
			case 1:
				cmd.Gender = func() *types.Gender { gender := types.GenderMale; return &gender }()
			case 0:
				cmd.Gender = func() *types.Gender { gender := types.GenderFemale; return &gender }()
			case 2:
				cmd.Gender = func() *types.Gender { gender := types.GenderUnKnow; return &gender }()
			default:
				err := errorx.New(fmt.Sprintf("invalid gender=%v", data.Gender))
				l.Error(errorx.WithStack(err))
			}
			//Birthday *Birthday       `json:"birthday,omitempty"`
			cmd.Birthday = func() *kyc.Birthday {
				t := time.Unix(int64(data.Birthday), 0)
				return &kyc.Birthday{
					Year:  t.Year(),
					Month: int(t.Month()),
					Day:   t.Day(),
				}
			}()
			//Medias   *[]Media        `json:"medias,omitempty"`
			cmd.Medias = func() *[]kyc.Media {
				r := []kyc.Media{}
				if data.ImgFront != "" {
					r = append(r, kyc.Media{
						Key:  "imgFront",
						Type: kyc.MediaTypeImage,
						URL:  data.ImgFront,
					})
				}
				if data.ImgBack != "" {
					r = append(r, kyc.Media{
						Key:  "imgBack",
						Type: kyc.MediaTypeImage,
						URL:  data.ImgBack,
					})
				}
				if data.ImgHandheld != "" {
					r = append(r, kyc.Media{
						Key:  "imgHandheld",
						Type: kyc.MediaTypeImage,
						URL:  data.ImgHandheld,
					})
				}
				return &r
			}()
			//More     *[]KeyValue     `json:"more,omitempty"`
			cmd.More = func() *[]kyc.KeyValue {
				return &[]kyc.KeyValue{
					{Key: "relationStore", Value: data.RelationStore},
					{Key: "employmentType", Value: data.EmploymentType},
					{Key: "currentAddress", Value: data.CurrentAddress},
					{Key: "nationality", Value: data.Nationality},
					{Key: "permanentAddress", Value: data.PermanentAddress},
					{Key: "mainSourceOfFunds", Value: data.MainSourceOfFunds},
					{Key: "occupation", Value: data.Occupation},
				}
			}()
		}
		id := types.NewID()
		_, err := kyc.ApplyService.Request(newDDDCtx(id), cmd)
		if err != nil {
			l.Error(errorx.WithStack(err))
			return err
		}
		switch data.Status {
		case 1:
		case 2:
			_, err = kyc.ApproveService.Request(newDDDCtx(id), kyc.ApproveCommand{
				Comment: "approved",
			})
		case 3:
			_, err = kyc.RejectService.Request(newDDDCtx(id), kyc.RejectCommand{
				Comment: "rejected",
			})
		default:
		}
		if err != nil {
			l.Error(errorx.WithStack(err))
			return err
		}
		return nil
	})
	for _, item := range kycs[user.User.ID] {
		wp.Submit(item)
	}
	wp.Stop()
	return nil
}

func newDDDCtx(aggregateID types.ID, transactionID ...string) *ddd.Context {
	ctx := &ddd.Context{
		AggregateID:   types.NewID(),
		Operator:      types.NewID().String(),
		SessionID:     types.NewID().String(),
		TraceID:       types.NewID(),
		TransactionID: types.NewID().String(),
		Caller:        "migration",
	}
	if aggregateID != types.EmptyID {
		ctx.AggregateID = aggregateID
	}
	if len(transactionID) != 0 && transactionID[0] != "" {
		ctx.TransactionID = transactionID[0]
	}
	return ctx
}

var (
	gameTypes = map[int32]betting.GameCategory{
		1:  betting.CategorySport,
		2:  betting.CategorySlot,
		3:  betting.CategoryLive,
		4:  betting.CategoryFish,
		5:  betting.CategoryChess,
		6:  betting.CategorySport,
		11: betting.CategoryChess,
		12: betting.CategorySlot,
	}
	games  = map[int32]*model.GameList{}
	brands = map[int32]string{
		1:  "JDB",
		2:  "BNG",
		6:  "CQ9",
		8:  "JiLi",
		9:  "FB",
		10: "HOT#_SLOTS",
		11: "TPG",
		13: "PB",
		14: "TBS",
		15: "Booming",
		26: "Filbet_Sports",
		27: "EVO",
		28: "RTG",
		29: "HABA",
		30: "FUNKY",
	}
)

func InitGame() {
	db := oldmain.Use(clientz.MainDB())
	{
		data1, err := db.GameList.Find()
		if err != nil {
			logrus.Error(errorx.WithStack(err))
			return
		}
		for _, item := range data1 {
			games[item.ID] = item
		}
	}
}

func getWinBetSlip(in *model.WinBetslips) (*WinBetSlips, errorx.Error) {
	data, err := winBetSlipCollection.GetOne(bson.M{"bet.id": in.ID})
	if err == nil && data != nil {
		return data, nil
	}
	data = &WinBetSlips{
		Entity: types.Entity{
			ID: types.NewID(),
		},
		Bet: in,
	}
	err = winBetSlipCollection.Create(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func Init() errorx.Error {
	InitGame()
	{
		data1, err := oldmain.Use(clientz.MainDB()).WinUserKyc.Find()
		if err != nil {
			logrus.Error(errorx.WithStack(err))
		}
		for _, item := range data1 {
			_, ok := kycs[item.UID]
			if !ok {
				kycs[item.UID] = []model.WinUserKyc{}
			}
			kycs[item.UID] = append(kycs[item.UID], *item)
		}
	}
	{
		data1, err := oldmain.Use(clientz.MainDB()).WinCoinDepositRecord.Find()
		if err != nil {
			return errorx.WithStack(err)
		}
		for _, record := range data1 {
			_, ok := deposits[record.UID]
			if !ok {
				deposits[record.UID] = map[int64]*model.WinCoinDepositRecord{}
			}
			deposits[record.UID][record.ID] = record
		}
	}
	{
		data1, err := oldmain.Use(clientz.MainDB()).WinCoinWithdrawalRecord.Find()
		if err != nil {
			return errorx.WithStack(err)
		}
		for _, record := range data1 {
			_, ok := withdraws[record.UID]
			if !ok {
				withdraws[record.UID] = map[int64]*model.WinCoinWithdrawalRecord{}
			}
			withdraws[record.UID][record.ID] = record
		}
	}
	{
		data1, err := oldsharding.Use(clientz.ShardingDB()).WinUserWallet.Find()
		if err != nil {
			return errorx.WithStack(err)
		}
		for _, item := range data1 {
			_, ok := userWallets[item.UID]
			if ok {
				continue
			}
			userWallets[item.UID] = item
		}
	}
	return nil
}
