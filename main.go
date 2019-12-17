package main

import (
	"github.com/astaxie/beego"
	"sales-project/models"
	_ "sales-project/routers"
)

func main() {
	models.Init()
	beego.Run()
}
