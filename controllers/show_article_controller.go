/**
 * @Author: caoduanxi
 * @Date: 2021/3/17 10:41
 * @Motto: Keep thinking, keep coding!
 */

package controllers

import (
	"myblog/models"
	"myblog/utils"
	"strconv"
)

type ShowArticleController struct {
	BaseController
}

func (this *ShowArticleController) Get() {
	idStr := this.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	// 获取文章信息
	article := models.QueryArticleWithId(id)
	this.Data["Title"] = article.Title
	//this.Data["Content"] = article.Content
	// 转化为markdown形式显示
	this.Data["Content"] = utils.SwitchMarkdownToHtml(article.Content)
	this.TplName = "show_article.html"
}
