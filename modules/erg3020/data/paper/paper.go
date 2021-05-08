/*
 * @Author: 		liziwei01
 * @Date: 			2021-04-19 15:00:00
 * @LastEditTime: 	2021-04-19 15:00:00
 * @LastEditors: 	liziwei01
 * @Description: 	搜索论文服务数据层：要从数据库获取或者要写入的数据在这里处理
 * @FilePath: 		go-liziwei01-appui/modules/erg3020/data/paper/paper.go
 */
package paper

import (
	"context"
	"time"

	paperDao "go-liziwei01-appui/modules/erg3020/dao/paper"
	paperModel "go-liziwei01-appui/modules/erg3020/model/paper"
	searchModel "go-liziwei01-appui/modules/erg3020/model/search"
)

/**
 * @description: 搜索论文服务后台数据层处理逻辑
 * @param {searchModel.PaperSearchParams} params
 * @return {[]paperModel.PaperInfo}
 */
func GetPaperList(ctx context.Context, params searchModel.PaperSearchParams) ([]paperModel.PaperInfo, error) {
	res, err := paperDao.GetPaperList(ctx, params)
	if err != nil {
		return make([]paperModel.PaperInfo, 0), err
	}
	return res, nil
}

/**
 * @description: 计数处理
 * @param {searchModel.PaperSearchParams} params
 * @return {map[string]interface{}}
 */
func GetPaperPagesCount(ctx context.Context, params searchModel.PaperSearchParams) (int64, error) {
	count, err := paperDao.GetPaperPagesCount(ctx, params)
	if err != nil {
		return 0, err
	}
	return count, nil
}

/**
 * @description: 处理数据库获取的数据
 * @param {searchModel.PaperSearchParams} params
 * @param {[]paperModel.PaperInfo} papersInfo
 * @param {int64} count
 * @return {map[string]interface{}}
 */
func FormatPaperInfo(ctx context.Context, params searchModel.PaperSearchParams, papersInfo []paperModel.PaperInfo, count int64) (map[string]interface{}, error) {
	var (
		res []map[string]interface{}
	)
	for _, v := range papersInfo {
		timeStr := time.Unix(v.PublishTime, 0).Format("2006-01-02_15:04:05")
		res = append(res, map[string]interface{}{
			"index_number": v.IndexNumber,
			"title":        v.Title,
			"authors":      v.Authors,
			"journal":      v.Journal,
			"publish_time": timeStr,
			"references":   v.References,
			"total_cites":  v.TotalCites,
			"score":        v.Score,
		})
	}
	return map[string]interface{}{
		"list":  res,
		"count": count,
	}, nil
}
