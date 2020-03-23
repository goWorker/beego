package routers

import (
	"zibian/mybook/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//beego.Router("/", &controllers.MainController{})
	beego.Router("/", &controllers.HomeController{},"get:Index")
	beego.Router("/explore", &controllers.ExploreController{},"get:Index")
	//beego.Router("/books/:key", &controllers.DocumentController{},"get:Index")
	//beego.Router("/read/:key/:id", &controllers.DocumentController{},"*:Read")
	//beego.Router("/read/:key/search", &controllers.DocumentController{},"post:Search")
	//
	//beego.Router("/api/:key/content/?:id", &controllers.DocumentController{},"*:Content")
	//beego.Router("/api/:key/edit/?:id", &controllers.DocumentController{},"*:Edit")
	//
	//beego.Router("/api/upload", &controllers.DocumentController{},"post:Upload")
	//beego.Router("/api/:key/edit/create", &controllers.DocumentController{},"post:Create")
	//beego.Router("/api/:key/edit/delete", &controllers.DocumentController{},"post:Delete")
	//
	//beego.Router("/search", &controllers.SearchController{},"get:Search")
	//beego.Router("/search/result", &controllers.SearchController{},"get:Result")
	//
	//beego.Router("/login", &controllers.AccountController{},"*:Login")
	//beego.Router("/regist", &controllers.AccountController{},"*:Regist")
	//beego.Router("/logout", &controllers.AccountController{},"*:Logout")
	//beego.Router("/doregist", &controllers.AccountController{},"post:Doregist")
	//
	//beego.Router("/book", &controllers.BookController{},"*:Index")
	//beego.Router("/book/create", &controllers.BookController{},"post:Create")
	//beego.Router("/book/:key/setting", &controllers.BookController{},"*:Setting")
	//beego.Router("/book/setting/upload", &controllers.BookController{},"post:UploadCover")
	//beego.Router("/book/star/:id", &controllers.BookController{},"*:Collection")
	//beego.Router("/book/setting/save", &controllers.BookController{},"post:SaveBook")
	//beego.Router("/book/:key/release", &controllers.BookController{},"post:Release")
	//beego.Router("/book/setting/save", &controllers.BookController{},"post:SaveBook")
	//beego.Router("/book/setting/token", &controllers.BookController{},"post:CreateToken")
	//
	//beego.Router("/user/:username", &controllers.UserController{}, "get:Index")
	//beego.Router("/user/:username/collection", &controllers.UserController{}, "get:Collection")
	//beego.Router("/user/:username/follow", &controllers.UserController{}, "get:Follow")
	//beego.Router("/user/:username/fans", &controllers.UserController{}, "get:Fans")
	//beego.Router("/follow/:uid", &controllers.BaseController{}, "get:SetFollow")
	//beego.Router("/book/score/:id", &controllers.BookController{}, "*:Score")
	//beego.Router("/book/comment/:id", &controllers.BookController{}, "post:Comment")
	//
	//
	//beego.Router("/setting", &controllers.SettingController{}, "*:Index")
	//beego.Router("/setting/upload", &controllers.SettingController{}, "*:Upload")
	//
	//beego.Router("/manager/category", &controllers.ManagerController{}, "post,get:Category")
	//beego.Router("/manager/update-cate", &controllers.ManagerController{}, "get:UpdateCate")
	//beego.Router("/manager/del-cate", &controllers.ManagerController{}, "get:DelCate")
	//beego.Router("/manager/icon-cate", &controllers.ManagerController{}, "post:UpdateCateIcon")
}
