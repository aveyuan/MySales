package models

import (
	"github.com/astaxie/beego/orm"
)

type Express struct {
	Id int
	Excmp string
	Exnum int64
	Remark string
	Sales *Sales `orm:"rel(one)"`
}

func (this *Express)Add()error  {
	o := orm.NewOrm()
	if _,err:=o.Insert(this);err !=nil{
		return err
	}
	return nil
}

func (this *Express)ExSales(id int)*Express  {
	o := orm.NewOrm()
	var ex Express
	o.QueryTable(Express{}).Filter("Sales",id).All(&ex)
	return &ex
}
