package esx

import (
	"context"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"pkg/config"
	"pkg/once"
)

type Config struct {
	Url      string `yaml:"url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

var conf = config.EtcdYamlConfig[Config](Config{
	Url:      "http://192.168.88.23:9201",
	Username: "elastic",
	Password: "1+u-89oHI=Pq3u8v39Wc",
})

var GetClient = once.F[*elastic.Client](func() *elastic.Client {
	c, err := elastic.NewClient(elastic.SetURL(conf.Get().Url), elastic.SetSniff(false), elastic.SetBasicAuth(conf.Get().Username, conf.Get().Password))
	if err != nil {
		logrus.Fatal(err)
	}
	{
		_, _, err = c.Ping(conf.Get().Url).Do(context.Background())
		if err != nil {
			logrus.Fatal(err)
		}
	}
	return c
})
