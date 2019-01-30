package models

import (
	"github.com/astaxie/beego/orm"
)

//快递表
//用于存放快递信息
type Express struct {
	Id int
	Excmp string //快递公司
	Exnum int64 //快递号码
	Remark string //备注
	Sales *Sales `orm:"rel(one)"` //与销售单一对一关联
}

//新增快递
func (this *Express)Add()error  {
	o := orm.NewOrm()
	if _,err:=o.Insert(this);err !=nil{
		return err
	}
	return nil
}

//通过销售单的id来查询到快递单号
func (this *Express)ExSales(id int)*Express  {
	o := orm.NewOrm()
	var ex Express
	o.QueryTable(Express{}).Filter("Sales",id).All(&ex)
	return &ex
}
