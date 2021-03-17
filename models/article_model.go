package models

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"log"
	"myblog/utils"
	"strconv"
)

type Article struct {
	Id         int
	Title      string
	Tags       string
	Short      string
	Content    string
	Author     string
	CreateTime int64
}

func AddArticle(article Article) (int64, error) {
	i, err := insertArticle(article)
	// 这里需要设置具体的总页数
	SetArticleRowsNum()
	return i, err
}

func insertArticle(article Article) (int64, error) {
	return utils.ModifyDB("insert into article (title,tags,short,content,author,create_time) "+
		"values(?,?,?,?,?,?)", article.Title, article.Tags, article.Short, article.Content, article.Author, article.CreateTime)
}

// 查询文章
func FindArticleWithPage(page int) ([]Article, error) {
	// 获取文件配置中每一页的文章数量
	num, _ := beego.AppConfig.Int("articleListPageNum")
	page -= 1 // 这里是因为页数是从1开始，但是mysql是从0开始所以减一
	fmt.Println("page ===>", page)
	return QueryArticleWithPage(page, num)
}

// 根据文章id获取文章
func QueryArticleWithId(id int) Article {
	row := utils.QueryRowDB("select id,title,tags,short,content,author,create_time from article where id = " + strconv.Itoa(id))
	fmt.Println(row)
	title := ""
	tags := ""
	short := ""
	content := ""
	author := ""
	var createTime int64
	createTime = 0
	// row.Scan查询数据的时候，一定要确保所有的字段一一对应才会进行赋值，否则值为空
	row.Scan(&id, &title, &tags, &short, &content, &author, &createTime)
	article := Article{id, title, tags, short, content, author, createTime}
	return article
}

func QueryArticleWithPage(page, num int) ([]Article, error) {
	sql := fmt.Sprintf("limit %d,%d", page*num, num)
	return QueryArticleWithOptions(sql)
}

func QueryArticleWithOptions(sql string) ([]Article, error) {
	sql = "select id,title,tags,short,content,author,create_time from article " + sql
	//sql = "select * from article " + sql
	rows, err := utils.QueryDB(sql)
	if err != nil {
		return nil, err
	}
	var artList []Article
	// 从查出来的多行数据中获取数值
	for rows.Next() {
		id := 0
		title := ""
		tags := ""
		short := ""
		content := ""
		author := ""
		var createTime int64
		createTime = 0
		rows.Scan(&id, &title, &tags, &short, &content, &author, &createTime)
		article := Article{id, title, tags, short, content, author, createTime}
		artList = append(artList, article)
	}
	fmt.Println("Article List:", artList)
	return artList, nil
}

// 更新数据库
func UpdateArticle(article Article) (int64, error) {
	return utils.ModifyDB("update article set title = ?, tags = ?,short=?,content=? where id = ?",
		article.Title, article.Tags, article.Short, article.Content, article.Id)
}

// 删除文章操作
// 注意删除文章之后需要更新文章总数目
func DeleteArticle(articleId int) (int64, error) {
	row, err := utils.ModifyDB("delete from article where id = ?", articleId)
	if err != nil {
		log.Printf("删除文章编号为%d的文章错误,error:%s", articleId, err.Error())
	}
	SetArticleRowsNum()
	return row, err
}

// 根据参数查询文章,这里是个根据标签查询
func QueryArticleWithParam(param string) []string {
	rows, err := utils.QueryDB(fmt.Sprintf("select %s from article", param))
	if err != nil {
		log.Println("query article with param meet error...", err)
	}
	var paramList []string
	for rows.Next() {
		args := ""
		rows.Scan(&args)
		paramList = append(paramList, args)
	}
	// 将标签转换为list个数输出
	return paramList
}

// 根据标签查询具体的文章
func QueryArticleWithTag(tag string) ([]Article, error) {
	sql := "where tags like '%&" + tag + "&%'"
	sql += "or tags like '%&" + tag + "'"
	sql += "or tags like '" + tag + "&%'"
	sql += "or tags like '" + tag + "'"
	return QueryArticleWithOptions(sql)
}
