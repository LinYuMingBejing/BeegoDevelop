package controllers

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
)

type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) Get() {
	now := time.Now()
	fmt.Println(now)
	c.Ctx.WriteString("新聞列表")
}

func (c *ArticleController) AddArticle() {
	c.Ctx.WriteString("增加新聞")
}

// func (c *ArticleController) EditArticle() {
// 	c.Ctx.WriteString("修改新聞")
// }

func (c *ArticleController) EditArticle() {
	id, err := c.GetInt("id")
	if err != nil {
		beego.Info(err)
		c.Ctx.WriteString("傳參錯誤")
	}
	beego.Info(id)
	c.Ctx.WriteString("修改新聞")
}
