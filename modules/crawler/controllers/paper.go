/*
 * @Author: liziwei01
 * @Date: 2021-04-19 15:00:00
 * @LastEditTime: 2022-02-26 17:02:18
 * @LastEditors: liziwei01
 * @Description: 搜索论文服务后台控制层：这一层负责与前端交互
 * @FilePath: /github.com/liziwei01/go-liziwei01-appui/modules/crawler/controllers/paper/paper.go
 */
package paper

import (
	"context"
	"net/http"

	"github.com/gogf/gf/util/gconv"
	"github.com/liziwei01/go-liziwei01-appui/library/time"
	crawlServices "github.com/liziwei01/go-liziwei01-appui/modules/crawler/services"
	errBase "github.com/liziwei01/go-liziwei01-library/model/error"
	"github.com/liziwei01/go-liziwei01-library/model/ghttp"
	"github.com/liziwei01/go-liziwei01-library/model/logit"

	searchModel "github.com/liziwei01/go-liziwei01-appui/modules/crawler/model/search"
)

var ctx = context.Background()

/**
 * @description: 搜索论文服务后台控制层处理逻辑
 * @param {http.ResponseWriter} response
 * @param {*http.Request} request
 * @return {*}
 */
func CrawlPaper(response http.ResponseWriter, request *http.Request) {
	logit.Logger.Info("/crawlPaper")
	g := ghttp.Default(&request, &response)
	// 获取前端传入的参数
	params, err := inputCrawlPaper(ctx, g)
	if err != nil {
		logit.Logger.Error(err)
	}
	err = crawlServices.CrawlPaper(ctx, params)
	g.Write("", errBase.ErrorNoSuccess, err)
}

/**
 * @description: 提取传入的http请求内的参数
 * @param {*http.Request} request http请求
 * @return {searchModel.PaperSearchParams}
 */
func inputCrawlPaper(ctx context.Context, g ghttp.Ghttp) (searchModel.PaperSearchParams, error) {
	// 客户端接受的参数处理
	pageIndexStr := g.Get("pageIndex")   // 选择显示页，默认第1页
	pageLengthStr := g.Get("pageLength") // 每页显示几条，默认10条
	title := g.Get("title")
	authors := g.Get("authors")
	publishStartTimeStr := g.Get("startTime") // 按发表时间筛选
	publishEndTimeStr := g.Get("endTime")     // 按发表时间筛选
	journal := g.Get("journal")

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
		publishStartTime, _ = time.GetTodayTimeStamp(ctx)
	} else {
		publishStartTime, _ = time.GetTodayTime(ctx, publishStartTime)
	}
	if publishEndTime == 0 {
		_, publishEndTime = time.GetTodayTimeStamp(ctx)
	} else {
		_, publishEndTime = time.GetTodayTime(ctx, publishEndTime)
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
