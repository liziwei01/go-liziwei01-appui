/*
 * @Author: 		liziwei01
 * @Date: 			2021-04-19 15:00:00
 * @LastEditTime: 	2021-04-19 15:00:00
 * @LastEditors: 	liziwei01
 * @Description: 	搜索明星服务数据层：要从数据库获取或者要写入的数据在这里处理
 * @FilePath: 		github.com/liziwei01/go-liziwei01-appui/modules/csc3170/data/star/star.go
 */
package star

import (
	"context"

	starDao "github.com/liziwei01/go-liziwei01-appui/modules/csc3170/dao/star"
	searchModel "github.com/liziwei01/go-liziwei01-appui/modules/csc3170/model/search"
	starModel "github.com/liziwei01/go-liziwei01-appui/modules/csc3170/model/star"
)

/**
 * @description: 插入用户数据后台数据层处理逻辑
 * @param {starModel.UserInfo} params
 * @return {[]starModel.StarInfo}
 */
func InsertUser(ctx context.Context, params starModel.UserInfo) error {
	err := starDao.InsertUser(ctx, params)
	if err != nil {
		return err
	}
	return nil
}

/**
 * @description: 搜索论文服务后台数据层处理逻辑
 * @param {searchModel.UserSearchParams} params
 * @return {[]starModel.UserInfo}
 */
func GetUserList(ctx context.Context, params searchModel.UserSearchParams) ([]starModel.UserInfo, error) {
	res, err := starDao.GetUserList(ctx, params)
	if err != nil {
		return make([]starModel.UserInfo, 0), err
	}
	return res, nil
}

/**
 * @description: 计数处理
 * @param {starModel.UserInfo} params
 * @return {map[string]interface{}}
 */
func GetUserPagesCount(ctx context.Context, params searchModel.UserSearchParams) (int64, error) {
	count, err := starDao.GetUserPagesCount(ctx, params)
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
func FormatUserInfo(ctx context.Context, params searchModel.UserSearchParams, userInfo []starModel.UserInfo, count int64) (map[string]interface{}, error) {
	var (
		list []map[string]interface{}
	)
	for _, v := range userInfo {
		list = append(list, map[string]interface{}{
			"user_id":  v.UserId,
			"name":     v.UserName,
			"password": v.Password,
		})
	}
	return map[string]interface{}{
		"list":   list,
		"count":  count,
		"errmsg": "success",
	}, nil
}
