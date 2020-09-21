package main

import (
	"officialsummary/controllers"
	_ "officialsummary/routers"
	_ "officialsummary/sysinit"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/static", "./static")
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}

