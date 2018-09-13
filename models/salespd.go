package models

import "github.com/astaxie/beego/orm"

type Salespd struct {
	Id int
	Sales	*Sales `orm:"rel(fk)"`
	Product	*Product `orm:"rel(fk)"`
	Unitprice	float32
	Number 		int
	SalesPdname	string
	SalesPdapnu	string
	SalesPddate  string
	SalesPdbatch string
	SalesPdShelflife       string
	SalesPdmanufacturer    string

}

func (this *Salespd)Add()error  {
	o := orm.NewOrm()
	if _,err := o.Insert(this);err !=nil{
		return err
	}
	return nil
}

func (this *Salespd)IdSales(id int)[]*Salespd {
	o := orm.NewOrm()
	var salespds []*Salespd
	o.QueryTable(Salespd{}).Filter("Sales",id).All(&salespds)
	return salespds
}