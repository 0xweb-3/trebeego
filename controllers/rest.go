package controllers

import "github.com/beego/beego/v2/server/web"

type RestController struct {
	web.Controller
}

func (c *RestController) Get() {
	c.Ctx.WriteString("GET rest")
}

func (c *RestController) Post() {
	c.Ctx.WriteString("POST rest")
}

func (c *RestController) Put() {
	c.Ctx.WriteString("PUT rest")
}

func (c *RestController) Delete() {
	c.Ctx.WriteString("DELETE rest")
}
