/*
 * @Author: 		liziwei01
 * @Date: 			2021-04-19 15:00:00
 * @LastEditTime: 	2021-04-19 15:00:00
 * @LastEditors: 	liziwei01
 * @Description: 	搜索明星服务后台控制层：这一层负责与前端交互
 * @FilePath: 		/std/github.com/liziwei01/go-liziwei01-appui/modules/csc3170/controllers/paper/star.go
 */
package star

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gogf/gf/util/gconv"

	"github.com/liziwei01/go-liziwei01-appui/modules/csc3170/model/search"
	searchModel "github.com/liziwei01/go-liziwei01-appui/modules/csc3170/model/search"
	starModel "github.com/liziwei01/go-liziwei01-appui/modules/csc3170/model/star"
	starService "github.com/liziwei01/go-liziwei01-appui/modules/csc3170/services/star"
)

var ctx = context.Background()

/**
 * @description: 插入用户数据后台控制层处理逻辑
 * @param {http.ResponseWriter} response
 * @param {*http.Request} request
 * @return {*}
 */
func InsertUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "text/json")
	response.Header().Set("Access-Control-Allow-Origin", "*")
	params, err := inputInsertUser(request)
	if err != nil {
		response.WriteHeader(400)
		response.Write([]byte(fmt.Sprintf("controller.InsertUser failed with err: %s", err.Error())))
		return
	}
	err = starService.InsertUser(ctx, params)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(fmt.Sprintf("controller.InsertUser failed with err: %s", err.Error())))
		return
	}
	response.WriteHeader(200)
	response.Write([]byte("{\"errmsg\": \"success\"}"))
}

/**
 * @description: 插入用户数据：获取postform形式的参数
 * @param {*http.Request} request
 * @return {*}
 */
func inputInsertUser(request *http.Request) (starModel.UserInfo, error) {
	var (
		ret          starModel.UserInfo
		nameFlag     = false
		idFlag       = false
		passwordFlag = false
	)
	request.ParseForm()
	for k, v := range request.PostForm {
		if k == "name" {
			ret.UserName = gconv.String(v[0])
			nameFlag = true
		}
		if k == "ID" {
			ret.UserId = gconv.Int64(v[0])
			idFlag = true
		}
		if k == "password" {
			ret.Password = gconv.String(v[0])
			passwordFlag = true
		}
	}
	if nameFlag && idFlag && passwordFlag {
		return ret, nil
	}
	return starModel.UserInfo{}, fmt.Errorf("inputInsertUser->params not enough")
}

/**
 * @description: 获取用户数据后台控制层处理逻辑
 * @param {http.ResponseWriter} response
 * @param {*http.Request} request
 * @return {*}
 */
func GetUserList(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "text/json")
	response.Header().Set("Access-Control-Allow-Origin", "*")
	params, err := inputGetUserList(request)
	if err != nil {
		response.WriteHeader(400)
		response.Write([]byte(fmt.Sprintf("controller.GetUserList failed with err: %s", err.Error())))
		return
	}
	fmt.Println(request)
	res, err := starService.GetUserList(ctx, params)
	if err != nil {
		response.WriteHeader(404)
		response.Write([]byte(fmt.Sprintf("controller.GetUserList failed with err: %s", err.Error())))
		return
	}
	ret, err := json.Marshal(res)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(fmt.Sprintf("controller.GetUserList failed with err: %s", err.Error())))
		return
	}
	response.WriteHeader(200)
	response.Write(ret)
}

/**
 * @description: 获取用户数据：获取postform形式的参数
 * @param {*http.Request} request
 * @return {*}
 */
func inputGetUserList(request *http.Request) (searchModel.UserSearchParams, error) {
	var (
		ret            searchModel.UserSearchParams
		nameFlag       = false
		pageIndexFlag  = false
		pageLengthFlag = false
	)
	err := request.ParseForm()
	if err != nil {
		log.Printf("csc3170.controller.inputGetUserList request.ParseForm failed with err: %s\n", err.Error())
		return search.UserSearchParams{}, err
	}
	fmt.Println("request.PostForm: ")
	fmt.Println(request.PostForm)
	for k, v := range request.PostForm {
		if k == "name" {
			ret.UserName = gconv.String(v[0])
			nameFlag = true
		}
		if k == "pageIndex" {
			ret.PageIndex = gconv.Uint(v[0])
			pageIndexFlag = true
		}
		if k == "pageLength" {
			ret.PageLength = gconv.Uint(v[0])
			pageLengthFlag = true
		}
	}
	if !pageIndexFlag {
		ret.PageIndex = gconv.Uint(1)
	}
	if !pageLengthFlag {
		ret.PageLength = gconv.Uint(10)
	}
	if nameFlag {
		return ret, nil
	}
	return search.UserSearchParams{}, fmt.Errorf("inputGetUserList->params not enough")
}
