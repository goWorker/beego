package controllers

type AboutMeController struct {
	BaseController
}

func (c *AboutMeController) Get() {
	c.Data["wechat"] = "微信： Nurruden"
	c.Data["qq"] = "QQ: 382599386"
	c.Data["tel"] = "Tel:12345678901"
	c.TplName = "aboutme.html"
}