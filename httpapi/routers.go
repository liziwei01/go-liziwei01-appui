/*
 * @Author: 		liziwei01
 * @Date: 			2021-04-19 15:00:00
 * @LastEditTime: 2021-05-30 01:40:03
 * @LastEditors: liziwei01
 * @Description: 	启动http服务器并开始监听8086端口
 * @FilePath: 		github.com/liziwei01/go-liziwei01-appui/httpapi/httpapi.go
 */
package httpapi

import (
	"io"
	"net/http"

	csc3170Routers "github.com/liziwei01/go-liziwei01-appui/modules/csc3170/routers"
	erg3020Routers "github.com/liziwei01/go-liziwei01-appui/modules/erg3020/routers"
)

/**
 * @description: 后台启动路由分发
 * @param {*}
 * @return {*}
 */
func InitRouters() {
	// init routers
	erg3020Routers.Init()
	csc3170Routers.Init()

	// 兜底路由
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		io.WriteString(rw, "Hello! THis is Ziwei. Use get method and routers to explore.")
	})
}
