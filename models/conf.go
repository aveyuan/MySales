package models

import 	"github.com/astaxie/beego/orm"

//设置基本信息，主要用户打印订单使用

type Conf struct {
	Id int
	Compony string //公司名称
	Address int64 //公司地址
	Phone string //公司电话
	Sender string //发货人
}

//信息保存
func (this *Conf)Add()error  {
	o := orm.NewOrm()
	if _,err:=o.Insert(this);err !=nil{
		return err
	}
	return nil
}

//产品更新
func (this *Conf)Update()error  {
	o := orm.NewOrm()
	if _,err := o.Update(this,"Compony","Address","Phone","Sender");err!=nil{
		return err
	}
	return nil
}