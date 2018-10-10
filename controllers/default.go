package controllers

import "sales-project/models"

type MainController struct {
	BaseController
}

func (this *MainController) Get() {
	this.IsLogin()
	//作为默认首页，显示客户数量，产品数量，销售笔数
	//客户数量获取
	client := &models.Client{}
	clints := client.List()
	this.Data["clients"]=len(clints)
	//产品数量获取
	product := &models.Product{}
	products := product.List()
	this.Data["products"]=len(products)
	//销售笔数
	sales := &models.Sales{}
	saless := sales.List()
	this.Data["saless"]=len(saless)
	//一共售卖
	salespd := &models.Salespd{}
	salespds := salespd.List()
	var sum float32
	for _,v := range salespds{
		sum+=v.Totail
	}
	this.Data["sum"]=sum
	this.Layout="public/layout.html"
	this.TplName = "index.html"

}
