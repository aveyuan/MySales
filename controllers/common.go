package controllers

import (
	"github.com/astaxie/beego"

	"time"
	"html/template"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController)IsPost()bool  {
	return this.Ctx.Request.Method=="POST"
}

func (this *BaseController)GetDateTime()string  {
	return time.Now().Format("2006-01-02 15:04:05")
}

//判断是否登录，如果未登录则强制跳转到登录页面
func (this *BaseController)IsLogin()  {
	if this.GetSession("username") == nil{
		this.Redirect(beego.URLFor("LoginController.Login"),302)

	}
}

func (this *BaseController)Xsrf()  {
	this.Data["xsrfdata"]=template.HTML(this.XSRFFormHTML())
}