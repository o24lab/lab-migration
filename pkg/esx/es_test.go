package esx

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"migration/pkg/model"
	"testing"
)

func TestOne(t *testing.T) {
	var winBetSlipCollection = Define[model.WinBetslips]("win_betslips")
	one, err := winBetSlipCollection.GetOne(elastic.NewTermQuery("id", 785242934239964770))
	if err != nil {
		logrus.Error(err)
	}
	fmt.Println(one)
}

func TestList(t *testing.T) {
	var winCoinLogCollection = Define[model.WinCoinLog]("win_coin_log")
	list, err := winCoinLogCollection.GetList(elastic.NewTermQuery("uid", 422404))
	if err != nil {
		logrus.Error(err)
		return
	}
	logrus.WithField("len", len(list)).Info("success")
}
