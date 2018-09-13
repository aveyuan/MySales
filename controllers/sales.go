package controllers

import (
	"sales-project/models"
	"strconv"
	"fmt"
)

type SalesController struct {
	BaseController
}

func (this *SalesController)Add()  {
	if this.IsPost(){
		clientid,_ := this.GetInt("client")
		idclient := &models.Client{Id:clientid}
		//查询clinet
		client := idclient.IdClinet()
		//先保存订单信息
		sals := &models.Sales{Client:client,SalesPhone:client.Phone,SalesPostid:client.Postid,SalesAddress:client.Address,SalesData:this.GetDateTime(),Status:"成功"}
		if err := sals.Add();err != nil{
			this.Ctx.WriteString("添加销售失败")
		}else {
			//首先获取到多个信息，因为要多行保存，肯定三个都是匹配的
			slunitprice := this.GetStrings("unitprice")
			slnumber := this.GetStrings("number")
			slproductid := this.GetStrings("product")
			//三个统一的key不会错的，利用key来一个一个保存
			for k,_ := range slnumber{
				//如果产品的id是空的，那就跳过
				if slproductid[k]==""{
					continue
				}
				//获取单价，再转换为float
				sunitprice := slunitprice[k]
				unitprice64,_ := strconv.ParseFloat(sunitprice,32)
				unitprice := float32(unitprice64)
				//获取数量，再转换为int
				snumber := slnumber[k]
				number,_ := strconv.Atoi(snumber)
				//获取产品id，再转换为int
				sproductid := slproductid[k]
				productid,_ := strconv.Atoi(sproductid)
				idproduct := &models.Product{Id:productid}
				//查询到产品
				product := idproduct.IdProduct()
				//将产品信息保存在销售产品pd表
				salespd := &models.Salespd{Product:product,Sales:sals,SalesPdbatch:product.Productionbatch,SalesPdapnu:product.Approvalnumber,SalesPddate:product.ProductionDate,SalesPdmanufacturer:product.Manufacturer,SalesPdShelflife:product.ShelfLife,SalesPdname:product.Name,Unitprice:unitprice,Number:number}
				if err := salespd.Add();err != nil{
					this.Ctx.WriteString("添加产品失败")
				}else {
					this.Redirect(this.URLFor(".List"),302)
				}
			}//结束了保存产品
		}//结束了销售单保存失败
		this.Ctx.WriteString("添加产品成功") //最后一个总的添加，表示订单和订单里面的产品都保存ok
	}else {
		client := &models.Client{}
		clients := client.List()
		product := &models.Product{}
		products := product.List()
		this.Data["clients"]=clients
		this.Data["products"]=products
		this.Data["pagetitle"]="新增销售"
		this.Layout="public/layout.html"
		this.TplName="sales/add.html"
	}
}

func (this *SalesController)List()  {
	sales := new(models.Sales)
	saless := sales.List()
	//只能在这里显示一下简单的信息了
	list := make([]map[string]interface{},len(saless))
	for k,v := range saless{
		//先获取到客户信息
		row := make(map[string]interface{})
		client := v.Client.IdClinet()
		row["id"]=v.Id
		row["clientname"]=client.Name
		row["address"]=v.SalesAddress
		row["postid"]=v.SalesPostid
		row["phone"]=v.SalesPhone
		row["date"]=v.SalesData
		row["status"]=v.Status
		salespd := new(models.Salespd)
		row["salespd"]=salespd.IdSales(v.Id)
		list[k]=row
	}
	this.Data["sales"]=list
	this.Data["pagetitle"]="显示销售单"
	this.Layout="public/layout.html"
	this.TplName="sales/list.html"
}

func (this *SalesController)Detail()  {
	//获取传送过来的slaesid并取得sales
	salesid,_ := this.GetInt("id")
	ssales := &models.Sales{Id:salesid}
	sales,err := ssales.GetSales()
	//如果在没有找到数据则有误，友好的提示，不用抛出异常
	if err !=nil{
		this.Ctx.WriteString("数据有误")
		this.StopRun()
	}
	fmt.Println(sales.SalesAddress)
	//重组数据
	list := make([]interface{},0)
	row := make(map[string]interface{})
	client := sales.Client.IdClinet()
	row["id"]=sales.Id
	row["clientname"]=client.Name
	fmt.Println(client.Name)
	row["address"]=sales.SalesAddress
	row["postid"]=sales.SalesPostid
	row["phone"]=sales.SalesPhone
	row["date"]=sales.SalesData
	salespd := new(models.Salespd)
	row["salespd"]=salespd.IdSales(salesid)
	list = append(list, row)
	this.Data["pagetitle"]="订单详情"
	this.Data["detail"]=list
	this.Layout="public/layout.html"
	this.TplName="sales/detail.html"
}

func (this *SalesController)Upstatus()  {
	salesid,_ := this.GetInt("id")
	ssales := &models.Sales{Id:salesid,Status:"取消"}
	if err := ssales.Cancel();err !=nil{
		this.Ctx.WriteString("更新失败")
	}
	this.Redirect(this.URLFor(".List"),302)
}