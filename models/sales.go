package models

import (
	"github.com/astaxie/beego/orm"
)

//销售数据库
//这个数据库与salespd对应两者为一对多关系，即:一个销售单里有多个产品销售内容
//这里的电话，地址又重新保持了一遍是因为客户可能会换地址的情况，本系统没有做地址薄

type Sales struct {
	Id         int
	Client     *Client `orm:"rel(fk)"` //客户
	SalesPhone string //电话
	SalesAddress string //地址
	SalesPostid	string //邮编
	Salespd		[]*Salespd `orm:"reverse(many)"` //产品
	SalesData	string //销售时间
	Status 		string //状态
	Remarks		string //备注
	Express *Express `orm:"reverse(one)"` //快递
}

//添加
func (this *Sales)Add()error  {
	o := orm.NewOrm()
	if _,err := o.Insert(this);err != nil{
		return err
	}
	return nil
}

//列表
func (this *Sales)List()[]*Sales  {
	var saless []*Sales
	o := orm.NewOrm()
	o.QueryTable(Sales{}).All(&saless)
	return saless
}

//查询
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

//获取
func (this *Sales)GetSales()(*Sales,error)  {
	o := orm.NewOrm()
	if err := o.Read(this);err !=nil{
		return nil,err
	}
	return this,nil
}

//状态
func (this *Sales)Cancel()error  {
	o := orm.NewOrm()
	if _,err := o.Update(this,"Status");err !=nil{
		return err
	}
	return nil
}