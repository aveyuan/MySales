package routers

import (
	"sales-project/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/client/add",&controllers.ClientController{},"*:Add")
    beego.Router("/client/list",&controllers.ClientController{},"get,post:List")
    beego.Router("/client/update",&controllers.ClientController{},"*:Update")
    //beego.Router("/client/delete",&controllers.ClientController{},"*:Delete")

    beego.Router("/product/add",&controllers.ProductController{},"*:Add")
    beego.Router("/product/list",&controllers.ProductController{},"get,post:List")
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

   beego.Router("/tag/add",&controllers.TagController{},"*:Add")
   beego.Router("/tag/list",&controllers.TagController{},"*:List")
   beego.Router("/tag/update",&controllers.TagController{},"*:Update")

   beego.Router("/express/add",&controllers.ExpressController{},"*:Add")

    //路由过滤
    var FilterUser = func(ctx *context.Context) {
        _, ok := ctx.Input.Session("username").(string)
        if !ok && ctx.Request.RequestURI != "/login" {
            ctx.Redirect(302, "/login")
        }
    }
    beego.InsertFilter("/*",beego.BeforeRouter,FilterUser)
}
