package models

import "github.com/astaxie/beego/orm"

//客户标签
//主要目的是可以通过标签来分类各类客户，目前设想的是一对多

type Tag struct {
	Id int
	Name string //标签名称
	Client []*Client `orm:"reverse(many)"` //客户
}

//添加标签
func (this *Tag)Add() error  {
	o := orm.NewOrm()
	if _,err := o.Insert(this);err != nil{
		return err
	}
	return nil
}

//更新标签
func (this *Tag)Update()error  {
	o := orm.NewOrm()
	if _,err := o.Update(this,"Name");err !=nil{
		return err
	}
	return nil
}

//标签列表
func (this *Tag)List()[]*Tag  {
	o := orm.NewOrm()
	var tag []*Tag
	o.QueryTable(Tag{}).OrderBy("-Id").All(&tag)
	return tag
}

//通过id返回标签
func (this *Tag)Idtag()*Tag  {
	o := orm.NewOrm()
	o.Read(this)
	return this
}