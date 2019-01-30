package models

import (
	"github.com/astaxie/beego/orm"
)

//销售数据库
/*
用于存放销售产品的快照,实现打印,查询
 */

type Sales struct {
	Id         int
	Client     *Client `orm:"rel(fk)"`
	SalesPhone string
	SalesAddress string
	SalesPostid	string
	Salespd		[]*Salespd `orm:"reverse(many)"`
	SalesData	string
	Status 		string
	Remarks		string
	Express *Express `orm:"reverse(one)"`
}

func (this *Sales)Add()error  {
	o := orm.NewOrm()
	if _,err := o.Insert(this);err != nil{
		return err
	}
	return nil
}

func (this *Sales)List()[]*Sales  {
	var saless []*Sales
	o := orm.NewOrm()
	o.QueryTable(Sales{}).All(&saless)
	return saless
}

func (this *Sales)ListLimit(limit,page int,key string)([]*Sales,[]*Sales)  {
	var saless []*Sales
	var num []*Sales
	o := orm.NewOrm()
	if key == "*"{
		o.QueryTable(Sales{}).Limit(limit,(page-1)*limit).OrderBy("-Id").All(&saless)
		o.QueryTable(Sales{}).All(&num)
	}else {
		con := orm.NewCondition()
		con1 := con.Or("Id__icontains",key).Or("SalesPhone__icontains",key).Or("SalesAddress__icontains",key)
		//分页数据显示
		o.QueryTable(Sales{}).SetCond(con1).Limit(limit,(page-1)*limit).OrderBy("-Id").OrderBy("-Id").All(&saless)
		//总共数据
		o.QueryTable(Sales{}).SetCond(con1).OrderBy("-Id").All(&num)

	}
	//如果没有查询到，再尝试是否是通过姓名来查询的，这里需要反向查询
	if len(num)==0{
		o.QueryTable(Sales{}).Filter("Client__Name", key).RelatedSel().Limit(limit,(page-1)*limit).OrderBy("-Id").OrderBy("-Id").All(&saless)
		o.QueryTable(Sales{}).Filter("Client__Name", key).OrderBy("-Id").All(&num)
	}
	return saless,num
}

func (this *Sales)GetSales()(*Sales,error)  {
	o := orm.NewOrm()
	if err := o.Read(this);err !=nil{
		return nil,err
	}
	return this,nil
}

func (this *Sales)Cancel()error  {
	o := orm.NewOrm()
	if _,err := o.Update(this,"Status");err !=nil{
		return err
	}
	return nil
}