package utils

import (
	"bytes"
	"crypto/md5"
	"database/sql"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	"github.com/russross/blackfriday"
	"github.com/sourcegraph/syntaxhighlight"
	"html/template"
	"log"
	"time"
)

// 定义数据库操作对象
var db *sql.DB

//func init() {
//	InitMysql()
//}

/**
InitMysql 实现具体的数据库的连接操作
*/
func InitMysql() {
	fmt.Println("Init MySQL...")
	driverName, err := beego.AppConfig.String("drivername")
	if err != nil {
		fmt.Println("get DriverName meet error....")
		return
	}
	//orm.RegisterModel()
	// 注册数据库驱动
	_ = orm.RegisterDriver("mysql", orm.DRMySQL)

	// 数据库连接
	user, _ := beego.AppConfig.String("mysqluser")
	pwd, _ := beego.AppConfig.String("mysqlpwd")
	host, _ := beego.AppConfig.String("host")
	port, _ := beego.AppConfig.String("port")
	dbname, _ := beego.AppConfig.String("dbname")

	dbConn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"
	//fmt.Println(dbConn)

	// 这里表示已经连接上了数据库
	//err = orm.RegisterDataBase("default", driverName, dbConn)
	db1, err := sql.Open(driverName, dbConn)
	if err != nil {
		log.Println("connected database error...")
		fmt.Println(err.Error())
		return
	}
	db = db1
	log.Println("connected database success...")
	CreateTableWithUser()
	CreateTableWithArticle()
	CreateTableWithAlbum()
}

// 创建图片表格
func CreateTableWithAlbum() {
	sql := `create table if not exists album(
		id int(4) primary key not null auto_increment,
		filepath varchar(255),
		filename varchar(64),
		status int(4),
		create_time int(10)
	)`
	ModifyDB(sql)
}

// 创建用户表
func CreateTableWithUser() {
	sql := `create table if not exists users(
		id int(4) primary key auto_increment not null,
		username varchar(50),
		password varchar(50),
		status int(4),
		create_time int(10)
	);`
	ModifyDB(sql)
}

// 操作db
func ModifyDB(sql string, args ...interface{}) (int64, error) {
	exec, err := db.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	// 获取影响的行数
	count, err := exec.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}

// 查询db，获取单条记录
func QueryRowDB(sql string) *sql.Row {
	return db.QueryRow(sql)
}

// 查询db，或获取多条记录
func QueryDB(sql string) (*sql.Rows, error) {
	return db.Query(sql)
}

// 获取MD5值
func MD5(str string) string {
	// 输出的是16进制
	md5str := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return md5str
}

// 创建文章表
func CreateTableWithArticle() {
	sql := `create table if not exists article(
	id int(4) primary key auto_increment not null,
	title varchar(30),
	author varchar(30),
	tags varchar(30),
	short varchar(30),
	content longtext,
	create_time int(10)
	);`
	ModifyDB(sql)
}

// 转化为时间戳
func SwitchTimeStampToData(createTime int64) string {
	timeString := time.Unix(createTime, 0).Format("2006-01-02 15:04:05")
	return timeString
}

// 将Markdown转化为HTML
func SwitchMarkdownToHtml(content string) template.HTML {
	// 对一般的Markdown进行处理
	markdown := blackfriday.MarkdownCommon([]byte(content))
	// 获取其中的html文档
	doc, _ := goquery.NewDocumentFromReader(bytes.NewReader(markdown))
	// 这里是对代码进行高亮处理操作
	doc.Find("code").Each(func(i int, selection *goquery.Selection) {
		light, _ := syntaxhighlight.AsHTML([]byte(selection.Text()))
		selection.SetHtml(string(light))
	})
	htmlString, _ := doc.Html()
	return template.HTML(htmlString)
}
