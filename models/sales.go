package models

import "github.com/astaxie/beego/orm"

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

func (this *Sales)ListLimit(limit,page int)[]*Sales  {
	var saless []*Sales
	o := orm.NewOrm()
	o.QueryTable(Sales{}).Limit(limit,(page-1)*limit).All(&saless)
	return saless
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