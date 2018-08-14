package main

import (
	_ "jjj/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)

func init() {
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbname := beego.AppConfig.String("dbname")
	if dbport == "" {
		dbport = "3306"
	}
	dsn := dbuser + ":" + dbpassword+ "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	//orm.RegisterDriver("mysql",orm.DRMySQL,)
	orm.RegisterDataBase("default", "mysql",dsn)
	//dbhost := "localhost"
	//dbport := "3306"
	//dbname := "default"
	//dbuser := "root"
	//dbpass :="123456"
	//dsn := strings.Join([]string{dbuser, ":", dbpass, "@tcp(", dbhost, ":", dbport, ")/", dbname, "?charset=utf8"}, "")
	//orm.RegisterDataBase("default", "mysql", dsn, 30)
	//dbUser := "root"
	//dbName := "root" //数据库名字
	//dbPwd := "123456"
	//dbLink := fmt.Sprintf("%s:%s:3306@/%s?charset=utf8", dbUser, dbPwd, dbName)
	//
	//orm.RegisterDriver("mysql", orm.DRMySQL)   //数据库连接字符串
	//orm.RegisterDataBase("default", "mysql", dbLink)
}
func main() {
	beego.Run()
}

