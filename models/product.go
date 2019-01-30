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

//添加产品
func (this *Product)Add()error  {
	o := orm.NewOrm()
	if _,err := o.Insert(this);err!=nil{
		return err
	}
	return nil
}

//产品列表
func (this *Product)List()[]*Product  {
	var products []*Product
	o := orm.NewOrm()
	o.QueryTable(Product{}).All(&products)
	return products
}

//产品查询列表
func (this *Product)ListLimit(limit,page int,key string)([]*Product,[]*Product)  {
	o := orm.NewOrm()
	var products []*Product
	var num []*Product
	if key == "*"{
		o.QueryTable(Product{}).Limit(limit,(page-1)*limit).OrderBy("-Id").All(&products)
		o.QueryTable(Product{}).All(&num)
	}else {
		con := orm.NewCondition()
		con1 := con.Or("Name__icontains",key).Or("Code__icontains",key).Or("Approvalnumber__icontains",key).Or("Manufacturer__icontains",key)
		o.QueryTable(Product{}).SetCond(con1).Limit(limit,(page-1)*limit).OrderBy("-Id").All(&products)
		o.QueryTable(Product{}).SetCond(con1).All(&num)
	}
	return products,num
}

//通过id查询到产品
func (this *Product)IdProduct()*Product {
	o := orm.NewOrm()
	o.Read(this)
	return this
}

//产品更新
func (this *Product)Update()error  {
	o := orm.NewOrm()
	if _,err := o.Update(this,"Name","Code","Approvalnumber","ProductionDate","Productionbatch","ShelfLife","Manufacturer","Updatetime");err!=nil{
		return err
	}
	return nil
}