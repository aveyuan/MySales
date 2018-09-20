package main

import (
	_ "sales-project/routers"
	"github.com/astaxie/beego"
	"os"
	"github.com/astaxie/beego/orm"
	"sales-project/models"
)

func main() {
	beego.Run()
}

func init()  {
	models.Init()
	if len(os.Args)>=2 && os.Args[1]=="install"{
		orm.RunSyncdb("default",false,true)
		user := new(models.User)
		user.InstallUser()
	}
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionProvider="file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "./tmp"
}

