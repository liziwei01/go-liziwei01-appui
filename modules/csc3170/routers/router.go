/*
 * @Author: 		liziwei01
 * @Date: 			2021-04-19 15:00:00
 * @LastEditTime: 	2021-04-19 15:00:00
 * @LastEditors: 	liziwei01
 * @Description: 	分发csc3170路由
 * @FilePath: 		go-liziwei01-appui/modules/csc3170/routers/router.go
 */
package routers

import (
	"net/http"

	"go-liziwei01-appui/modules/csc3170/controllers/star"
)

/**
 * @description: 后台csc3170路由分发
 * @param {*}
 * @return {*}
 */
func Init() {
	http.HandleFunc("/insertStars", star.GetStarList)
	http.HandleFunc("/insertPrograms", star.GetStarList)
	http.HandleFunc("/insertCompany", star.GetStarList)
	http.HandleFunc("/insertPlatform", star.GetStarList)
	http.HandleFunc("/insertUsers", star.GetStarList)
}
