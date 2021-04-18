/*
 * @Author: 		liziwei01
 * @Date: 			2021-04-19 15:00:00
 * @LastEditTime: 	2021-04-19 15:00:00
 * @LastEditors: 	liziwei01
 * @Description: 	搜索论文服务后台控制层
 * @FilePath: 		/std/go-liziwei01-appui/modules/erg3020/controllers/paper/paper.go
 */
package paper

import (
	"context"
	"fmt"
	"io"
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
	fmt.Println("controller->GetPaperList")

	params, valid := inputGetPaperList(ctx, request)
	if valid == false {
		io.WriteString(response, "hello, world! error\n")
		return
	}
	res, err := paperService.GetPaperList(ctx, params)
	if err != nil {
		io.WriteString(response, "hello, world! error\n")
		return
	}

	io.WriteString(response, res["count"].(string))
}

/**
 * @description: 提取传入的http请求内的参数
 * @param {*http.Request} request http请求
 * @return {searchModel.PaperSearchParams}
 */
func inputGetPaperList(ctx context.Context, request *http.Request) (searchModel.PaperSearchParams, bool) {
	// 客户端接受的参数处理
	query 				:= request.URL.Query()
	pageIndexStr 		:= query.Get("pageIndex")   // 选择显示页，默认第1页
	pageLengthStr 		:= query.Get("pageLength") // 每页显示几条，默认10条
	authors 			:= query.Get("authors")
	publishStartTimeStr := query.Get("startTime") // 按发表时间筛选
	publishEndTimeStr 	:= query.Get("endTime")     // 按发表时间筛选

	pageIndex 			:= gconv.Uint(pageIndexStr)
	pageLength 			:= gconv.Uint(pageLengthStr)
	publishStartTime 	:= gconv.Int64(publishStartTimeStr)
	publishEndTime 		:= gconv.Int64(publishEndTimeStr)
	authorsStr 			:= gconv.String(authors)

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
		PageIndex:		pageIndex,
		PageLength:		pageLength,
		Authors: 		authorsStr,
		StartTime:		publishStartTime,
		EndTime:		publishEndTime,
	}
	return params, true
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
