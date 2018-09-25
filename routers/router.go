package routers

import (
	"sales-project/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/client/add",&controllers.ClientController{},"*:Add")
    beego.Router("/client/list",&controllers.ClientController{},"get,post:List")
    beego.Router("/client/update",&controllers.ClientController{},"*:Update")
    //beego.Router("/client/delete",&controllers.ClientController{},"*:Delete")

    beego.Router("/product/add",&controllers.ProductController{},"*:Add")
    beego.Router("/product/list",&controllers.ProductController{},"get:List")
    beego.Router("/product/update",&controllers.ProductController{},"*:Update")

    beego.Router("/sales/add",&controllers.SalesController{},"*:Add")
    beego.Router("/sales/list",&controllers.SalesController{},"*:List")
    beego.Router("/sales/detail",&controllers.SalesController{},"*:Detail")
    beego.Router("/sales/status",&controllers.SalesController{},"*:Upstatus")
    beego.Router("/login",&controllers.LoginController{},"*:Login")
    beego.Router("/reg",&controllers.LoginController{},"*:Reg")
    beego.Router("/updatepass",&controllers.LoginController{},"*:UpdatePass")
    beego.Router("/updatenike",&controllers.LoginController{},"*:UpdateNike")
    beego.Router("/logout",&controllers.LoginController{},"*:Logout")
}
