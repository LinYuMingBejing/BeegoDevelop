package controllers

import (
	"github.com/astaxie/beego"
)

type CmsController struct {
	beego.Controller
}

func (c *CmsController) Get() {
	cmsId := c.Ctx.Input.Param(":id")
	c.Ctx.WriteString("CMS接口----" + cmsId)
}
