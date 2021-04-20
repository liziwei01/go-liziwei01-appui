/*
 * @Author: 		liziwei01
 * @Date: 			2021-04-19 15:00:00
 * @LastEditTime: 	2021-04-19 15:00:00
 * @LastEditors: 	liziwei01
 * @Description: 	搜索论文服务论文模型
 * @FilePath: 		go-liziwei01-appui/modules/erg3020/model/paper/paper.go
 */
package paper

type PaperInfo struct {
	IndexNumber int64  `db:"index_number"	json:"index_number" c:"-"`
	Title       string `db:"title" 			json:"title" 		c:"-"`
	Authors     string `db:"author" 		json:"authors" 		c:"-"`
	PublishTime int64  `db:"publish_time"	json:"publish_time" c:"-"`
	Journal     string `db:"journal" 		json:"journal" 		c:"-"`
	References  string `db:"references" 	json:"references" 	c:"-"`
	Point       int64  `db:"point" 			json:"point"		c:"-"`
}
