package controllers

import (
	"sales-project/models"
	"fmt"
)

type ProductController struct {
	BaseController
}

func (this *ProductController)Add()  {
	this.IsLogin()
	if this.IsPost(){
		name := this.Ctx.Request.PostForm.Get("name")
		code := this.Ctx.Request.PostForm.Get("code")
		approvalnumber := this.Ctx.Request.PostForm.Get("approvalnumber")
		productiondate := this.Ctx.Request.PostForm.Get("productiondate")
		productionbatch := this.Ctx.Request.PostForm.Get("productionbatch")
		shelflife := this.Ctx.Request.PostForm.Get("shelflife")
		manufacturer := this.Ctx.Request.PostForm.Get("manufacturer")
		product := &models.Product{Name:name,Code:code,Approvalnumber:approvalnumber,ProductionDate:productiondate,Productionbatch:productionbatch,ShelfLife:shelflife,Manufacturer:manufacturer,Createtime:this.GetDateTime()}
		if err := product.Add();err !=nil{
			this.Ctx.WriteString("添加失败")
		}else {
			this.Redirect(this.URLFor(".List"),302)

		}
	}else {
		this.Data["pagetitle"]="添加产品"
		this.Layout="public/layout.html"
		this.TplName="product/add.html"
	}
}

func (this *ProductController)List()  {
	this.IsLogin()
	product := &models.Product{}
	products := product.List()
	this.Data["pagetitle"]="产品显示页面"
	this.Data["products"]=products
	this.Layout="public/layout.html"
	this.TplName="product/list.html"
}

func (this *ProductController)Update()  {
	if this.IsPost(){
		id,err := this.GetInt("id")
			if err !=nil{
				fmt.Println("数据有误")
				this.StopRun()
			}
		//查询数据库里面有没有数据
		name := this.GetString("name")
		code := this.GetString("code")
		approvalnumber := this.GetString("approvalnumber")
		productiondate := this.GetString("productiondate")
		productionbatch := this.GetString("productionbatch")
		shelflife := this.GetString("shelflife")
		manufacturer := this.GetString("manufacturer")
		product := &models.Product{Id:id,Name:name,Code:code,Approvalnumber:approvalnumber,Productionbatch:productionbatch,ProductionDate:productiondate,ShelfLife:shelflife,Manufacturer:manufacturer,Updatetime:this.GetDateTime()}
		if err := product.Update();err !=nil{
			this.Ctx.WriteString("更新失败")
		}
		this.Redirect(this.URLFor(".List"),302)
	}else {
		id,err := this.GetInt("id")
		if err!=nil{
			this.Ctx.WriteString("数据有误")
		}
		product := &models.Product{Id:id}
		this.Data["product"]=product.IdProduct()
		this.Data["pagetitle"]="修改用户信息页面"
		this.Layout="public/layout.html"
		this.TplName="product/update.html"
	}
}
