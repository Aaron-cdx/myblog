package models

import (
	"bytes"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"html/template"
	"myblog/utils"
	"strconv"
	"strings"
)

/**
控制首页展示
*/

type HomeBlockParam struct {
	Id         int
	Title      string
	Tags       []TagLink
	Short      string
	Content    string
	Author     string
	CreateTime string
	// 查看文章地址
	Link string

	// 修改文章地址
	UpdateLink string
	DeleteLink string

	// 记录是否登录
	IsLogin bool
}

type TagLink struct {
	TagName string
	TagUrl  string
}

// 添加分页结构体
type HomeFooterPageCode struct {
	HasPre   bool
	HasNext  bool
	ShowPage string
	PreLink  string
	NextLink string
}

// 首页显示内容,f
func MakeHomeBlocks(articles []Article, isLogin bool) template.HTML {
	htmlHome := ""
	// for index, value := range objects{} 实现遍历
	for _, art := range articles {
		// 转换为模板所需要的数据
		homePageParam := HomeBlockParam{}
		homePageParam.Id = art.Id
		homePageParam.Title = art.Title
		homePageParam.Tags = createTagsLinks(art.Tags)
		homePageParam.Short = art.Short
		homePageParam.Content = art.Content
		homePageParam.Author = art.Author
		homePageParam.CreateTime = utils.SwitchTimeStampToData(art.CreateTime)
		homePageParam.Link = "/article/" + strconv.Itoa(art.Id)
		homePageParam.UpdateLink = "/article/update?id=" + strconv.Itoa(art.Id)
		homePageParam.DeleteLink = "/article/delete?id=" + strconv.Itoa(art.Id)
		homePageParam.IsLogin = isLogin

		// 处理变量，利用ParseFile解析该文件，用于插入变量
		t, _ := template.ParseFiles("views/block/home_block.html")
		buffer := bytes.Buffer{}
		t.Execute(&buffer, homePageParam)
		htmlHome += buffer.String()
	}
	fmt.Println("htmlHome ===>", htmlHome)
	return template.HTML(htmlHome)
}

// tag字符串转化为首页需要的数据结构,因为可能存在多个tag标签，使用&标签进行分隔
func createTagsLinks(tag string) []TagLink {
	var tagLink []TagLink
	tagsParam := strings.Split(tag, "&")
	for _, tag := range tagsParam {
		tagLink = append(tagLink, TagLink{tag, "/?tag=" + tag})
	}
	return tagLink
}

// 实现翻页功能
func ConfigHomeFooterPageCode(page int) HomeFooterPageCode {
	pageCode := HomeFooterPageCode{}
	// 总条数
	num := GetArticleRowsNum()
	// 获取每一页的条数
	pageRow, _ := beego.AppConfig.Int("articleListPageNum")
	// 总页数
	allPageNum := (num-1)/pageRow + 1
	pageCode.ShowPage = fmt.Sprintf("%d/%d", page, allPageNum)

	if page <= 1 {
		// 前一页需要关闭
		pageCode.HasPre = false
	} else {
		pageCode.HasPre = true
	}

	if page >= allPageNum {
		pageCode.HasNext = false
	} else {
		pageCode.HasNext = true
	}
	// 这里需要设置访问的前一页和后一页的链接
	pageCode.PreLink = "/?page=" + strconv.Itoa(page-1)
	pageCode.NextLink = "/?page=" + strconv.Itoa(page+1)
	return pageCode
}

var articleRowsNum = 0

// 获取表的行数
func GetArticleRowsNum() int {
	if articleRowsNum == 0 {
		articleRowsNum = QueryArticleRowNum()
	}
	return articleRowsNum
}

func QueryArticleRowNum() int {
	row := utils.QueryRowDB("select count(id) from article")
	num := 0
	row.Scan(&num)
	return num
}

// 新增文章或者删除文章的时候会发生具体的页数变化
func SetArticleRowsNum(){
	articleRowsNum = QueryArticleRowNum()
}
