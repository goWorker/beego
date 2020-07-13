package main

import (
	_ "officialsummary/routers"
	_ "officialsummary/sysinit"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/static", "./static")
	beego.Run()
}

