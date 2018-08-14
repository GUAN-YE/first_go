package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

type User struct {
	Id      int
	Name    string
}

func init() {

	//dbUser := "root"
	//dbName := "root" //数据库名字
	//dbPwd := "123456"
	//dbLink := fmt.Sprintf("%s:%s:3306@/%s?charset=utf8", dbUser, dbPwd, dbName)
	orm.RegisterModel(new(User))
	//orm.RegisterDriver("mysql", orm.DRMySQL)   //数据库连接字符串
	//orm.RegisterDataBase("default", "mysql", dbLink)
}

func AddUser(username string) (int64, error){
	o := orm.NewOrm()
	o.Using("default")
	user := new(User)
	user.Name = username
	fmt.Println(user)
	return o.Insert(user)
}
func Login(username string) (int64, error) {
	o:=orm.NewOrm()
	o.Using("default")
	err := o.Raw("select * from user where name = '"+ username+"'")
	if err == nil{
		return o.Insert(User{Name:username})
	}
	return o.Insert(User{Name:""})
}
