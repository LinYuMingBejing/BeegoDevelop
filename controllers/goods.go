package controllers

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
)

type GoodsController struct {
	beego.Controller
}

func (c *GoodsController) Get() {
	c.Data["title"] = "beego.me"
	c.Data["product"] = "astaxie@gmail.com"
	c.TplName = "goods.html"
}

type Product struct {
	Title   string `form:"title"`
	Content string `form:"content"`
}

func (c *GoodsController) DoAdd() {
	c.Ctx.WriteString("執行增加操作")
}

func (c *GoodsController) DoEdit() {
	p := Product{}
	err := c.ParseForm(&p)
	if err != nil {
		c.Ctx.WriteString("執行修改失敗")
		return
	}
	fmt.Printf("%#v", p)

	c.Ctx.WriteString("執行修改操作")
}

func (c *GoodsController) DoDelete() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Ctx.WriteString("參數錯誤")
		return
	}

	c.Ctx.WriteString("執行刪除操作" + strconv.Itoa(id))
}
