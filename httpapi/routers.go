/*
 * @Author: 		liziwei01
 * @Date: 			2021-04-19 15:00:00
 * @LastEditTime: 	2021-04-19 15:00:00
 * @LastEditors: 	liziwei01
 * @Description: 	启动http服务器并开始监听8086端口
 * @FilePath: 		go-liziwei01-appui/httpapi/httpapi.go
 */
package httpapi

import (
	"fmt"
	"net/http"
	"log"

	csc3170Routers "go-liziwei01-appui/modules/csc3170/routers"
	erg3020Routers "go-liziwei01-appui/modules/erg3020/routers"
)

/**
 * @description: 后台启动路由分发
 * @param {*}
 * @return {*}
 */
func InitRouters() {
	fmt.Println("init routers")

	erg3020Routers.Init()
	csc3170Routers.Init()
	err := http.ListenAndServe("localhost:8086", nil)
	if err != nil {
		log.Panic("server 500")
	}
}
