package controllers

import (
	"github.com/astaxie/beego"
	"zibian/mybook/models"
)

type HomeController struct{
	BaseController
}

func (c *HomeController)Index() {
	if cates,err := new(models.Category).GetCates(-1,1);err==nil{
		c.Data["Cates"] = cates
	}else{
		beego.Error(err.Error())
	}
	c.TplName = "home/list.html"
}