package controllers

import (
	"encoding/json"
	"io"
	"strconv"
	"trybeego/models"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/validation"
	beego "github.com/beego/beego/v2/server/web"
)

type UserController struct {
	beego.Controller
}

type User struct {
	Name  string `json:"name"`  // 用户名，必填
	Email string `json:"email"` // 邮箱，必填
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
	var user User

	body := c.Ctx.Input.RequestBody
	if err := json.Unmarshal(body, &user); err != nil {
		c.CustomAbort(400, "Invalid input")
	}

	// 校验字段
	valid := validation.Validation{}
	valid.Required(user.Name, "name")
	valid.MaxSize(user.Name, 20, "name")
	valid.Required(user.Email, "email")
	valid.Email(user.Email, "email")

	if valid.HasErrors() {
		// 遍历错误并返回
		errs := map[string]string{}
		for _, err := range valid.Errors {
			errs[err.Key] = err.Message
		}
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{"error": errs}
		c.ServeJSON()
		return
	}

	o := orm.NewOrm()
	_, err := o.Insert(&models.User{
		Name:  user.Name,
		Email: user.Email,
	})
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
