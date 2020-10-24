package routers

import (
	"WEB/controllers"
	"WEB/controllers/test_orm"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/article", &controllers.ArticleController{})
	beego.Router("/article/add", &controllers.ArticleController{}, "get:AddArticle")
	beego.Router("/article/edit", &controllers.ArticleController{}, "get:EditArticle")

	// 模擬獲取前端傳入參數
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/user/add", &controllers.UserController{}, "get:AddUser")
	beego.Router("/user/doAdd", &controllers.UserController{}, "post:DoAddUser")
	beego.Router("/user/getUser", &controllers.UserController{}, "get:GetUser")

	// 模擬Restful API各種開發狀況
	beego.Router("/goods", &controllers.GoodsController{})
	beego.Router("/goods/add", &controllers.GoodsController{}, "post:DoAdd")
	beego.Router("/goods/edit", &controllers.GoodsController{}, "put:DoEdit")
	beego.Router("/goods/delete", &controllers.GoodsController{}, "delete:DoDelete")

	// 動態路由說明
	beego.Router("/api/:id", &controllers.ApiController{})

	// 配置正則路由 [路由－－偽靜態]
	beego.Router("/cms_:id([0-9]).html", &controllers.CmsController{})

	// 模擬頁面跳轉
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/doLogin", &controllers.LoginController{}, "post:DoLogin")

	// 獲取各種配置文件的變量
	beego.Router("/config", &controllers.MainController{}, "get:ConfigGet")

	// 獲取cookie
	beego.Router("/cookie/get", &controllers.CookieController{})
	beego.Router("/cookie/set", &controllers.CookieController{}, "get:Set")

	// orm
	beego.Router("/orm", &test_orm.OrmController{})

}
