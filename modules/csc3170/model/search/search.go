/*
 * @Author: liziwei01
 * @Date: 2021-04-19 15:00:00
 * @LastEditTime: 2021-05-30 02:23:24
 * @LastEditors: liziwei01
 * @Description: 搜索模型
 * @FilePath: /github.com/liziwei01/go-liziwei01-appui/modules/csc3170/model/search/search.go
 */
package search

type StarSearchParams struct {
	PageIndex  uint   `json:"page_index"`
	PageLength uint   `json:"page_length"`
	Title      string `json:"title"`
	Authors    string `json:"authors"`
	StartTime  int64  `json:"start_time"`
	EndTime    int64  `json:"end_time"`
	Journal    string `json:"journal"`
}

type UserSearchParams struct {
	PageIndex    uint   `json:"page_index"`
	PageLength   uint   `json:"page_length"`
	UserName     string `json:"user_name"`
	UserId       int64  `json:"user_id"`
	UserPassword string `json:"user_password"`
}
