package main

import (
	_ "officialsummary/routers"
	_ "officialsummary/sysinit"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

