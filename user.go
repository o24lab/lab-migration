package migration

import (
	"context"
	"domain/user"
	"domain/user/basicinfo"
	"domain/user/username"
	"encoding/json"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"github.com/svc0a/worker"
	"go.mongodb.org/mongo-driver/bson"
	"migration/pkg/clientz"
	"migration/pkg/model"
	"migration/pkg/oldmain"
	"migration/pkg/oldsharding"
	"pkg/errorx"
	"pkg/types"
	"sync"
)

var (
	userWallets = sync.Map{}
	usersMap    = sync.Map{}
)

func User() errorx.Error {
	{
		data1, err := oldsharding.Use(clientz.ShardingDB()).WinUserWallet.Find()
		if err != nil {
			return errorx.WithStack(err)
		}
		wp := worker.New[model.WinUserWallet, errorx.Error](workerNum, worker.WithChanSize[model.WinUserWallet, errorx.Error](int64(len(data1))), worker.WithErrHandler[model.WinUserWallet, errorx.Error](func(err errorx.Error) {
			if err != nil {
				logrus.Error(err)
			}
		}))
		for _, item := range data1 {
			wp.Submit(*item)
		}
		wp.Start(func(record model.WinUserWallet) errorx.Error {
			userWallets.Store(record.UID, record)
			return nil
		})
		wp.Stop()
	}
	db := oldmain.Use(clientz.MainDB()).WinUser
	data, err := db.Find()
	if err != nil {
		return errorx.WithStack(err)
	}
	{
		wp := worker.New[model.WinUser, errorx.Error](workerNum, worker.WithChanSize[model.WinUser, errorx.Error](int64(len(data))), worker.WithErrHandler[model.WinUser, errorx.Error](func(err errorx.Error) {
			if err != nil {
				logrus.Error(err)
			}
		}))
		for _, item := range data {
			wp.Submit(*item)
		}
		wp.Start(func(data model.WinUser) errorx.Error {
			id := types.NewID()
			u := WinUser{
				Entity: types.Entity{
					ID: id,
				},
				User: data,
			}
			wallet, ok := userWallets.Load(data.ID)
			if ok {
				wallet1, ok := wallet.(model.WinUserWallet)
				if ok {
					u.Balance = decimal.NewFromFloat(wallet1.Coin)
				}
			}
			err1 := winUserCollection.Create(&u)
			if err1 != nil {
				return err1
			}
			usersMap.Store(data.ID, u.ID)
			return nil
		})
		wp.Stop()
	}
	{
		users, err1 := winUserCollection.GetList(bson.M{}, context.TODO())
		if err1 != nil {
			return err1
		}
		wp := worker.New[WinUser, errorx.Error](workerNum, worker.WithChanSize[WinUser, errorx.Error](int64(len(data))), worker.WithErrHandler[WinUser, errorx.Error](func(err errorx.Error) {
			if err != nil {
				logrus.Error(err)
			}
		}))
		for _, item := range *users {
			wp.Submit(item)
		}
		wp.Start(func(data WinUser) errorx.Error {
			{
				parentID, ok := usersMap.Load(data.User.SupUid1)
				if !ok {
					return nil
				}
				pid, ok := parentID.(types.ID)
				if !ok {
					return nil
				}
				data.ParentID = pid
			}
			{
				cmd := user.CreateUserCommand{}
				if data.User.Mobile != "" {
					cmd.PhoneNumber = &types.PhoneNumber{
						CountryCode: "63",
						Number:      data.User.Mobile,
					}
				}
				if data.User.Email != "" {
					cmd.Email = func() *types.Email {
						email := types.Email(data.User.Email)
						return &email
					}()
				}
				if data.User.Username != "" {
					cmd.Username = func() *username.Username {
						username1 := username.Username(data.User.Username)
						return &username1
					}()
				}
				cmd.BasicInfo = &basicinfo.UpdateCommand{}
				if data.User.Avatar != "" {
					cmd.BasicInfo.AvatarUrl = func() *basicinfo.AvatarUrl {
						n := basicinfo.AvatarUrl(data.User.Avatar)
						return &n
					}()
				}
				if data.User.Username != "" {
					cmd.BasicInfo.Nickname = func() *basicinfo.Nickname {
						n := basicinfo.Nickname(data.User.Username)
						return &n
					}()
				}
				if !data.ParentID.IsEmpty() {
					cmd.ParentID = &data.ParentID
				}
				cmd.BasicInfo.Gender = func() *types.Gender { g := types.GenderUnKnow; return &g }()
				data.CreateIn = func() string {
					marshal, _ := json.Marshal(cmd)
					return string(marshal)
				}()
			}
			err2 := winUserCollection.UpdateWithOptimisticLock(&data, context.TODO())
			if err2 != nil {
				return err2
			}
			return nil
		})
		wp.Stop()
	}
	users, err1 := winUserCollection.GetList(bson.M{}, context.TODO())
	if err1 != nil {
		return err1
	}
	{
		wp := worker.New[WinUser, errorx.Error](workerNum, worker.WithChanSize[WinUser, errorx.Error](int64(len(data))), worker.WithErrHandler[WinUser, errorx.Error](func(err errorx.Error) {
			if err != nil {
				logrus.Error(err)
			}
		}))
		for _, item := range *users {
			wp.Submit(item)
		}
		wp.Start(func(data WinUser) errorx.Error {
			cmd := user.CreateUserCommand{}
			{
				err2 := json.Unmarshal([]byte(data.CreateIn), &cmd)
				if err2 != nil {
					return errorx.WithStack(err2)
				}
			}
			_, err2 := user.CreateUserService.Request(newDDDCtx(data.ID), cmd)
			if err2 != nil {
				return err2
			}
			return nil
		})
		wp.Stop()
	}
	return nil
}
