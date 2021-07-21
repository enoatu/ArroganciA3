package main

import (
	_ "arrogancia/routers"
	// _ "arrogancia/tasks"
	// "github.com/davecgh/go-spew/spew"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	sqlconn, _ := beego.AppConfig.String("sqlconn")
	// spew.Dump(sqlconn)
	orm.RegisterDataBase("arrogancia", "mysql", sqlconn)
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
