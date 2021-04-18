/*
 * @Author: 		liziwei01
 * @Date: 			2021-04-19 15:00:00
 * @LastEditTime: 	2021-04-19 15:00:00
 * @LastEditors: 	liziwei01
 * @Description: 	启动APP
 * @FilePath: 		go-liziwei01-appui/bootstrap/bootstrap.go
 */
package bootstrap

import (
	"fmt"
	"go-liziwei01-appui/httpapi"
)

/**
 * @description: 后台启动APP
 * @param {*}
 * @return {*}
 */
func Init() {
	fmt.Println("bootstrap")
	httpapi.InitRouters()
}