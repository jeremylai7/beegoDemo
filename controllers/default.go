package controllers

import (
	"encoding/json"
	"fmt"

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

type UserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c *MainController) Post() {
	var req UserReq
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	fmt.Println(req.Username, req.Password)

	c.Data["json"] = map[string]interface{}{
		"success": true,
		"code":    200,
		"msg":     "你传参 name 是" + req.Username,
	}
	c.ServeJSON()
}
