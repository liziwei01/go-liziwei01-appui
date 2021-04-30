/*
 * @Author: 		liziwei01
 * @Date: 			2021-04-19 15:00:00
 * @LastEditTime: 	2021-04-19 15:00:00
 * @LastEditors: 	liziwei01
 * @Description: 	搜索论文服务后台控制层：这一层负责与前端交互
 * @FilePath: 		/std/go-liziwei01-appui/modules/erg3020/controllers/paper/paper.go
 */
package paper

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gogf/gf/util/gconv"

	searchModel "go-liziwei01-appui/modules/erg3020/model/search"
	paperService "go-liziwei01-appui/modules/erg3020/services/paper"
)

var ctx = context.Background()

/**
 * @description: 搜索论文服务后台控制层处理逻辑
 * @param {http.ResponseWriter} response
 * @param {*http.Request} request
 * @return {*}
 */
func GetPaperList(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "text/json")
	response.Header().Set("Access-Control-Allow-Origin", "*")
	log.Printf("Request from %s\n", request.URL)
	log.Print(request)
	params, err := inputGetPaperList(ctx, request)
	if err != nil {
		response.WriteHeader(400)
		response.Write([]byte(fmt.Sprintf("controller.GetPaperList failed with err: %s", err.Error())))
		return
	}
	res, err := paperService.GetPaperList(ctx, params)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(fmt.Sprintf("controller.GetPaperList failed with err: %s", err.Error())))
		return
	}
	ret, err := json.Marshal(res)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(fmt.Sprintf("controller.GetPaperList json marshal failed with err: %s", err.Error())))
		return
	}
	io.WriteString(response, string(ret))
}

/**
 * @description: 提取传入的http请求内的参数
 * @param {*http.Request} request http请求
 * @return {searchModel.PaperSearchParams}
 */
func inputGetPaperList(ctx context.Context, request *http.Request) (searchModel.PaperSearchParams, error) {
	// 客户端接受的参数处理
	query := request.URL.Query()
	name, _ := ioutil.ReadAll(request.Body)
	fmt.Println(name)
	pageIndexStr := query.Get("pageIndex")   // 选择显示页，默认第1页
	pageLengthStr := query.Get("pageLength") // 每页显示几条，默认10条
	title := query.Get("title")
	authors := query.Get("authors")
	publishStartTimeStr := query.Get("startTime") // 按发表时间筛选
	publishEndTimeStr := query.Get("endTime")     // 按发表时间筛选
	journal := query.Get("journal")

	pageIndex := gconv.Uint(pageIndexStr)
	pageLength := gconv.Uint(pageLengthStr)
	publishStartTime := gconv.Int64(publishStartTimeStr)
	publishEndTime := gconv.Int64(publishEndTimeStr)
	authorsStr := gconv.String(authors)

	if pageIndex == 0 {
		pageIndex = gconv.Uint("1")
	}
	if pageLength == 0 {
		pageLength = gconv.Uint("10")
	}
	// 默认今天
	if publishStartTime == 0 {
		publishStartTime, _ = getTodayTimeStamp(ctx, time.Now().Unix())
	} else {
		publishStartTime, _ = getTodayTimeStamp(ctx, publishStartTime)
	}
	if publishEndTime == 0 {
		_, publishEndTime = getTodayTimeStamp(ctx, time.Now().Unix())
	} else {
		_, publishEndTime = getTodayTimeStamp(ctx, publishEndTime)
	}

	params := searchModel.PaperSearchParams{
		PageIndex:  pageIndex,
		PageLength: pageLength,
		Title:      title,
		Authors:    authorsStr,
		StartTime:  publishStartTime,
		EndTime:    publishEndTime,
		Journal:    journal,
	}
	return params, nil
}

/**
 * 获取当天00:00:00 和23:59:59对应的时间戳
 * @param {context.Context} ctx
 * @return {*}
 */
func getTodayTimeStamp(ctx context.Context, timeInt int64) (int64, int64) {
	timeStr := time.Unix(timeInt, 0).Format("2006-01-02")
	//获取当前时区
	loc, _ := time.LoadLocation("Local")
	//日期当天0点时间戳(拼接字符串)
	startDate := timeStr + "_00:00:00"
	startTime, _ := time.ParseInLocation("2006-01-02_15:04:05", startDate, loc)
	//日期当天23时59分时间戳
	endDate := timeStr + "_23:59:59"
	end, _ := time.ParseInLocation("2006-01-02_15:04:05", endDate, loc)
	//返回当天0点和23点59分的时间戳
	return startTime.Unix(), end.Unix()
}
