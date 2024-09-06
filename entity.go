package migration

import (
	"application/query/userquery"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"migration/pkg/model"
	"pkg/mongox"
	"pkg/types"
)

var userCollection = mongox.Define[WinUser]("migration", mongox.Indexes(mongo.IndexModel{
	Keys: bson.D{
		{Key: "user.id", Value: 1},
	},
	Options: options.Index().SetUnique(true),
}))

var userQueryCollection = mongox.Define[userquery.User]("application")

type WinUser struct {
	types.Entity `bson:",inline"`
	User         model.WinUser `bson:"user"`
}

type WinBetSlips struct {
	types.Entity `bson:",inline"`
	Bet          *model.WinBetslips `json:"bet" bson:"bet"`
}

var winBetSlipCollection = mongox.Define[WinBetSlips]("migration", mongox.Indexes(mongo.IndexModel{
	Keys: bson.D{
		{Key: "bet.id", Value: 1},
	},
	Options: options.Index().SetUnique(true),
}))
