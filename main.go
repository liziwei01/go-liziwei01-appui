/*
 * @Author: 		liziwei01
 * @Date: 			2021-04-19 15:00:00
 * @LastEditTime: 	2021-04-19 15:00:00
 * @LastEditors: 	liziwei01
 * @Description: 	main
 * @FilePath: 		go-liziwei01-appui/main.go
 */
package main

import (
	"go-liziwei01-appui/httpapi"

	"github.com/liziwei01/go-liziwei01-library/bootstrap"
)

/**
 * @description: main函数
 * @param {*}
 * @return {*}
 */
func main() {
	bootstrap.Init()
	httpapi.InitRouters()
}
