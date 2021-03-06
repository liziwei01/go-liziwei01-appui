/*
 * @Author: liziwei01
 * @Date: 2021-04-19 15:00:00
 * @LastEditTime: 2022-02-27 18:19:31
 * @LastEditors: liziwei01
 * @Description: 搜索模型
 * @FilePath: /github.com/liziwei01/go-liziwei01-appui/modules/erg3020/model/search/search.go
 */
package search

type PaperSearchParams struct {
	PageIndex  uint   `json:"page_index"`
	PageLength uint   `json:"page_length"`
	Title      string `json:"title"`
	Authors    string `json:"authors"`
	Journal    string `json:"journal"`
	StartTime  int64  `json:"start_time"`
	EndTime    int64  `json:"end_time"`
	Ref        string `json:"ref"`
}
