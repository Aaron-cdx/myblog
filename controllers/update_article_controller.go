/**
 * @Author: caoduanxi
 * @Date: 2021/3/17 14:03
 * @Motto: Keep thinking, keep coding!
 */

package controllers

import "myblog/models"

type UpdateArticleController struct {
	BaseController
}

// 更新文章控制器,获取展示的页面
func (this *UpdateArticleController) Get() {
	id, _ := this.GetInt("id")
	article := models.QueryArticleWithId(id)
	this.Data["Title"] = article.Title
	this.Data["Tags"] = article.Tags
	this.Data["Short"] = article.Short
	this.Data["Content"] = article.Content
	this.Data["Id"] = article.Id
	this.TplName = "write_article.html"
}

// 文章修改，使用post实现
func (this *UpdateArticleController) Post() {
	id, _ := this.GetInt("id")
	// 获取表单中具体的数据
	title := this.GetString("title")
	tags := this.GetString("tags")
	short := this.GetString("short")
	content := this.GetString("content")
	// 获取具体的实体对象，实例化，如果是使用默认值可以通过 键值对的形式传递即可
	article := models.Article{Id: id, Title: title, Tags: tags, Short: short, Content: content}
	// 更新数据库
	_, err := models.UpdateArticle(article)
	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "更新成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "更新失败"}
	}
	this.ServeJSON()
}
