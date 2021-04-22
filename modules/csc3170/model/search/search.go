/*
 * @Author: 		liziwei01
 * @Date: 			2021-04-19 15:00:00
 * @LastEditTime: 	2021-04-19 15:00:00
 * @LastEditors: 	liziwei01
 * @Description: 	搜索模型
 * @FilePath: 		go-liziwei01-appui/modules/csc3170/model/search/search.go
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
