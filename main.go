package main

import (
	"awesomeProject1/models"
	_ "awesomeProject1/routers"

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
	// 注册数据库
	orm.RegisterDataBase(
		"default",
		"mysql",
		"root:lzc123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=true&loc=Local",
	)

	// 注册模型
	orm.RegisterModel(new(models.User))

	// 自动建表
	// 第一个 true：没有表就创建
	// 第二个 true：打印 SQL
	orm.RunSyncdb("default", false, true)

}
