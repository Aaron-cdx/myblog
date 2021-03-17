/**
 * @Author: caoduanxi
 * @Date: 2021/3/17 14:31
 * @Motto: Keep thinking, keep coding!
 */

package controllers

import (
	"log"
	"myblog/models"
)

type DeleteArticleController struct {
	BaseController
}

// 删除操作
func (this *DeleteArticleController) Get() {
	articleId, _ := this.GetInt("id")
	_, err := models.DeleteArticle(articleId)
	if err != nil {
		log.Println("delete article error...", err.Error())
	}
	// 重定向操作
	this.Redirect("/", 302)
}
