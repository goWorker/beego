package main

import (
	_ "newblog/routers"
	"github.com/astaxie/beego"
	"newblog/utils"
)

func main() {
	beego.SetStaticPath("/static", "./static")

	utils.InitMysql()

	beego.Run()
}

