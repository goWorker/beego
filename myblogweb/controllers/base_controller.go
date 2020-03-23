package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	IsLogin	bool
	Loginuser	interface{}
}

func (this *BaseController)  Prepare(){
	loginuser := this.GetSession("loginuser")
	fmt.Println("loginuser----->",loginuser)
	if loginuser != nil {
		this.IsLogin = true
		this.Loginuser = loginuser
	}else {
		this.IsLogin = false
	}
	this.Data["IsLogin"] = this.IsLogin
}