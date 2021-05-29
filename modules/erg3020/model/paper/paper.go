/*
 * @Author: liziwei01
 * @Date: 2021-04-19 15:00:00
 * @LastEditTime: 2021-05-30 02:26:57
 * @LastEditors: liziwei01
 * @Description: 搜索论文服务论文模型
 * @FilePath: /github.com/liziwei01/go-liziwei01-appui/modules/erg3020/model/paper/paper.go
 */
package paper

type PaperInfo struct {
	IndexNumber     int64  `db:"index_number" json:"index_number"`
	Title           string `db:"title" json:"title"`
	Authors         string `db:"author" json:"authors"`
	PublishTime     int64  `db:"publish_time" json:"publish_time"`
	Journal         string `db:"journal" json:"journal"`
	References      string `db:"ref" json:"references"`
	TotalCites      int64  `db:"total_cites" json:"total_cites"`
	Score           int64  `db:"score" json:"score"`
	ScoreSimilarity int    `json:"score_similarity"`
}
