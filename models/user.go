package models

import "github.com/beego/beego/v2/client/orm"

type User struct {
	ID    int    `orm:"pk;auto" json:"id"`
	Name  string `orm:"size(100)" json:"name"`
	Email string `orm:"size(100)" json:"email"`
}

func init() {
	// Register model
	orm.RegisterModel(new(User))
}
