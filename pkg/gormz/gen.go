package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	//// 配置第一个数据库
	//db1, err := gorm.Open(mysql.Open("root:KK#%213Sse!!aq@tcp(192.168.88.23:3306)/filbet_prod_main"), &gorm.Config{})
	//if err != nil {
	//	panic("failed to connect to database 1")
	//}

	// 配置第二个数据库
	db2, err := gorm.Open(mysql.Open("root:KK#%213Sse!!aq@tcp(192.168.88.23:3306)/filbet_prod_sharding"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database 2")
	}

	// 生成代码
	//generateCode(db1, "oldmain")
	generateSharding(db2, "oldsharding")
}

func generateCode(db *gorm.DB, outputDir string) {
	g := gen.NewGenerator(gen.Config{
		OutPath: outputDir,
		Mode:    gen.WithDefaultQuery | gen.WithoutContext,
	})
	g.UseDB(db)

	// 生成代码，使用数据库的模型
	g.ApplyBasic(g.GenerateAllTable()...)
	g.Execute()
}

func generateSharding(db *gorm.DB, outputDir string) {
	g := gen.NewGenerator(gen.Config{
		OutPath: outputDir,
		Mode:    gen.WithDefaultQuery | gen.WithoutContext,
	})
	g.UseDB(db)

	// 生成代码，使用数据库的模型
	g.ApplyBasic(g.GenerateModel("win_betslips_0"),
		g.GenerateModel("win_betslips_details_0"),
		g.GenerateModel("win_coin_log_0"),
		g.GenerateModel("win_user_wallet"),
	)
	g.Execute()
}
