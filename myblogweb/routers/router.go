package routers

import (
	"myblogweb/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/register", &controllers.RegisterController{})
	//登录
	beego.Router("/login", &controllers.LoginController{})

	//写文章
	beego.Router("/article/add", &controllers.AddArticleController{})
	//显示文章内容
	beego.Router("/article/:id", &controllers.ShowArticleController{})
	//更新文章
	beego.Router("article/update", &controllers.UpdateArticleController{})
	beego.Router("/article/delete", &controllers.DeleteArticleController{})
	beego.Router("/tags", &controllers.TagsController{})
	beego.Router("/album", &controllers.AlbumController{})

	//文件上传
	beego.Router("/upload", &controllers.UploadController{})
	beego.Router("/aboutme", &controllers.AboutMeController{})
}