/**
 * @Author: caoduanxi
 * @Date: 2021/3/17 16:17
 * @Motto: Keep thinking, keep coding!
 */

package models

import "myblog/utils"

type Album struct {
	Id         int
	Filepath   string
	Filename   string
	Status     int
	CreateTime int64
}

// 插入具体的数据
func InsertAlbum(album Album) (int64, error) {
	return utils.ModifyDB("insert into album(filepath,filename,status,create_time) values (?,?,?,?)",
		album.Filepath, album.Filename, album.Status, album.CreateTime)
}

// 查找所有的图片
func FindAllAlbums() ([]Album, error) {
	rows, err := utils.QueryDB("select id,filepath,filename,status,create_time from album")
	if err != nil {
		return nil, err
	}
	var albums []Album
	for rows.Next() {
		id := 0
		filepath := ""
		filename := ""
		status := 0
		var createTime int64
		createTime = 0
		rows.Scan(&id, &filepath, &filename, &status, &createTime)
		album := Album{id, filepath, filename, status, createTime}
		albums = append(albums, album)
	}
	return albums, nil
}
