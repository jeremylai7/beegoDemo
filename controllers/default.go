package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.vip"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["AEmail"] = "xxxxx@qq.com"
	c.TplName = "index.tpl"
}

func (c *MainController) Post() {
	c.Data["json"] = map[string]interface{}{
		"success": true,
		"code":    200,
		"msg":     "这是返回的json数据",
	}
	c.ServeJSON()
}
