package controllers

import "myblog/models"

type HomeController struct {
	BaseController
}

/**
this.Data用于设置模板html中的数据
this.TplName用于跳转到具体的模板html中

因为涉及到tags文章标签跳转的问题，所以在这个地方需要对首页的数据进行改造
*/
func (this *HomeController) Get() {
	/*page, _ := this.GetInt("page")
	if page <= 0 {
		page = 1
	}
	var artList []models.Article
	artList, _ = models.FindArticleWithPage(page)
	fmt.Println("article list:", artList)
	this.Data["PageCode"] = models.ConfigHomeFooterPageCode(page)
	this.Data["HasFooter"] = true
	fmt.Println("IsLogin:", this.IsLogin, this.Loginuser)
	this.Data["Content"] = models.MakeHomeBlocks(artList, this.IsLogin)
	this.TplName = "home.html"*/

	tag := this.GetString("tag")
	page, _ := this.GetInt("page")
	// 获取具体的文章数
	var articleList []models.Article
	if len(tag) > 0 {
		// 按照标签搜索文章
		articleList, _ = models.QueryArticleWithTag(tag)
		this.Data["HasFooter"] = false
	} else {
		if page <= 0 {
			page = 1
		}
		// 否则的话查找当前页面的page
		articleList, _ = models.FindArticleWithPage(page)
		this.Data["PageCode"] = models.ConfigHomeFooterPageCode(page)
		this.Data["HasFooter"] = true
	}
	// 上面的行为是无视了tag和page的分别请求，只是在展示的时候作为人为适当的处理操作
	this.Data["Content"] = models.MakeHomeBlocks(articleList, this.IsLogin)
	this.TplName = "home.html"
}
