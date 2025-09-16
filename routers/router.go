package routers

import (
	"trybeego/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	//beego.Router("/", &controllers.MainController{})

	// user
	beego.Router("/users", &controllers.UserController{}, "get:GetAll;post:Create")
	beego.Router("/users/:id", &controllers.UserController{}, "get:GetOne;put:Update;delete:Delete")
}
