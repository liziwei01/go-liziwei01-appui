/*
 * @Author: liziwei01
 * @Date: 2021-04-19 15:00:00
 * @LastEditTime: 2021-05-30 02:25:46
 * @LastEditors: liziwei01
 * @Description: 分发csc3170路由
 * @FilePath: /github.com/liziwei01/go-liziwei01-appui/modules/csc3170/routers/router.go
 */
package routers

import (
	"net/http"

	"github.com/liziwei01/go-liziwei01-appui/modules/csc3170/controllers/star"
)

/**
 * @description: 后台csc3170路由分发
 * @param {*}
 * @return {*}
 */
func Init() {
	http.HandleFunc("/userList", star.GetUserList)
	http.HandleFunc("/insertUser", star.InsertUser)
}
