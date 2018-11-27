package main

import (
	_ "sales-project/routers"
	"github.com/astaxie/beego"
	"sales-project/models"
)

func main() {
	models.Init()
	beego.Run()
}
