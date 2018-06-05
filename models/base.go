package models

import (
	"fmt"
	"github.com/robfig/config"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	fmt.Println("start base_model init")
	GetConnDB()
}

func GetConnDB() {
	// PostgreSQL 配置
	orm.RegisterDriver("postgres", orm.DRPostgres) // 注册驱动

	// set default database
	DbConfigInfo := GetDbConfig()
	dataSource := "user=" + DbConfigInfo["DB_USER"] + " password=" + DbConfigInfo["DB_PASS"] + " dbname=" + DbConfigInfo["DB_NAME"] + " host=" + DbConfigInfo["DB_HOST"] + " port=" + DbConfigInfo["DB_PORT"] + " sslmode=disable"

	orm.RegisterDataBase("default", "postgres", dataSource)

	orm.RegisterModel(new(User))
	// create table
	orm.RunSyncdb("default", false, true)
	orm.Debug = true

	fmt.Println("PQSQL", DbConfigInfo["DB_USER"]+":"+DbConfigInfo["DB_PASS"]+"@/"+DbConfigInfo["DB_NAME"]+"?charset=utf8")

}
func GetDbConfig() map[string]string {
	iniconf, _ := config.ReadDefault("D:/GoWork/src/goproject/conf/db.conf")
	DbConfigInfo := make(map[string]string)
	DbConfigInfo["DB_HOST"], _ = iniconf.String("DB", "DB_HOST")
	DbConfigInfo["DB_PORT"], _ = iniconf.String("DB", "DB_PORT")
	DbConfigInfo["DB_USER"], _ = iniconf.String("DB", "DB_USER")
	DbConfigInfo["DB_PASS"], _ = iniconf.String("DB", "DB_PASS")
	DbConfigInfo["DB_NAME"], _ = iniconf.String("DB", "DB_NAME")

	return DbConfigInfo
}

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
