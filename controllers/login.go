package controllers

import (
	"sales-project/models"
	"sales-project/libs"
	"github.com/astaxie/beego"
)

type LoginController struct {
	BaseController
}

func (this *LoginController)Login()  {
	if this.IsPost(){
		user := this.GetString("username")
		pass := this.GetString("password")
		password := libs.Passwords(pass)
		users := &models.User{Username:user,Password:password}
		if err := users.Login();err !=nil{
			this.Ctx.WriteString("验证失败")
			this.StopRun()
		}
		this.SetSession("username",user)
		this.Redirect(beego.URLFor("MainController.Get"),302)
	}else {
		this.Xsrf()
		this.Data["pagetitle"]="登录系统"
		this.TplName="login/index.html"
	}
}

func (this *LoginController)Reg()  {
	if this.IsPost(){
		user := this.GetString("username")
		pass1 := this.GetString("password1")
		pass2 := this.GetString("password2")
		nikename := this.GetString("nikename")
		if pass1 != pass2{
			this.Ctx.WriteString("两次密码不一致")
			this.StopRun()
		}
		password := libs.Passwords(pass1) //使用libs库加密
		users := &models.User{Username:user,Password:password,Nikename:nikename}
		if err := users.Add();err != nil{
			this.Ctx.WriteString("注册失败")
		}
		this.Ctx.WriteString("注册成功")
	}else {
		this.Xsrf()
		this.Data["pagetitle"]="注册账号"
		this.TplName="login/reg.html"
	}
}

func (this *LoginController)UpdatePass()  {
	if this.IsPost(){
		userid,err := this.GetInt("id")
		if err !=nil{
			this.Ctx.WriteString("获取参数有误")
			this.StopRun()
		}
		pass1 := this.GetString("password1")
		pass2 := this.GetString("password2")
		if pass1 != pass2{
			this.Ctx.WriteString("两次密码不一致")
			this.StopRun()
		}
		password := libs.Passwords(pass1) //使用libs库加密
		user := &models.User{Id:userid,Password:password}
		if err := user.UpdatePassword();err != nil{
			this.Ctx.WriteString("修改密码失败")
		}
		this.Ctx.WriteString("修改密码成功")

	}else {
		this.Data["pagetitle"]="修改密码"
		user := &models.User{}
		this.Xsrf()
		this.Data["user"]=user.List()
		this.TplName="login/update.html"
	}
}

func (this *LoginController)UpdateNike()  {
	if this.IsPost(){
		userid,err := this.GetInt("id")
		if err !=nil{
			this.Ctx.WriteString("获取参数有误")
			this.StopRun()
		}
		nike := this.GetString("nikename")
		user := &models.User{Id:userid,Nikename:nike}
		if err := user.UpdateNikename();err != nil{
			this.Ctx.WriteString("修改昵称失败")
		}
		this.Ctx.WriteString("修改昵称成功")

	}else {
		this.Data["pagetitle"]="修改昵称"
		user := &models.User{}
		this.Xsrf()
		this.Data["user"]=user.List()
		this.TplName="login/updatenike.html"
	}
}

func (this *LoginController)Logout()  {
	this.DestroySession()
	this.Redirect(beego.URLFor("LoginController.Login"),302)
}