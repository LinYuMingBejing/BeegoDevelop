package controllers

import (
	"github.com/astaxie/beego"
)

type CookieController struct {
	beego.Controller
}

func (c *CookieController) Set() {
	// 1. 設置cookie
	c.Ctx.SetCookie("username", "John&Mary")

	// 2. 設置cookie過期時間 (單位:秒)
	c.Ctx.SetCookie("username", "Larry", 20)

	// 3. 設置cookie的訪問路徑
	c.Ctx.SetCookie("username", "Larry", 20, "/article")

	// 4. 設置cookie的路徑Domain
	c.Ctx.SetCookie("username", "Larry", 20, "/", ".etu.com")

	// 5. SetCookie無法配置中文cookie
	c.Ctx.SetCookie("username", "中文")

	// 6. cookie的加密，設置中文cookie，第一個參數表示密鑰
	c.Ctx.SetSecureCookie("aaaaa", "username", "李四", 1000)

}

func (c *CookieController) Get() {
	// 1. 獲取普通Cookie，實現數據共享
	username := c.Ctx.GetCookie("username")
	c.Data["username"] = username

	// 2. 獲取加密cookie
	username2, _ := c.Ctx.GetSecureCookie("aaaaa", "username")
	beego.Info(username2)

}
