/*
 * @Author: 		liziwei01
 * @Date: 			2021-04-19 15:00:00
 * @LastEditTime: 	2021-04-19 15:00:00
 * @LastEditors: 	liziwei01
 * @Description: 	搜索明星服务服务层：这一层提供完整的服务
 * @FilePath: 		go-liziwei01-appui/modules/csc3170/services/paper/star.go
 */
package star

import (
	"context"

	starData "go-liziwei01-appui/modules/csc3170/data/star"
	searchModel "go-liziwei01-appui/modules/csc3170/model/search"
)

/**
 * @description: 搜索明星服务后台服务层处理逻辑
 * @param {searchModel.StarSearchParams} params
 * @return {map[string]interface{}}
 */
func GetStarList(ctx context.Context, params searchModel.StarSearchParams) (map[string]interface{}, error) {
	ret, err := starData.GetStarList(ctx, params)
	if err != nil {
		return nil, err
	}
	pages, err := starData.GetStarPagesCount(ctx, params)
	if err != nil {
		return nil, err
	}
	res, err := starData.FormatStarInfo(ctx, params, ret, pages)
	if err != nil {
		return nil, err
	}
	return res, nil
}
