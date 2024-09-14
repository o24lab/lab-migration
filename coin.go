package migration

import (
	"context"
	"domain/balance"
	"domain/betting"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"github.com/svc0a/worker"
	"go.mongodb.org/mongo-driver/bson"
	"migration/pkg/clientz"
	"migration/pkg/model"
	"migration/pkg/oldmain"
	"pkg/ddd"
	"pkg/errorx"
	"pkg/types"
	"sync"
)

const workerNum = 10000

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

func Coin() errorx.Error {
	{
		err := initData()
		if err != nil {
			return err
		}
	}
	data, err1 := winUserCollection.GetList(bson.M{}, context.TODO())
	if err1 != nil {
		return err1
	}
	wp := worker.New[WinUser, errorx.Error](workerNum, worker.WithChanSize[WinUser, errorx.Error](int64(len(*data))), worker.WithErrHandler[WinUser, errorx.Error](func(err errorx.Error) {
		if err != nil {
			if errorx.ErrPhoneNumberInvalid.Is(err) {
				return
			}
			logrus.Error(err)
		}
	}))
	for _, item := range *data {
		wp.Submit(item)
	}
	wp.Start(func(user1 WinUser) errorx.Error {
		winCoinLogs, err := winCoinLogCollection.GetList(elastic.NewTermQuery("uid", user1.User.ID))
		if err != nil {
			return err
		}
		m, err := initBet(user1.User.ID)
		if err != nil {
			return err
		}
		for _, winCoinLog := range winCoinLogs {
			coinLog1 := NewWinCoinLog(&user1, &winCoinLog, m)
			err1 := coinLog1.Validate()
			if err1 != nil {
				logrus.Error(err1)
			}
			//1-存款 2-提款 3-投注 4-派彩 5-返水 6-佣金 7-活动(奖励) 8-系统调账 9-退款 10-佣金钱包转主账户余额 11-小费" json:"category
			switch winCoinLog.Category {
			case 1:
				err2 := coinLog1.Deposit()
				if err2 != nil {
					logrus.Error(err2)
				}
			case 2:
				err2 := coinLog1.Withdraw()
				if err2 != nil {
					logrus.Error(err2)
				}
			case 3:
				one, err2 := winBetSlipCollection.GetOne(elastic.NewTermQuery("id", winCoinLog.ReferID))
				if err2 == nil && one != nil {
					err3 := coinLog1.PlaceBet()
					if err3 != nil {
						logrus.Error(err3)
					}
				}
			case 4:
				one, err2 := winBetSlipCollection.GetOne(elastic.NewTermQuery("id", winCoinLog.ReferID))
				if err2 == nil && one != nil {
					err3 := coinLog1.Payout()
					if err3 != nil {
						logrus.Error(err3)
					}
				}
			case 5, 6, 7, 8, 10, 11:
				err2 := coinLog1.Adjust()
				if err2 != nil {
					logrus.Error(err2)
				}
			case 9:
				one, err2 := winBetSlipCollection.GetOne(elastic.NewTermQuery("id", winCoinLog.ReferID))
				if err2 == nil && one != nil {
					err3 := coinLog1.Cancel()
					if err3 != nil {
						logrus.Error(err3)
					}
				}
			default:
				err2 := coinLog1.Adjust()
				if err2 != nil {
					logrus.Error(err2)
				}
			}
		}
		{
			b, err1 := balance.GetUserCurrentBalanceService.Request(newDDDCtx(user1.ID), user1.ID)
			if err1 != nil {
				return err1
			}
			dif := user1.Balance.Sub(b.Available)
			if !dif.IsZero() {
				_, err2 := balance.RecordService.Request(newDDDCtx(user1.ID), balance.RecordCommand{
					Action: balance.ActionAdjust,
					Amount: dif,
					Force:  true,
				})
				if err2 != nil {
					logrus.Error(err2)
				}
			}
		}
		return nil
	})
	wp.Stop()
	return nil
}

func initBet(id int32) (*sync.Map, errorx.Error) {
	m := &sync.Map{}
	winBetslips, err := winBetSlipCollection.GetList(elastic.NewTermQuery("xb_uid", id))
	if err != nil {
		return nil, err
	}
	wp := worker.New[model.WinBetslips, errorx.Error](workerNum, worker.WithChanSize[model.WinBetslips, errorx.Error](int64(len(winBetslips))), worker.WithErrHandler[model.WinBetslips, errorx.Error](func(err errorx.Error) {
		if err != nil {
			logrus.Error(err)
		}
	}))
	for _, betslip := range winBetslips {
		wp.Submit(betslip)
	}
	wp.Start(func(data model.WinBetslips) errorx.Error {
		m.Store(data.ID, data)
		return nil
	})
	wp.Stop()
	return m, nil
}

type CtxOption func(ctx *ddd.Context)

func WithTransactionID(in string) CtxOption {
	return func(ctx *ddd.Context) {
		ctx.TransactionID = in
	}
}

func WithTimestamp(in int64) CtxOption {
	return func(ctx *ddd.Context) {
		ctx.Timestamp = in
	}
}

func newDDDCtx(aggregateID types.ID, options ...CtxOption) *ddd.Context {
	ctx := &ddd.Context{
		AggregateID:   aggregateID,
		Operator:      types.NewID().String(),
		SessionID:     types.NewID().String(),
		TraceID:       types.NewID(),
		TransactionID: types.NewID().String(),
		Caller:        "migration",
	}
	for _, o := range options {
		o(ctx)
	}
	return ctx
}

func initData() errorx.Error {
	db := oldmain.Use(clientz.MainDB())
	data1, err := db.GameList.Find()
	if err != nil {
		logrus.Error(errorx.WithStack(err))
		return errorx.WithStack(err)
	}
	for _, item := range data1 {
		games[item.ID] = item
	}
	return nil
}
