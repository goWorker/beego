package controllers

import (
	"fmt"
	"myblogweb/models"
)

type TagsController struct {
	BaseController
}

func (this * TagsController) Get() {
	tags := models.QueryArticleWithParam("tags")
	fmt.Println(models.HandleTangsListData(tags))
	this.Data["Tags"] = models.HandleTangsListData(tags)
	this.TplName = "tags.html"
}