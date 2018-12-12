package models

import (
	"github.com/astaxie/beego"
	"net/url"
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
	_"github.com/mattn/go-sqlite3"
	"strings"
	"os"
)

func Init()  {
	beego.Info("正在初始化数据库配置.")
	adapter := beego.AppConfig.String("db_adapter")
	prefix := beego.AppConfig.String("db.prefix")
	//如果配置的是mysql
	if strings.EqualFold(adapter, "mysql") {
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
		orm.RegisterDriver("default", orm.DRMySQL)
	}else if strings.EqualFold(adapter, "sqlite3"){
		orm.RegisterDriver("sqlite",orm.DRSqlite)
		orm.RegisterDataBase("default","sqlite3","sales_data.db")
	}else {
		beego.Error("不支持的数据库类型.")
		os.Exit(1)
	}
	orm.RegisterModelWithPrefix(prefix,new(Tag),new(Client),new(Product),new(Sales),new(User),new(Salespd),new(Express))
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
	beego.Info("数据库初始化完成.")

	if len(os.Args)>=2 && os.Args[1]=="install"{
		orm.RunSyncdb("default",false,true)
		user := new(User)
		user.InstallUser()
	}
}
