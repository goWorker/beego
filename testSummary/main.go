package main

import (
	_ "testSummary/routers"
	_ "testSummary/sysinit"
	"github.com/astaxie/beego"
)

func main() {
	//beego.SetStaticPath("/static", "./static")
	beego.Run()
}

