package routers

import (
	"trybeego/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	// 注册路由注解
	beego.Include(&controllers.AnnotationController{})

	// user
	beego.Router("/users", &controllers.UserController{}, "get:GetAll;post:Create")
	beego.Router("/users/:id", &controllers.UserController{}, "get:GetOne;put:Update;delete:Delete")

	// 自动路由模式
	beego.AutoRouter(&controllers.AutoController{})
}
