package routers

import (
	"newblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.HomeController{})
    beego.Router("/register",&controllers.RegisterController{})

    beego.Router("/login",&controllers.LoginController{})
	beego.Router("/exit", &controllers.ExitController{})
	//写文章
	beego.Router("/article/add", &controllers.AddArticleController{})
    beego.Router("/article/:id",&controllers.ShowArticleController{})
    beego.Router("/article/update",&controllers.UpdateArticleController{})
	beego.Router("/article/delete",&controllers.DeleteArticleController{})
	beego.Router("/tags", &controllers.TagsController{})
	beego.Router("/album", &controllers.AlbumController{})
	beego.Router("/upload", &controllers.UploadController{})
	beego.Router("/aboutme", &controllers.AboutMeController{})
}

//"get:ShowRegister;post:HandleRegister")
//beego.Router("/login", &controllers.LoginController{}, "get:ShowLogin;post:HandleLogin")
