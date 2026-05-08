package controllers

import (
	"awesomeProject1/models"
	"encoding/json"
	"fmt"
	"time"

	"github.com/beego/beego/v2/client/orm"
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

// 新增 User
func (c *MainController) Post() {
	var req UserReq
	// 解析成功，err 为空。解析失败，返回错误信息
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"code":    500,
			"msg":     "参数错误",
		}
		c.ServeJSON()
		return
	}

	fmt.Println(req.Username, req.Password)

	o := orm.NewOrm()
	user := new(models.User)
	user.Name = req.Username
	user.SubmitTime = time.Now()
	id, err := o.Insert(user)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"code":    500,
			"msg":     "新增失败" + err.Error(),
		}
	}

	c.Data["json"] = map[string]interface{}{
		"success": true,
		"code":    200,
		"msg":     "新增成功",
		"id":      id,
	}
	c.ServeJSON()
}

// 分页查询列表
func (c *MainController) List() {
	pageIndex, _ := c.GetInt("pageIndex", 1)
	pageSize, _ := c.GetInt("pageSize", 10)
	if pageIndex <= 0 {
		pageIndex = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	o := orm.NewOrm()

	var users []models.User
	qs := o.QueryTable(new(models.User))
	total, err := qs.Count()
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"code":    500,
			"msg":     err.Error(),
		}
		c.ServeJSON()
		return
	}

	_, err = qs.OrderBy("-id").Limit(pageSize, (pageIndex-1)*pageSize).All(&users)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"code":    500,
			"msg":     err.Error(),
		}
		c.ServeJSON()
		return
	}

	c.Data["json"] = map[string]interface{}{
		"success":   true,
		"code":      200,
		"users":     users,
		"total":     total,
		"pageIndex": pageIndex,
		"pageSize":  pageSize,
	}
	c.ServeJSON()

}
