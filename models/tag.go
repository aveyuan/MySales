package models

import "github.com/astaxie/beego/orm"

type Tag struct {
	Id int
	Name string
	Client []*Client `orm:"reverse(many)"`
}

func (this *Tag)Add() error  {
	o := orm.NewOrm()
	if _,err := o.Insert(this);err != nil{
		return err
	}
	return nil
}

func (this *Tag)Update()error  {
	o := orm.NewOrm()
	if _,err := o.Update(this,"Name");err !=nil{
		return err
	}
	return nil
}

func (this *Tag)List()[]*Tag  {
	o := orm.NewOrm()
	var tag []*Tag
	o.QueryTable(Tag{}).OrderBy("-Id").All(&tag)
	return tag
}

func (this *Tag)Idtag()*Tag  {
	o := orm.NewOrm()
	o.Read(this)
	return this
}