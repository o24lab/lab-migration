package migration

import (
	"github.com/sirupsen/logrus"
	"migration/pkg/model"
	"pkg/logger"
	"pkg/types"
	"testing"
)

func TestUser(t *testing.T) {
	logger.Init()
	err := Migrate()
	if err != nil {
		logrus.Error(err)
	}
	logrus.Info("migrate success")
}

func TestTb(t *testing.T) {
	err := userCollection.Create(&WinUser{
		Entity: types.Entity{
			ID: types.NewID(),
		},
		User: model.WinUser{
			ID: 111,
		},
	})
	if err != nil {
		logrus.Error(err)
	}
}
