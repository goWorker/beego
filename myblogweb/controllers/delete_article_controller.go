package controllers

import (
	"fmt"
	"log"
	"myblogweb/models"
)

type DeleteArticleController struct {
	BaseController
}

func (this *DeleteArticleController) Get(){
	artID, _:= this.GetInt("id")
	fmt.Println("删除 ID：",artID)

	_,err := models.DeleteArticle(artID)
	if err != nil {
		log.Println(err)
	}
	this.Redirect("/",302)
}