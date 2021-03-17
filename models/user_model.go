package models

import (
	"fmt"
	"myblog/utils"
)

type User struct {
	Id         int
	Username   string
	Password   string
	Status     int
	CreateTime int64
}

// 插入新用户
func InsertUser(user User) (int64, error) {
	fmt.Println(user)
	return utils.ModifyDB("insert into users(username,password,status,create_time) values (?,?,?,?)",
		user.Username, user.Password, user.Status, user.CreateTime)
}

// 根据用户名查询id
func QueryUserWithUsername(username string) int {
	// 格式化以个string字符串
	sql := fmt.Sprintf("where username = '%s'", username)
	return QueryUserWithOptions(sql)
}

// 按条件查询
func QueryUserWithOptions(options string) int {
	sql := fmt.Sprintf("select id from users %s", options)
	fmt.Println(sql)
	row := utils.QueryRowDB(sql)
	id := 0
	row.Scan(&id) // 传入地址的话在修改过程中id的值会发生变化
	//if err != nil {
	//	log.Println("query user with options error :", err)
	//}
	return id
}

// 根据用户名和密码查询用户id
func QueryUserWithParam(username, password string) int {
	sql := fmt.Sprintf("where username = '%s' and password = '%s'", username, password)
	return QueryUserWithOptions(sql)
}
