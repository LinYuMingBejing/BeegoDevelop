package models

import "github.com/astaxie/beego/orm"

type User struct {
	Id      int    `orm:"pk"`
	Name    string `orm:"column(username)"`
	Age     int
	Address string
}

func (u *User) TableName() string {
	return "sys_user"
}

func init() {
	orm.RegisterModel(new(User))
}
