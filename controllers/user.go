package controllers

import (
	"encoding/json"
	"io"
	"strconv"
	"trybeego/models"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) GetAll() {
	o := orm.NewOrm()
	var users []*models.User
	_, err := o.QueryTable(new(models.User)).All(&users)

	if err != nil {
		c.CustomAbort(500, err.Error())
	}
	c.Data["Users"] = users
	c.ServeJSON()
}

func (c *UserController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)

	o := orm.NewOrm()
	user := models.User{ID: id}
	err := o.Read(&user)

	if err == orm.ErrNoRows {
		c.CustomAbort(404, "User not found")
	}
	c.Data["json"] = user
	c.ServeJSON()
}

func (c *UserController) Create() {
	//c.Ctx.Input.RequestBody 在 Beego 2.x 有时会返回空，尤其是在请求体已经被读取过的情况下。
	//使用 io.ReadAll(c.Ctx.Request.Body) 可以直接从 HTTP 请求里读取完整原始 body，保证 json.Unmarshal 可以正常解析。
	body, err := io.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		logs.Error("Read body error: ", err)
		c.CustomAbort(400, "Cannot read request body")
		return
	}
	logs.Info("Request body raw: %s", string(body))
	logs.Info("Request Method: %s", c.Ctx.Request.Method)
	logs.Info("Request Header: %+v", c.Ctx.Request.Header)

	var user models.User
	if err := json.Unmarshal(body, &user); err != nil {
		logs.Error("RunSyncdb error: ", err)
		c.CustomAbort(400, "Invalid input")
	}
	o := orm.NewOrm()
	_, err = o.Insert(&user)
	if err != nil {
		c.CustomAbort(500, err.Error())
	}
	c.Data["json"] = user
	c.ServeJSON()
}

func (c *UserController) Update() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)

	var user models.User
	body, err := io.ReadAll(c.Ctx.Request.Body)
	if err := json.Unmarshal(body, &user); err != nil {
		c.CustomAbort(400, "Invalid input")
	}
	user.ID = id

	o := orm.NewOrm()
	_, err = o.Update(&user)
	if err != nil {
		c.CustomAbort(500, err.Error())
	}
	c.Data["json"] = user
	c.ServeJSON()
}

func (c *UserController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)

	o := orm.NewOrm()
	_, err := o.Delete(&models.User{ID: id})
	if err != nil {
		c.CustomAbort(500, err.Error())
	}
	c.Data["json"] = map[string]string{"message": "User deleted successfully"}
	c.ServeJSON()
}
