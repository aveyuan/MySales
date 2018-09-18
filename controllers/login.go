package controllers

import (
	"sales-project/models"
)

type LoginController struct {
	BaseController
}

func (this *LoginController)Login()  {
	if this.IsPost(){
		user := this.GetString("username")
		pass := this.GetString("password")
		users := &models.User{Username:user,Password:pass}
		if err := users.Login();err !=nil{
			this.Ctx.WriteString("验证失败")
			this.StopRun()
		}
		this.Ctx.WriteString("验证成功")
	}else {
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
		//md5pass := hex.EncodeToString([]byte) 加密明天再看吧
		users := &models.User{Username:user,Password:pass2,Nikename:nikename}
		if err := users.Add();err != nil{
			this.Ctx.WriteString("注册失败")
		}
		this.Ctx.WriteString("注册成功")
	}else {
		this.Data["pagetitle"]="注册账号"
		this.TplName="login/reg.html"
	}
}