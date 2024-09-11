package main

import (
	"context"
	"fmt"
	"github.com/libi/dcron/cron"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"go.mongodb.org/mongo-driver/bson"
	"math"
	"migration"
	"migration/pkg"
	"pkg/logger"
	"pkg/mongox"
	"time"
)

func main() {
	logger.Init()
	go func() {
		startTime := time.Now()
		var count int64
		c := cron.New(cron.WithSeconds())
		c.Start()
		_, err := c.AddFunc("0 * * * * *", func() {
			t := time.Now()
			result := bson.M{}
			err := mongox.GetClient().Database("migration").RunCommand(context.TODO(), bson.D{{"collStats", "migration_winuser"}}).Decode(&result)
			if err != nil {
				logrus.Error(err)
				return
			}
			dif := cast.ToInt64(result["count"]) - count
			count = cast.ToInt64(result["count"])
			logrus.WithField("start", startTime.Format(time.DateTime)).WithField("currentTime", t.Format(time.DateTime)).WithField("duration(min)", math.Ceil(t.Sub(startTime).Minutes())).WithField("userTotal", count).WithField("users/min", dif).Info("statistics")
			go func() {
				err1 := pkg.TgSend(fmt.Sprintf("start=%v,\ncurrent=%v,\nduration(%v min),\ntotal=%v,\nuser/min=%v", startTime.Format(time.DateTime), t.Format(time.DateTime), math.Ceil(t.Sub(startTime).Minutes()), count, dif))
				logrus.Error(err1)
			}()
		})
		if err != nil {
			logrus.Error(err)
		}
	}()
	err := migration.Migrate()
	if err != nil {
		logrus.Error(err)
	}
	logrus.Info("migrate success")
}
