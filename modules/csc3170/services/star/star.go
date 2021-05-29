/*
 * @Author: liziwei01
 * @Date: 2021-04-19 15:00:00
 * @LastEditTime: 2021-05-30 02:25:53
 * @LastEditors: liziwei01
 * @Description: 搜索明星服务服务层：这一层提供完整的服务
 * @FilePath: /github.com/liziwei01/go-liziwei01-appui/modules/csc3170/services/star/star.go
 */
package star

import (
	"context"

	starData "github.com/liziwei01/go-liziwei01-appui/modules/csc3170/data/star"
	searchModel "github.com/liziwei01/go-liziwei01-appui/modules/csc3170/model/search"
	starModel "github.com/liziwei01/go-liziwei01-appui/modules/csc3170/model/star"
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
	// 从数据库获取前端要求数量的用户名单，如：要求搜索第二页5个人则limit 1,5
	ret, err := starData.GetUserList(ctx, params)
	if err != nil {
		return nil, err
	}
	// 统计符合关键词要求的数据总量交给前端分页
	pages, err := starData.GetUserPagesCount(ctx, params)
	if err != nil {
		return nil, err
	}
	// 处理数据，如：将时间戳转换为便于理解的年月日
	res, err := starData.FormatUserInfo(ctx, params, ret, pages)
	if err != nil {
		return nil, err
	}
	return res, nil
}
