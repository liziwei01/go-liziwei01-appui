/*
 * @Author: liziwei01
 * @Date: 2021-05-29 15:14:24
 * @LastEditors: liziwei01
 * @LastEditTime: 2021-05-30 02:21:35
 * @Description: file content
 * @FilePath: /github.com/liziwei01/go-liziwei01-appui/bootstrap/bootstrap.go
 */

package bootstrap

import (
	"context"
	"log"
)

/**
 * @description: start APP
 * @param {*}
 * @return {*}
 */
func Init() {
	// 解析应用配置
	config, err := ParserAppConfig(appConfPath)
	if err != nil {
		log.Fatal("ParserAppConfig failed")
	}
	app := NewApp(context.Background(), config)

	// 启动APP
	app.Start()
}
