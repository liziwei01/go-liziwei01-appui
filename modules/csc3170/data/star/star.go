/*
 * @Author: 		liziwei01
 * @Date: 			2021-04-19 15:00:00
 * @LastEditTime: 	2021-04-19 15:00:00
 * @LastEditors: 	liziwei01
 * @Description: 	搜索明星服务数据层：要从数据库获取或者要写入的数据在这里处理
 * @FilePath: 		go-liziwei01-appui/modules/csc3170/data/star/star.go
 */
package star

import (
	"context"

	starDao "go-liziwei01-appui/modules/csc3170/dao/star"
	starModel "go-liziwei01-appui/modules/csc3170/model/star"
	searchModel "go-liziwei01-appui/modules/csc3170/model/search"

	"github.com/gogf/gf/util/gconv"
)

/**
 * @description: 搜索明星服务后台数据层处理逻辑
 * @param {searchModel.StarSearchParams} params
 * @return {[]starModel.StarInfo}
 */
func GetStarList(ctx context.Context, params searchModel.StarSearchParams) ([]starModel.StarInfo, error) {
	res, err := starDao.GetStarList(ctx, params)
	if err != nil {
		return make([]starModel.StarInfo, 0), err
	}
	return res, nil
}

/**
 * @description: 计数处理
 * @param {searchModel.StarSearchParams} params
 * @return {map[string]interface{}}
 */
func GetStarPagesCount(ctx context.Context, params searchModel.StarSearchParams) (int64, error) {
	count, err := starDao.GetStarPagesCount(ctx, params)
	if err != nil {
		return 0, err
	}
	return count, nil
}

/**
 * @description: 处理数据库获取的数据
 * @param {searchModel.StarSearchParams} params
 * @param {[]starModel.StarInfo} starsInfo
 * @param {int64} count
 * @return {map[string]interface{}}
 */
func FormatStarInfo(ctx context.Context, params searchModel.StarSearchParams, starsInfo []starModel.StarInfo, count int64) (map[string]interface{}, error) {
	// to do

	return map[string]interface{}{
		"params": params,
		"list":   starsInfo,
		"count":  gconv.String(count),
	}, nil
}
