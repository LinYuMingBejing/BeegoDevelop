package test_orm

import (
	"WEB/models"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type OrmController struct {
	beego.Controller
}

func (t *OrmController) Get() {
	o := orm.NewOrm()
	// o.Using("default")  默認使用default數據庫
	user := models.User{
		Name:    "梁靜茹",
		Age:     42,
		Address: "馬來西亞",
	}
	o.Insert(&user)

	// 批量插入
	users := []models.User{{Name: "孫燕姿", Age: 40, Address: "新加坡"}, {Name: "林俊傑", Age: 37, Address: "新加坡"}}
	o.InsertMulti(100, users)

	// 查詢
	query := models.User{Id: 4}
	o.Read(&query)

	queryName := models.User{Name: "孫燕姿"}
	o.Read(&queryName, "Name")

	// 查詢不到便創建
	queryUser := models.User{Name: "梁靜茹", Age: 42, Address: "馬來西亞"}
	isNew, id, err := o.ReadOrCreate(&queryUser, "Name", "Age", "Address")
	fmt.Println(isNew, id, err)

	// update 指定條件
	updateUser := models.User{Id: 4, Name: "林俊傑", Age: 22, Address: "新加坡"}
	queryErr := o.Read(&updateUser)
	if queryErr == nil {
		user.Name = "蔡健雅"
		user.Age = 23
		user.Address = "台灣台北"
		o.Update(&user)
	}

	// 刪除 前提：先查詢，再更新
	user = models.User{Name: "蔡健雅", Age: 23, Address: "台灣台北"}

	o.Delete(&user, "Name", "Age", "Address")

	// insertOrUpdate
	user = models.User{Name: "蔡健雅", Age: 23, Address: "台灣台北"}
	o.InsertOrUpdate(&user, "Id")

	t.TplName = "orm.html"
}
