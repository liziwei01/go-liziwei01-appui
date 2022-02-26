/*
 * @Author: liziwei01
 * @Date: 2021-04-19 15:00:00
 * @LastEditTime: 2022-02-26 19:51:16
 * @LastEditors: liziwei01
 * @Description: 分发crawler路由
 * @FilePath: /github.com/liziwei01/go-liziwei01-appui/modules/crawler/routers/router.go
 */
package routers

import (
	"net/http"

	crawler "github.com/liziwei01/go-liziwei01-appui/modules/crawler/controllers"
)

/**
 * @description: 搜索论文服务后台路由分发
 * @param {*}
 * @return {*}
 */
func Init() {
	http.HandleFunc("/crawlPaper", crawler.CrawlPaper)
}
