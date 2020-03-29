package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Testtable struct {
	Id int
	Name string
}

func init(){
	orm.RegisterDataBase("default","mysql","root:wandl123@tcp(127.0.0.1:3306)/testtable?charset=utf8",30)
	orm.RegisterModel(new(Testtable))
}

func PrintUserByORM(){
	o := orm.NewOrm() //定义orm对象
	user := Testtable{Id:2}
	o.Read(&user)
	fmt.Println(user)
}