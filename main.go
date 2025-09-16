package main

import (
	"fmt"
	// 路由的关联
	_ "trybeego/routers"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"

	// 👇 这里必须导入 MySQL 驱动
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// Read config
	config := beego.AppConfig
	driver, _ := config.String("db_driver")
	user, _ := config.String("db_user")
	password, _ := config.String("db_password")
	host, _ := config.String("db_host")
	port, _ := config.String("db_port")
	name, _ := config.String("db_name")

	// 数据库连接参数
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user, password, host, port, name)

	// Register database
	orm.RegisterDataBase("default", driver, dsn)

	// 表初始化
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		logs.Error("RunSyncdb error: ", err)
	}
}

func main() {
	beego.Run()
}
