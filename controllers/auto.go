package controllers

import beego "github.com/beego/beego/v2/server/web"

type AutoController struct {
	beego.Controller
}

func (c *AutoController) List() {
	c.Ctx.WriteString("List all items")
}

func (c *AutoController) Detail() {
	c.Ctx.WriteString("Item detail")
}
