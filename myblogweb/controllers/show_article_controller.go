package controllers

import (
	"fmt"
	"myblogweb/models"
	"myblogweb/utils"
	"strconv"
)

type ShowArticleController struct{
	BaseController
}

func (this *ShowArticleController) Get(){
	idStr := this.Ctx.Input.Param(":id")
	id,_ := strconv.Atoi(idStr)
	fmt.Println("ID: ",id)

	art := models.QueryArticleWithId(id)
	this.Data["Title"] = art.Title
	this.Data["Content"] = utils.SwitchMarkdownToHtml(art.Content)
	this.TplName="show_article.html"
}