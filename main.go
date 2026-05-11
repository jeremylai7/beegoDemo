package main

import (
	"awesomeProject1/models"
	_ "awesomeProject1/routers"
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

func init() {
	// 打印 sql
	orm.Debug = true

	driver, _ := beego.AppConfig.String("db_driver")
	user, _ := beego.AppConfig.String("db_user")
	password, _ := beego.AppConfig.String("db_password")
	host, _ := beego.AppConfig.String("db_host")
	port, _ := beego.AppConfig.String("db_port")
	dbName, _ := beego.AppConfig.String("db_name")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		user,
		password,
		host, port,
		dbName)

	// 注册数据库
	orm.RegisterDataBase("default", driver, dsn)

	// 注册模型
	orm.RegisterModel(new(models.User))

	// 自动建表
	// 第一个 true：没有表就创建
	// 第二个 true：建表打印 SQL
	orm.RunSyncdb("default", false, true)

}
