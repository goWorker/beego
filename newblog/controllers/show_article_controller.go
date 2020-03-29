package controllers

import (
	"fmt"
	"newblog/models"
	"strconv"
)

type ShowArticleController struct {
	BaseController
}

func (this *ShowArticleController) Get(){
	idStr := this.Ctx.Input.Param(":id")
	id,_ := strconv.Atoi(idStr)
	fmt.Println("id:",id)
	art := models.QueryArticleWithId(id)
	this.Data["Title"] = art.Title
	this.Data["Content"] = art.Content
	this.TplName = "show_article.html"
}