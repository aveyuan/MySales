package models

import (
	"github.com/astaxie/beego/orm"
)

//产品基本数据设定

type Product struct {
	Id              int
	Name            string
	Code            string
	Approvalnumber	string
	ProductionDate  string
	Productionbatch string
	ShelfLife       string
	Manufacturer    string
	Createtime 		string
	Updatetime 		string
	Salespd			[]*Salespd `orm:"reverse(many)"`
}

func (this *Product)Add()error  {
	o := orm.NewOrm()
	if _,err := o.Insert(this);err!=nil{
		return err
	}
	return nil
}

func (this *Product)List()[]*Product  {
	var products []*Product
	o := orm.NewOrm()
	o.QueryTable(Product{}).All(&products)
	return products
}

func (this *Product)IdProduct()*Product {
	o := orm.NewOrm()
	o.Read(this)
	return this
}

func (this *Product)Update()error  {
	o := orm.NewOrm()
	if _,err := o.Update(this,"Name","Code","Approvalnumber","ProductionDate","Productionbatch","ShelfLife","Manufacturer","Updatetime");err!=nil{
		return err
	}
	return nil
}