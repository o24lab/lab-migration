package clientz

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"pkg/config"
	"pkg/once"
)

type Config struct {
	Main     string `yaml:"main"`
	Sharding string `yaml:"sharding"`
}

//	var conf = config.EtcdYamlConfig[Config](Config{
//		Main:     "go_acct:BJf6dcY56UFQucr71kWf@tcp(192.168.88.24:3306)/filbet_prod_main",
//		Sharding: "go_acct:BJf6dcY56UFQucr71kWf@tcp(192.168.88.24:3306)/filbet_prod_sharding",
//	})
var conf = config.EtcdYamlConfig[Config](Config{
	Main:     "root:KK#%213Sse!!aq@tcp(192.168.88.23:3306)/filbet_prod_main",
	Sharding: "root:KK#%213Sse!!aq@tcp(192.168.88.23:3306)/filbet_prod_sharding",
})

var (
	mainDB     *gorm.DB
	shardingDB *gorm.DB
)

func MainDB() *gorm.DB {
	if mainDB != nil {
		return mainDB
	}
	db := once.F[*gorm.DB](func() *gorm.DB {
		// 配置第一个数据库
		db1, err := gorm.Open(mysql.Open(conf.Get().Main), &gorm.Config{})
		if err != nil {
			panic("failed to connect to database 1")
		}
		// 获取底层 *sql.DB 对象
		sqlDB, err := db1.DB()
		if err != nil {
			panic("failed to get *sql.DB from gorm.DB")
		}
		// 配置连接池
		sqlDB.SetMaxOpenConns(300)  // 设置最大打开连接数
		sqlDB.SetMaxIdleConns(300)  // 设置最大闲置连接数
		sqlDB.SetConnMaxLifetime(0) // 设置连接最大生命周期，0表示不限制
		return db1
	})
	return db()
}

func ShardingDB() *gorm.DB {
	if shardingDB != nil {
		return shardingDB
	}
	// 配置第二个数据库
	db2, err := gorm.Open(mysql.Open(conf.Get().Sharding), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database 2")
	}
	shardingDB = db2
	return shardingDB
}
