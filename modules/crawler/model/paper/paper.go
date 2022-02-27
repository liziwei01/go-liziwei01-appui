/*
 * @Author: liziwei01
 * @Date: 2021-04-19 15:00:00
 * @LastEditTime: 2022-02-27 09:37:12
 * @LastEditors: liziwei01
 * @Description: 搜索论文服务论文模型
 * @FilePath: /github.com/liziwei01/go-liziwei01-appui/modules/erg3020/model/paper/paper.go
 */
package paper

type PaperInfo struct {
	Title   string `db:"title" json:"title"`
	Ref     string `db:"ref" json:"ref"`
	Content string `db:"content" json:"content"`
}
