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
	Order		int64	`ddb:"order" 		json:"order" 		c:"-"`
	Title		string	`ddb:"title" 		json:"title" 		c:"-"`
	Authors		string	`ddb:"author" 		json:"authors" 		c:"-"`
	PublishTime	int64	`ddb:"publish_time" json:"publish_time" c:"-"`
	Journal		string	`ddb:"journal" 		json:"journal" 		c:"-"`
	References	string	`ddb:"references" 	json:"references" 	c:"-"`
	Point		int64	`ddb:"point" 		json:"point"		c:"-"`
}