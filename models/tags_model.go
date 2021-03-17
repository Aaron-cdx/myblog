/**
 * @Author: caoduanxi
 * @Date: 2021/3/17 14:50
 * @Motto: Keep thinking, keep coding!
 */

package models

import "strings"

/**
将tags数组通过遍历查询，获取到其中每个标签对应的文章个数
*/
func HandleTagsListData(tags []string) map[string]int {
	tagsMap := make(map[string]int)
	// 注意标签存入的时候是通过&符号进行拼接的，例: a&b&c&d
	for _, tag := range tags {
		tagList := strings.Split(tag, "&")
		// 然后对标签进行计数操作
		for _, v := range tagList {
			tagsMap[v]++
		}
	}
	return tagsMap
}
