package models

import (
	"github.com/astaxie/beego/orm"
)

//产品销售表
//这个表存放的是销售时候的产品表，是产品的一个副本内容，因为产品的生产日期会改变，所以，相当于实现了一个快照的功能
type Salespd struct {
	Id int
	Sales	*Sales `orm:"rel(fk)"`
	Product	*Product `orm:"rel(fk)"`
	Unitprice	float32 //单价
	Number 		int //数量
	Totail		float32 //总价
	SalesPdname	string //销售产品名称
	SalesPdapnu	string //销售产品编号
	SalesPddate  string //产品生产日期
	SalesPdbatch string //产品生产批号
	SalesPdShelflife       string //保质期
	SalesPdmanufacturer    string //生产厂商

}

//添加产品
func (this *Salespd)Add()error  {
	o := orm.NewOrm()
	if _,err := o.Insert(this);err !=nil{
		return err
	}
	return nil
}

//根据sales主表id返回其下的产品
func (this *Salespd)IdSales(id int)[]*Salespd {
	o := orm.NewOrm()
	var salespds []*Salespd
	o.QueryTable(Salespd{}).Filter("Sales",id).All(&salespds)
	return salespds
}

//产品销售产品列表
func (this *Salespd)List() []*Salespd {
	o := orm.NewOrm()
	var salespd []*Salespd
	o.QueryTable(Salespd{}).All(&salespd)
	return salespd
}
