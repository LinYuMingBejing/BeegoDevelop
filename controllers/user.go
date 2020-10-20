package controllers

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	c.Ctx.WriteString("用戶中心")
}

func (c *UserController) AddUser() {
	c.TplName = "user.html"
}

func (c *UserController) DoAddUser() {
	id, err := c.GetInt("ID")
	if err != nil {
		beego.Info(err)
		c.Ctx.WriteString("ID必須是數字")
		return
	}
	username := c.GetString("username")
	password := c.GetString("password")
	hobby := c.GetStrings("hobby")
	fmt.Printf("值 : %v === 類型:%T", hobby, hobby)
	c.Ctx.WriteString("增加用戶--" + strconv.Itoa(id) + username + password)
}

// 定義一個User結構體
type User struct {
	Username string   `form:"username json:"username"`
	Password string   `form:"password" json:"password"`
	Hobby    []string `from:hobby json:"hobby"`
}

func (c *UserController) GetUser() {
	u := User{
		Username: "情歌天后",
		Password: "a123456",
		Hobby:    []string{"作曲", "作詞"},
	}
	// 返回一個json字串
	c.Data["json"] = u
	c.ServeJSON()
}
