package main

import (
    _ "api/routers"

    beego "github.com/beego/beego/v2/server/web"
    "github.com/beego/beego/v2/client/orm"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    sqlconn, _ := beego.AppConfig.String("sqlconn")
    orm.RegisterDataBase("default", "mysql", sqlconn)
    if beego.BConfig.RunMode == "dev" {
        beego.BConfig.WebConfig.DirectoryIndex = true
        beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
    }
    beego.Run()
}

