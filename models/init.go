package models

import (
	"github.com/astaxie/beego"
	"net/url"
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"

)

func Init()  {
	prefix := beego.AppConfig.String("db.prefix")
	dbhost := beego.AppConfig.String("db.host")
	dbport := beego.AppConfig.String("db.port")
	dbuser := beego.AppConfig.String("db.user")
	dbpassword := beego.AppConfig.String("db.password")
	dbname := beego.AppConfig.String("db.name")
	timezone := beego.AppConfig.String("db.timezone")
	if dbport == "" {
		dbport = "3306"
	}
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	if timezone != "" {
		dsn = dsn + "&loc=" + url.QueryEscape(timezone)
	}
	orm.RegisterDataBase("default", "mysql", dsn)
	orm.RegisterDriver("default",orm.DRMySQL)
	orm.RegisterModelWithPrefix(prefix,new(Client),new(Product),new(Sales),new(User),new(Salespd))
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}
