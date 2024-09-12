package migration

import (
	"context"
	"domain/user/kyc"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/svc0a/worker"
	"go.mongodb.org/mongo-driver/bson"
	"migration/pkg/clientz"
	"migration/pkg/model"
	"migration/pkg/oldmain"
	"pkg/errorx"
	"pkg/types"
	"time"
)

func Kyc() errorx.Error {
	{
		data, err := oldmain.Use(clientz.MainDB()).WinUserKyc.Find()
		if err != nil {
			logrus.Error(errorx.WithStack(err))
		}
		wp := worker.New[model.WinUserKyc, errorx.Error](workerNum, worker.WithChanSize[model.WinUserKyc, errorx.Error](int64(len(data))), worker.WithErrHandler[model.WinUserKyc, errorx.Error](func(err errorx.Error) {
			if err != nil {
				logrus.Error(err)
			}
		}))
		for _, item := range data {
			wp.Submit(*item)
		}
		wp.Start(func(data model.WinUserKyc) errorx.Error {
			user, err1 := winUserCollection.GetOne(bson.M{"user.id": data.UID})
			if err1 != nil {
				return err1
			}
			userKyc := WinUserKyc{
				Uid:    user.ID,
				Kyc:    data,
				Status: kyc.StatusPending,
			}
			{
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
						logrus.Error(err)
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
				marshal, _ := json.Marshal(cmd)
				userKyc.ApplyIn = string(marshal)
			}
			err2 := winKycCollection.Create(&userKyc)
			if err2 != nil {
				return err2
			}
			return nil
		})
		wp.Stop()
	}
	{
		data, err := winKycCollection.GetList(bson.M{
			"kyc.firstname": bson.M{"$ne": ""},
			"kyc.lastname":  bson.M{"$ne": ""},
		}, context.TODO())
		if err != nil {
			return err
		}
		wp := worker.New[WinUserKyc, errorx.Error](workerNum, worker.WithChanSize[WinUserKyc, errorx.Error](int64(len(*data))), worker.WithErrHandler[WinUserKyc, errorx.Error](func(err errorx.Error) {
			if err != nil {
				logrus.Error(err)
			}
		}))
		for _, item := range *data {
			wp.Submit(item)
		}
		wp.Start(func(data WinUserKyc) errorx.Error {
			cmd := kyc.ApplyCommand{}
			err2 := json.Unmarshal([]byte(data.ApplyIn), &cmd)
			if err2 != nil {
				return errorx.WithStack(err2)
			}
			_, err1 := kyc.ApplyService.Request(newDDDCtx(data.Uid, WithTimestamp(int64(data.Kyc.CreatedAt))), cmd)
			if err1 != nil {
				data.ApplyOut = err1.Error()
				err3 := winKycCollection.UpdateWithOptimisticLock(&data)
				if err3 != nil {
					return err3
				}
				return err1
			}
			switch data.Kyc.Status {
			case 1:
			case 2:
				_, err = kyc.ApproveService.Request(newDDDCtx(data.Uid), kyc.ApproveCommand{
					Comment: "approved",
				})
			case 3:
				_, err = kyc.RejectService.Request(newDDDCtx(data.Uid), kyc.RejectCommand{
					Comment: "rejected",
				})
			default:
			}
			if err != nil {
				return err
			}
			return nil
		})
		wp.Stop()
		return nil
	}
}
