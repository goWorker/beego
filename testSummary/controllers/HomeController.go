package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"testSummary/models"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Index() {
	//models.PrintUsers()
	//models.PrintUserByORM()
	if table,err:= new(models.AutoSummary).HomeData();err==nil{
		fmt.Println(table)
		c.Data["Contents"]=table
	}else{
		beego.Error(err.Error())
	}
	c.TplName = "index.html"

}
