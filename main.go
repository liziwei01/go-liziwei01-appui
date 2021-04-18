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
	"fmt"

	"go-liziwei01-appui/bootstrap"
)

/**
 * @description: main函数
 * @param {*}
 * @return {*}
 */
func main() {
	fmt.Println("main start")
	bootstrap.Init()
}