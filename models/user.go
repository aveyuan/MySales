package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"sales-project/libs"
	"github.com/astaxie/beego"
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


func (this *User)InstallUser()  {
	o := orm.NewOrm()
	pass := libs.Passwords("123456")
	user := User{Username:"admin",Password:pass,Nikename:"管理员"}
	//先查询是否有用户存在
	if err := o.Read(&user,"Username");err !=nil{
		//没有查询到那就添加用户
		if _,err := o.Insert(&user);err !=nil{
			beego.Info("添加用户失败")
		}else {
			beego.Info("\n请记住用户名:admin\n密码:123456\n请及时修改用户名和密码，祝您使用愉快！")
		}
	}else {
		beego.Info("\n用户已经存在，跳过")
	}
}