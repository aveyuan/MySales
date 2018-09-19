package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
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

func (this *User)UpdatePassword()error  {
	o := orm.NewOrm()
	if _,err := o.Update(this,"Password","Nikename");err !=nil{
		return err
	}
	return nil
}

func (this *User)UpdateNikename()error  {
	o := orm.NewOrm()
	if _,err := o.Update(this,"Nikename");err !=nil{
		return err
	}
	return nil
}

func (this *User)List()*User  {
	o := orm.NewOrm()
	user := User{}
	o.QueryTable(User{}).One(&user)
	fmt.Println(user.Nikename)
	return &user
}