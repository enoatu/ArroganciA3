package main

import (
	_ "arrogancia/routers"
	"arrogancia/schedules"
	// "github.com/davecgh/go-spew/spew"

	"github.com/astaxie/beego/toolbox"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	sqlconn, _ := beego.AppConfig.String("sqlconn")
	orm.RegisterDataBase("default", "mysql", sqlconn)
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	toolbox.AddTask("collectTask", schedules.GetCollectTask())
	toolbox.StartTask()
	defer toolbox.StopTask()

	beego.Run()
}
