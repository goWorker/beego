package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

//每个子类Controller公用方法调用前，都执行一下Prepare方法
//func (c *BaseController) Prepare() {
//	c.Member = models.NewMember() //初始化
//	c.EnableAnonymous = false
//	//从session中获取用户信息
//	if member, ok := c.GetSession(common.SessionName).(models.Member); ok && member.MemberId > 0 {
//		c.Member = &member
//	} else {
//		//如果Cookie中存在登录信息，从cookie中获取用户信息
//		if cookie, ok := c.GetSecureCookie(common.AppKey(), "login"); ok {
//			var remember CookieRemember
//			err := utils.Decode(cookie, &remember)
//			if err == nil {
//				member, err := models.NewMember().Find(remember.MemberId)
//				if err == nil {
//					c.SetMember(*member)
//					c.Member = member
//				}
//			}
//		}
//	}
//	if c.Member.RoleName == "" {
//		c.Member.RoleName = common.Role(c.Member.MemberId)
//	}
//	c.Data["Member"] = c.Member
//	c.Data["BaseUrl"] = c.BaseUrl()
//	c.Data["SITE_NAME"] = "MBOOK"
//	//设置全局配置
//	c.Option = make(map[string]string)
//	c.Option["ENABLED_CAPTCHA"] = "false"
//}