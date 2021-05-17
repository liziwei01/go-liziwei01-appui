/*
 * @Author: 		liziwei01
 * @Date: 			2021-04-19 15:00:00
 * @LastEditTime: 	2021-04-30 20:00:00
 * @LastEditors: 	liziwei01
 * @Description: 	搜索论文服务后台控制层：这一层负责与前端交互
 * @FilePath: 		/std/go-liziwei01-appui/modules/erg3020/controllers/paper/paper.go
 */
package paper

import (
	"context"
	"net/http"
	"time"

	"github.com/gogf/gf/util/gconv"
	errBase "github.com/liziwei01/go-liziwei01-library/model/error"
	"github.com/liziwei01/go-liziwei01-library/model/ghttp"
	"github.com/liziwei01/go-liziwei01-library/model/logit"

	searchModel "go-liziwei01-appui/modules/erg3020/model/search"
	paperService "go-liziwei01-appui/modules/erg3020/services/paper"
	paperScript "go-liziwei01-appui/script/erg3020/readcsv"
)

var ctx = context.Background()

/**
 * @description: 搜索论文服务后台控制层处理逻辑
 * @param {http.ResponseWriter} response
 * @param {*http.Request} request
 * @return {*}
 */
func GetPaperList(response http.ResponseWriter, request *http.Request) {
	logit.Logger.Info("/paperList")
	g := ghttp.Default(&request, &response)
	// 获取前端传入的参数
	params, err := inputGetPaperList(ctx, g)
	if err != nil {
		g.Write(params, errBase.ErrorNoClient, err)
		logit.Logger.Error(err)
	}
	// 获取根据评分和相似度排序的论文列表
	res, err := paperService.GetPaperList(ctx, params)
	if err != nil {
		g.Write(res, errBase.ErrorNoServer, err)
		logit.Logger.Error(err)
	}
	// 返回论文列表给前端
	g.Write(res, errBase.ErrorNoSuccess, err)
}

/**
 * @description: 提取传入的http请求内的参数
 * @param {*http.Request} request http请求
 * @return {searchModel.PaperSearchParams}
 */
func inputGetPaperList(ctx context.Context, g ghttp.Ghttp) (searchModel.PaperSearchParams, error) {
	// 客户端接受的参数处理
	pageIndexStr := g.Get("pageIndex")   // 选择显示页，默认第1页
	pageLengthStr := g.Get("pageLength") // 每页显示几条，默认10条
	title := g.Get("title")
	authors := g.Get("authors")
	publishStartTimeStr := g.Get("startTime") // 按发表时间筛选
	publishEndTimeStr := g.Get("endTime")     // 按发表时间筛选
	journal := g.Get("journal")

	types := g.Get("type")
	key := g.Get("key")
	switch types {
	case "title":
		title = key
	case "author":
		authors = key
	case "journal":
		journal = key
	default:
		// do nothing
	}

	pageIndex := gconv.Uint(pageIndexStr)
	pageLength := gconv.Uint(pageLengthStr)
	publishStartTime := gconv.Int64(publishStartTimeStr)
	publishEndTime := gconv.Int64(publishEndTimeStr)
	authorsStr := gconv.String(authors)

	if pageIndex == 0 {
		pageIndex = gconv.Uint("1")
	}
	if pageLength == 0 {
		pageLength = gconv.Uint("20")
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
 * @description: 搜索论文服务后台控制层处理逻辑
 * @param {http.ResponseWriter} response
 * @param {*http.Request} request
 * @return {*}
 */
func AddPaperList(response http.ResponseWriter, request *http.Request) {
	logit.Logger.Info("/addPaperList")
	g := ghttp.Default(&request, &response)
	params, err := inputAddPaperList(ctx, request)
	if err != nil {
		g.Write(params, errBase.ErrorNoClient, err)
		logit.Logger.Error(err)
	}
	res, err := paperScript.ParseBatchCsv(ctx, params)
	if err != nil {
		g.Write(res, errBase.ErrorNoServer, err)
		logit.Logger.Error(err)
	}
	err = paperScript.AddBatchAsync(ctx, res)
	g.Write(res, errBase.ErrorNoSuccess, err)
}

func inputAddPaperList(ctx context.Context, request *http.Request) (string, error) {
	query := request.URL.Query()
	fileName := query.Get("file_name")
	return fileName, nil
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
