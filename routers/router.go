package routers

import (
	"awesomeProject1/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	// 查询列表
	beego.Router("/user/list", &controllers.MainController{}, "get:List")
}
