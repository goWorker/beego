package main

import (
	_ "zibian/mybook/routers"
	_ "zibian/mybook/sysinit"
	"github.com/astaxie/beego"
)

func main() {
	//beego.SetStaticPath("/static", "./static")
	beego.Run()
}

