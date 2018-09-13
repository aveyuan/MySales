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
	}
}

