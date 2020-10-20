package main

import (
	_ "WEB/routers"

	"github.com/astaxie/beego"
)

func main() {
	// 1. 配置redis數據庫的連接地址
	beego.AppConfig.Set("redisuser", "admin666")
	// 2. 加載其他配置文件
	beego.LoadAppConfig("INI", "config/redis.conf")
	// beego.AddFuncMap("unixToDate", models.unixToDate)
	beego.Run()
}
