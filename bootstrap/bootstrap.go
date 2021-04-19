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
	"context"
	"fmt"
	"log"
)

/**
 * @description: 后台启动APP
 * @param {*}
 * @return {*}
 */
func Init() {
	fmt.Println("bootstrap")
	config, err := ParserAppConfig(appConfPath)
	if err != nil {
		log.Fatal("init app failed")
	}
	fmt.Printf("listen: %s", config.HTTPServer.Listen)
	app := NewApp(context.Background(), config)
	app.Start()
}
