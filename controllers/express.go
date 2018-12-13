package controllers

import (
	"sales-project/models"
	"strconv"
)

type ExpressController struct {
	BaseController
}

func (this *ExpressController)Add()  {
	if this.IsPost(){
		salesid,_ := this.GetInt("id")
		sales := &models.Sales{Id:salesid}
		excmp := this.GetString("excmp")
		exnum,_ := this.GetInt64("exnum")
		remark:= this.GetString("remark")
		express := &models.Express{Excmp:excmp,Exnum:exnum,Remark:remark,Sales:sales}
		if err:= express.Add();err !=nil{
			this.Ctx.WriteString("添加快递单号失败")
		}
		this.Redirect(this.URLFor("SalesController.Detail")+"?id="+strconv.Itoa(salesid),302)

	}else {
		salesid,_ := this.GetInt("id")
		this.Data["salesid"]=salesid
		this.Xsrf()
		this.Layout="public/layout.html"
		this.TplName="express/add.html"
	}
}
