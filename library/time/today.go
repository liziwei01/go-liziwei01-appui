/*
 * @Author: liziwei01
 * @Date: 2022-02-26 16:49:10
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-02-26 16:51:16
 * @Description: file content
 */
package time

import (
	"context"
	"time"
)

func GetTodayTimeStamp(ctx context.Context) (int64, int64) {
	return GetTodayTime(ctx, time.Now().Unix())
}

/**
 * 获取当天00:00:00 和23:59:59对应的时间戳
 * @param {context.Context} ctx
 * @return {*}
 */
func GetTodayTime(ctx context.Context, timeInt int64) (int64, int64) {
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
