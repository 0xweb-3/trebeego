package routers

import (
	"trybeego/controllers"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, Cors)

	// 注册路由注解
	beego.Include(&controllers.AnnotationController{})

	// user
	beego.Router("/users", &controllers.UserController{}, "get:GetAll;post:Create")
	beego.Router("/users/:id", &controllers.UserController{}, "get:GetOne;put:Update;delete:Delete")

	// 自动路由模式
	beego.AutoRouter(&controllers.AutoController{})

	// RESTful风格路由
	beego.Router("/rest", &controllers.RestController{})

}

var Cors = cors.Allow(&cors.Options{
	AllowAllOrigins: true, // 允许所有域名访问
	AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	AllowHeaders: []string{
		"content-type", "DeviceId", "MarketChannel", "DeviceModel", "deviceid", "ApiVersionNum", "marketchannel",
		"ClientVersion", "Platform", "DeviceModel", "ApiVersionNum", "TimeStamp", "BundleId", "OsVersion", "Signature", "User-Agent", "language",
		"Authorization", "SessionId", "NetworkType",
	},
	ExposeHeaders:    []string{"Content-Length"},
	AllowCredentials: false,
})
