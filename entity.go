package migration

import (
	"domain/user/kyc"
	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"migration/pkg/esx"
	"migration/pkg/model"
	"pkg/mongox"
	"pkg/types"
)

const databaseName = "migration"

var winUserCollection = mongox.Define[WinUser](databaseName, mongox.Indexes(mongo.IndexModel{
	Keys: bson.D{
		{Key: "user.id", Value: 1},
	},
	Options: options.Index().SetUnique(true),
}))

type WinUser struct {
	types.Entity `bson:",inline"`
	ParentID     types.ID        `bson:"parentID"`
	User         model.WinUser   `bson:"user"`
	Balance      decimal.Decimal `bson:"balance"`
	CreateIn     string          `bson:"cmd"`
	CreateOut    string          `bson:"cmdOut"`
}

type WinUserKyc struct {
	types.Entity `bson:",inline"`
	Uid          types.ID         `bson:"uid"`
	Kyc          model.WinUserKyc `bson:"kyc"`
	Status       kyc.Status       `bson:"status"`
	ApplyIn      string           `bson:"applyIn"`
	ApplyOut     string           `bson:"applyOut"`
}

var winKycCollection = mongox.Define[WinUserKyc](databaseName, mongox.Indexes(mongo.IndexModel{
	Keys: bson.D{
		{Key: "kyc.id", Value: 1},
	},
	Options: options.Index().SetUnique(true),
}))

var winBetSlipCollection = esx.Define[model.WinBetslips]("win_betslips")

var winCoinLogCollection = esx.Define[model.WinCoinLog]("win_coin_log")
