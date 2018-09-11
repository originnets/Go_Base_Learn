package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id int	//ID默认是主键,auto
	//IdInfo	int `orm:pk,auto`
	Name string `orm:"size(100)"` //varchar(100)
	User_order []*User_order `orm:"reverse(many)"` // 设置一对多的反向关系
}


type User_order struct {
	Id int
	Order_data string `orm:"size(100)""`
	User *User `orm:"rel(fk)"` //设置一对多关系 , 一个User对应多个User_order
}

func init() {
	//设置连接串
	username := beego.AppConfig.String("username")
	hostname := beego.AppConfig.String("hostname")
	password := beego.AppConfig.String("password")
	dbname := beego.AppConfig.String("dbname")
	port := beego.AppConfig.String("port")

	if port == "" {
		port = "3306"
	}

	conn := username + ":" + password + "@tcp(" + hostname + ":" + port + ")/" + dbname + "?charset=utf8"

	// set default database
	//orm.RegisterDataBase("default", "mysql", "root:!QAZ2wsx@tcp(192.168.1.218:3306)/gotest?charset=utf8", 30)
	orm.RegisterDataBase("default", "mysql", conn, 30)

	// register model
	orm.RegisterModel(new(User), new(User_order)) // 创建表

	orm.Debug = true
	// create table
	orm.RunSyncdb("default", false, false) // true 改成false，如果表存在则会给出提示，如果改成false则不会提示 ， 这句话没有会报主键不存在的错误
}