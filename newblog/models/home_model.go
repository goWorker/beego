package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
)

type HomeFooterPageCode struct {
	HasPre bool
	HasNext bool
	ShowPage string
	PreLink string
	NextLink string
}

func ConfigHomeFooterPageCode(page int) HomeFooterPageCode {
	pageCode := HomeFooterPageCode{}
	num := GetArticleRowsNum()

	pageRow,_ := beego.AppConfig.Int("articleListPageNum")

	allPageNum := (num-1)/pageRow + 1
	pageCode.ShowPage = fmt.Sprintf("%d/%d",page,allPageNum)

	if page < 1 {
		pageCode.HasPre = false
	}else {
		pageCode.HasPre = true
	}
	if page >= allPageNum{
		pageCode.HasNext = false
	}else {
		pageCode.HasNext = true
	}
	pageCode.PreLink = "/?page=" + strconv.Itoa(page-1)
	pageCode.NextLink = "/?page=" + strconv.Itoa(page+1)
	return pageCode
}