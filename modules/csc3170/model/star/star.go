/*
 * @Author: liziwei01
 * @Date: 2021-04-19 15:00:00
 * @LastEditTime: 2021-05-30 02:25:18
 * @LastEditors: liziwei01
 * @Description: 搜索明星服务明星模型
 * @FilePath: /github.com/liziwei01/go-liziwei01-appui/modules/csc3170/model/star/star.go
 */
package star

type StarInfo struct {
	IndexNumber int64  `db:"index_number" json:"index_number" c:"-"`
	Title       string `db:"title" json:"title" c:"-"`
	Authors     string `db:"author" json:"authors" c:"-"`
	PublishTime int64  `db:"publish_time" json:"publish_time" c:"-"`
	Journal     string `db:"journal" json:"journal" c:"-"`
	References  string `db:"references" json:"references" c:"-"`
	Point       int64  `db:"point" json:"point" c:"-"`
}

type UserInfo struct {
	UserId   int64  `db:"user_id" json:"user_id" c:"-"`
	UserName string `db:"name" json:"user_name" c:"-"`
	Password string `db:"password" json:"password" c:"-"`
}
