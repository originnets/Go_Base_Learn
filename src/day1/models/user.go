package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//插入
func inserUser() {
	o := orm.NewOrm()
	user := User{}
	user.Name = "huo"
	id, err := o.Insert(&user)
	if err != nil {
		beego.Info("insert err")
		return
	}
	beego.Info("insert success, id = ", id)
}


//查询
func queryUser() {
	o := orm.NewOrm()
	user := User{Id:1}
	err := o.Read(&user)
	if err != nil{
		beego.Info("query is err")
		return
	}
	beego.Info("query success, user = ", user)
}

//更新
func updateUser() {
	o := orm.NewOrm()
	user := User{Id:1, Name:"GO"}
	_, err := o.Update(&user)
	if err != nil {
		beego.Info("update err")
		return
	}
	beego.Info("update success")
}

//删除
func deleteUser() {
	o := orm.NewOrm()
	user := User{Id:2}
	_, err := o.Delete(&user)
	if err != nil {
		beego.Error("delete err")
		return
	}
	beego.Error("delete success")
}

//多表关联插入
func insertoder() {
	o := orm.NewOrm()
	order := User_order{}
	order.Order_data = "this is GO"

	user := User{Id:2}
	order.User = &user

	id, err := o.Insert(&order)
	if err != nil {
		beego.Info("insert err")
		return
	}
	beego.Info("insert success id = ", id)
}

//多表关联查询
func queryOrder() {
	var orders []*User_order
	o := orm.NewOrm()
	qs := o.QueryTable("User_order")
	_, err := qs.Filter("user__id", 1).All(&orders) //把查询的结果放到orders中 //user__id 表示user.id 这个表示查询user.id =1 的order中有几个
	if err != nil {
		beego.Error("query err")
		return
	}
	for _, data := range orders{
		beego.Info("query roder success, order = ", data)
	}
	fmt.Scan("test")
}

