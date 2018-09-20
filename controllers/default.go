package controllers

type MainController struct {
	BaseController
}

func (this *MainController) Get() {
	this.IsLogin()
	this.TplName = "public/layout.html"
}
