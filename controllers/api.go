package controllers

import (
	"github.com/astaxie/beego"
)

type ApiController struct {
	beego.Controller
}

func (c *ApiController) Get() {
	// 獲取動態路由的值
	id := c.Ctx.Input.Param(":id")
	c.Ctx.WriteString("api接口----" + id)
}
