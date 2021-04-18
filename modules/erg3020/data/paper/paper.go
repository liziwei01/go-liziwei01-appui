/*
 * @Author: 		liziwei01
 * @Date: 			2021-04-19 15:00:00
 * @LastEditTime: 	2021-04-19 15:00:00
 * @LastEditors: 	liziwei01
 * @Description: 	搜索论文服务数据层
 * @FilePath: 		go-liziwei01-appui/modules/erg3020/data/paper/paper.go
 */
package paper

import (
	"context"
	"fmt"

	paperDao "go-liziwei01-appui/modules/erg3020/dao/paper"
	paperModel "go-liziwei01-appui/modules/erg3020/model/paper"
	searchModel "go-liziwei01-appui/modules/erg3020/model/search"

	"github.com/gogf/gf/util/gconv"
)

/**
 * @description: 搜索论文服务后台数据层处理逻辑
 * @param {searchModel.PaperSearchParams} params
 * @return {[]paperModel.PaperInfo}
 */
func GetPaperList(ctx context.Context, params searchModel.PaperSearchParams) ([]paperModel.PaperInfo, error) {
	fmt.Println("service->GetPaperList")
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
	// to do

	return 0, nil
}

/**
 * @description: 处理数据库获取的数据
 * @param {searchModel.PaperSearchParams} params
 * @param {[]paperModel.PaperInfo} papersInfo
 * @param {int64} count
 * @return {map[string]interface{}}
 */
func FormatPaperInfo(ctx context.Context, params searchModel.PaperSearchParams, papersInfo []paperModel.PaperInfo, count int64) (map[string]interface{}, error) {
	// to do

	return map[string]interface{}{
		"list":  papersInfo,
		"count": gconv.String(count),
	}, nil
}
