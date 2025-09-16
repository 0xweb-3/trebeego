package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

// AnnotationController
type AnnotationController struct {
	beego.Controller
}

// @router /annotation/:id [get]
func (c *AnnotationController) Get() {
	id := c.Ctx.Input.Param(":id")
	c.Ctx.WriteString("User Id = " + id)
}
