package controllers

import (
	"myblog/models"
	"time"
)

type AddArticleController struct {
	BaseController
}

// 跳转到文章页面
func (this *AddArticleController) Get(){
	this.TplName = "write_article.html"
}

// 添加文件逻辑post
func (this *AddArticleController) Post(){
	title := this.GetString("title")
	tags := this.GetString("tags")
	short := this.GetString("short")
	content := this.GetString("content")

	// 转化为Article实体
	article := models.Article{0,title,tags,short,content,"aaroncao",time.Now().Unix()}
	_, err := models.AddArticle(article)
	var response map[string]interface{}
	if err == nil {
		response = map[string]interface{}{"code": 1, "message": "ok"}
	} else {
		response = map[string]interface{}{"code": 0, "message": "error"}
	}

	this.Data["json"] = response
	this.ServeJSON()

}
