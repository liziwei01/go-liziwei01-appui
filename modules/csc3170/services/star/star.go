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
	starModel "go-liziwei01-appui/modules/csc3170/model/star"
)

/**
 * @description: 插入用户数据后台服务层处理逻辑
 * @param {starModel.UserInfo} params
 * @return {map[string]interface{}}
 */
func InsertUser(ctx context.Context, params starModel.UserInfo) error {
	err := starData.InsertUser(ctx, params)
	if err != nil {
		return err
	}
	return nil
}

/**
 * @description: 获取用户数据后台服务层处理逻辑
 * @param {starModel.UserInfo} params
 * @return {map[string]interface{}}
 */
func GetUserList(ctx context.Context, params searchModel.UserSearchParams) (map[string]interface{}, error) {
	ret, err := starData.GetUserList(ctx, params)
	if err != nil {
		return nil, err
	}
	pages, err := starData.GetUserPagesCount(ctx, params)
	if err != nil {
		return nil, err
	}
	res, err := starData.FormatUserInfo(ctx, params, ret, pages)
	if err != nil {
		return nil, err
	}
	return res, nil
}
