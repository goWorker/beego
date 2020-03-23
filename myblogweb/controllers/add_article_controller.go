package controllers

import (
	"fmt"
	"time"
	"myblogweb/models"
)

type AddArticleController struct {
	BaseController
}

func (this *AddArticleController) Get(){
	this.TplName = "write_article.html"
}

func (this *AddArticleController) Post() {
	title := this.GetString("title")
	tags := this.GetString("tags")
	short := this.GetString("short")
	content := this.GetString("content")
	fmt.Printf("title:%s,tags:%s\n", title, tags)

	art := models.Article{0, title, tags, short, content, "Allan", time.Now().Unix()}
	_, err := models.AddArticle(art)

	//返回数据给浏览器
	var response map[string]interface{}
	if err == nil {
		//无误
		response = map[string]interface{}{"code": 1, "message": "ok"}
	} else {
		response = map[string]interface{}{"code": 0, "message": "error"}
	}

	this.Data["json"] = response
	this.ServeJSON()

}