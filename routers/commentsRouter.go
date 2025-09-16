package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["trybeego/controllers:AnnotationController"] = append(beego.GlobalControllerRouter["trybeego/controllers:AnnotationController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/annotation/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
