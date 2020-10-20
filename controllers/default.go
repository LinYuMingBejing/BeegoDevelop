package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type Article struct {
	Title   string
	Content string
}

func (c *MainController) Get() {
	// 1. 模板中綁定基本數據類型 字符串 數值 布爾值
	c.Data["Website"] = "beego 課程"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["num"] = 12
	c.Data["flag"] = true

	// 2. 模板中綁定結構體數據
	article := Article{
		Title:   "我是Golang課程",
		Content: "實戰項目",
	}
	c.Data["article"] = article

	// 3. 模板中循環遍歷 range，模擬循環切片
	c.Data["sliceList"] = []string{"php", "python", "hadoop", "elk"}

	// 4. 模板中循環遍歷 range，模擬循環Map
	userinfo := make(map[string]interface{})
	userinfo["username"] = "靜茹"
	userinfo["age"] = 20

	c.TplName = "index.html"
}

func (c *MainController) ConfigGet() {
	// 	獲取app.conf裡面的配置參數
	// 1. String 獲取一個值
	mysqluser := beego.AppConfig.String("mysqluser")
	beego.Info(mysqluser)

	// 2. Strings 獲取一個數組
	mysqlpass := beego.AppConfig.String("mysqlpass")
	beego.Info(mysqlpass)

	// 3. 獲取beego.AppConfig.Set()的值
	redisuser := beego.AppConfig.String("redisuser")
	beego.Info(redisuser)

	// 4. 獲取其他配置文件信息
	redispass := beego.AppConfig.String("redispass")
	beego.Info(redispass)

	c.Ctx.WriteString("獲取配置參數")
}
