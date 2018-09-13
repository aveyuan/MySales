package controllers

import (
	"sales-project/models"
)

type ClientController struct {
	BaseController
}

func (this *ClientController)Add()  {
	if this.IsPost(){
		name := this.Ctx.Request.PostForm.Get("name")
		phone := this.Ctx.Request.PostForm.Get("phone")
		address := this.Ctx.Request.PostForm.Get("address")
		postid := this.Ctx.Request.PostForm.Get("postid")
		remarks := this.Ctx.Request.PostForm.Get("remarks")
		client := &models.Client{Name:name,Phone:phone,Address:address,Postid:postid,Remarks:remarks,Createtime:this.GetDateTime()}
		if err := client.Add();err !=nil{
			this.Ctx.WriteString("添加失败")
		}else {
			this.Redirect(this.URLFor(".List"),302)
		}
	}else {
		this.Data["pagetitle"]="新增客户"
		this.Layout="public/layout.html"
		this.TplName="client/add.html"
	}
}

func (this *ClientController)List()  {
	client := &models.Client{}
	clients := client.List()
	this.Data["clients"]=clients
	this.Data["pagetitle"]="用户列表"
	this.Layout="public/layout.html"
	this.TplName="client/list.html"
}

func (this *ClientController)Update()  {
	if this.IsPost(){
		id,err := this.GetInt("id")
			if err!=nil{
				this.Ctx.WriteString("获ID有误")
				this.StopRun()
			}
		name := this.GetString("name")
		address := this.GetString("address")
		postid := this.GetString("postid")
		phone := this.GetString("phone")
		remarks := this.GetString("remarks")
		client := &models.Client{Id:id,Name:name,Address:address,Postid:postid,Phone:phone,Remarks:remarks,Updatetime:this.GetDateTime()}
		if err := client.Update();err!=nil{
			this.Ctx.WriteString("更新失败")
		}else{
			this.Redirect(this.URLFor(".List"),302)
		}
	}else {
		id,err := this.GetInt("id")
		if err!=nil{
			this.Ctx.WriteString("数据有误")
		}
		client := &models.Client{Id:id}
		this.Data["client"]=client.IdClinet()
		this.Data["pagetitle"]="修改用户信息页面"
		this.Layout="public/layout.html"
		this.TplName="client/update.html"
	}
}
//不允许删除，功能暂时不用
//func (this *ClientController)Delete()  {
//	id,err := this.GetInt("id")
//	if err!=nil{
//		this.Ctx.WriteString("数据有误")
//	}
//	client := &models.Client{Id:id}
//	if err := client.Delete();err !=nil{
//		this.Ctx.WriteString("删除失败")
//	}else {
//		this.Redirect(this.URLFor(".List"),302)
//	}
//
//}