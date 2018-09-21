package models

import (
	"github.com/astaxie/beego/orm"
)

//客户

type Client struct {
	Id      int
	Name 	string
	Phone   string
	Address string
	Postid	string
	Remarks string
	Createtime string
	Updatetime string
	Sales []*Sales `orm:"reverse(many)"`
}

func (this *Client)Add()error  {
	o := orm.NewOrm()
	if _,err := o.Insert(this);err != nil{
		return err
	}
	return nil
}

func (this *Client)List()[]*Client  {
	o := orm.NewOrm()
	var clients []*Client
	o.QueryTable(Client{}).All(&clients)
	return clients
}

func (this *Client)ListLimit(limit,page int)[]*Client  {
	o := orm.NewOrm()
	var clients []*Client
	o.QueryTable(Client{}).Limit(limit,(page-1)*limit).All(&clients)
	return clients
}

func (this *Client)IdClinet()*Client {
	o := orm.NewOrm()
	o.Read(this)
	return this
}

func (this *Client)Update()error  {
	o :=orm.NewOrm()
	if _,err := o.Update(this,"Name","Phone","Address","Postid","Remarks","Updatetime");err != nil{
		return err
	}
	return nil
}

//不允许删除，功能暂时不用
//func (this *Client)Delete()error  {
//	o := orm.NewOrm()
//	if _,err := o.Delete(this);err !=nil{
//		return err
//	}
//	return nil
//}