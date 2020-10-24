package main

import (
	_ "WEB/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	username := beego.AppConfig.String(username)
	pwd := beego.AppConfig.String(pwd)
	host := beego.AppConfig.String(host)
	port := beego.AppConfig.String(port)
	db := beego.AppConfig.String(db)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", username+":"+"pwd"+"@tcp("+host+":"+port+")/"+"db"+"?charset=utf8")

}

func main() {
	// 1. 配置redis數據庫的連接地址
	beego.AppConfig.Set("redisuser", "admin666")
	// 2. 加載其他配置文件
	beego.LoadAppConfig("INI", "config/redis.conf")
	// beego.AddFuncMap("unixToDate", models.unixToDate)
	beego.Run()
}
