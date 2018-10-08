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
	Tag	*Tag `orm:"rel(fk)"`
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

func (this *Client)TagGet()*Client  {
	o := orm.NewOrm()
	o.Read(this)
	if this.Tag != nil{
		o.Read(this.Tag)
	}
	return this
}

//一共返回两个变量，一个是显示当前的，另外一个是没有分页的，可以很好的返回总页数
func (this *Client)ListLimit(limit,page int,key string)([]*Client,[]*Client)  {
	o := orm.NewOrm()
	var clients []*Client
	var num []*Client
	if key=="*"{
		o.QueryTable(Client{}).Limit(limit,(page-1)*limit).OrderBy("-Id").All(&clients)
		o.QueryTable(Client{}).All(&num)
	}else {
		con := orm.NewCondition()
		con1 := con.Or("Name__icontains",key).Or("Phone__icontains",key).Or("Address__icontains",key)
		o.QueryTable(Client{}).SetCond(con1).Limit(limit,(page-1)*limit).OrderBy("-Id").All(&clients)
		o.QueryTable(Client{}).SetCond(con1).All(&num)

	}
	return clients,num
}

func (this *Client)IdClinet()*Client {
	o := orm.NewOrm()
	o.Read(this)
	return this
}

func (this *Client)Update()error  {
	o :=orm.NewOrm()
	//tag标签更新暂未实现
	if _,err := o.Update(this,"Id","Name","Phone","Address","Postid","Remarks","Updatetime","Tag");err != nil{
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