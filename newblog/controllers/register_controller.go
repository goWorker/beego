package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"newblog/models"
	"newblog/utils"
	"time"
)

type RegisterController struct{
	beego.Controller
}

func (this *RegisterController) Get(){
	logs.Info("this is get")
	this.TplName = "register.html"
}

func (this *RegisterController) Post(){
	logs.Info("this is post")
	username := this.GetString("username")
	password := this.GetString("password")
	repassword := this.GetString("repassword")

	fmt.Println(username,password,repassword)
	//logs.Info(username,password,repassword)

	id := models.QueryUserWithUsername(username)
	fmt.Println("id:", id)

	if id > 0{
		this.Data["json"] = map[string]interface{}{"code":0,"message":"用户名已经存在"}
		fmt.Println("用户名已经存在")
		this.ServeJSON()
		return
	}
	fmt.Printf("md5前：",password)
	password = utils.MD5(password)
	fmt.Println("md5后：",password)
	fmt.Printf("username is:",username)
	user := models.User{0,username,password,0,time.Now().Unix()}

	_,err :=  models.InsertUser(user)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code":0,"message":"注册失败"}
	}else{
		this.Data["json"] = map[string]interface{}{"code":1,"message":"注册成功"}

	}
	this.ServeJSON()
	}