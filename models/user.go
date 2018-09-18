package models

import (
	"github.com/astaxie/beego/orm"
)

//用户表

type User struct {
	Id int
	Username string
	Password string
	Nikename string
}

func (this *User)Login()error  {
	o := orm.NewOrm()
	if err :=o.Read(this,"Username","Password");err !=nil{
		return err
	}
	return nil
}

func (this *User)Add()error  {
	o := orm.NewOrm()
	if _,err := o.Insert(this);err != nil{
		return err
	}
	return nil
}