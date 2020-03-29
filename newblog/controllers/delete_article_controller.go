package controllers

import (
	"fmt"
	"log"
	"newblog/models"
)

type DeleteArticleController struct {

	BaseController
}

func (this *DeleteArticleController) Get(){
	artID,_:=this.GetInt("id")
	fmt.Println("Delete id: ",artID)
	_,err := models.DeleteArticle(artID)
	if err != nil {
		log.Println(err)
	}
	this.Redirect("/",302)
}