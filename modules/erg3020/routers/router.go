/*
 * @Author: liziwei01
 * @Date: 2021-04-19 15:00:00
 * @LastEditTime: 2022-02-26 19:57:37
 * @LastEditors: liziwei01
 * @Description: 分发erg3020路由
 * @FilePath: /github.com/liziwei01/go-liziwei01-appui/modules/erg3020/routers/router.go
 */
package routers

import (
	"net/http"

	"github.com/liziwei01/go-liziwei01-appui/modules/erg3020/controllers/paper"
)

/**
 * @description: 搜索论文服务后台路由分发
 * @param {*}
 * @return {*}
 */
func Init() {
	http.HandleFunc("/paperList", paper.GetPaperList)
	http.HandleFunc("/paper", paper.GetPaper)
	http.HandleFunc("/addPaperList", paper.AddPaperList)
}
