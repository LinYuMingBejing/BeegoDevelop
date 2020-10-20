package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "login.html"
}

func (c *LoginController) DoLogin() {
	username := c.GetString("username")
	password := c.GetString("password")
	fmt.Printf("用戶名 : %v === 密碼:%v", username, password)
	// 執行跳轉
	c.Redirect("/", 302)
}
