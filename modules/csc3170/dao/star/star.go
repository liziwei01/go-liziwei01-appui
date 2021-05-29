/*
 * @Author: 		liziwei01
 * @Date: 			2021-04-19 15:00:00
 * @LastEditTime: 	2021-04-19 15:00:00
 * @LastEditors: 	liziwei01
 * @Description: 	搜索明星服务数据库层：在这里访问数据库获取数据
 * @FilePath: 		github.com/liziwei01/go-liziwei01-appui/modules/csc3170/dao/star/star.go
 */
package star

import (
	"context"
	"log"

	"github.com/liziwei01/go-liziwei01-appui/modules/csc3170/constant"
	searchModel "github.com/liziwei01/go-liziwei01-appui/modules/csc3170/model/search"
	starModel "github.com/liziwei01/go-liziwei01-appui/modules/csc3170/model/star"

	baseDao "github.com/liziwei01/go-liziwei01-library/model/mysql"
)

const (
	// 用户信息表
	USER_TABLE_NAME = "tb_star_user_info"
)

/**
 * @description: 搜索明星服务后台数据层处理逻辑
 * @param {searchModel.StarSearchParams} params
 * @return {[]starModel.StarInfo}
 */
func InsertUser(ctx context.Context, params starModel.UserInfo) error {
	client, err := baseDao.GetMysqlClient(ctx, constant.SERVICE_CONF_DB_NEWAPP_LIZIWEI)
	if err != nil {
		log.Printf("csc3170.dao.InsertUser GetMysqlClient failed with err: %s\n", err.Error())
		return err
	}
	err = client.Insert(ctx, USER_TABLE_NAME, map[string]interface{}{
		"user_id":  params.UserId,
		"name":     params.UserName,
		"password": params.Password,
	})
	if err != nil {
		log.Printf("csc3170.dao.InsertUser Insert failed with err: %s\n", err.Error())
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
	var (
		res      []starModel.UserInfo
		intStart = (params.PageIndex - 1) * params.PageLength
	)
	// 获取Client
	client, err := baseDao.GetMysqlClient(ctx, constant.SERVICE_CONF_DB_NEWAPP_LIZIWEI)
	if err != nil {
		log.Printf("csc3170.dao.GetUserList GetMysqlClient failed with err: %s\n", err.Error())
		return make([]starModel.UserInfo, 0), err
	}
	// 搜索关键词
	where := map[string]interface{}{
		"_orderby":  "user_id asc",
		"_limit":    []uint{intStart, params.PageLength},
		"name like": params.UserName,
	}
	// 需求的列
	columns := []string{"*"}
	// 使用Client进行搜索并在res中获取返回数据
	err = client.Query(ctx, USER_TABLE_NAME, where, columns, &res)
	if err != nil {
		log.Printf("csc3170.dao.GetUserList Query failed with err: %s\n", err.Error())
		return make([]starModel.UserInfo, 0), err
	}
	return res, nil
}

func GetUserPagesCount(ctx context.Context, params searchModel.UserSearchParams) (int64, error) {
	var (
		userCount = make([]struct {
			UserCount int64 `db:"count"`
		}, 1)
	)
	client, err := baseDao.GetMysqlClient(ctx, constant.SERVICE_CONF_DB_NEWAPP_LIZIWEI)
	if err != nil {
		log.Printf("csc3170.dao.GetUserPagesCount GetMysqlClient failed with err: %s\n", err.Error())
		return 0, err
	}
	where := map[string]interface{}{
		"name like": params.UserName,
	}
	columns := []string{"count(user_id) as count"}
	err = client.Query(ctx, USER_TABLE_NAME, where, columns, &userCount)
	if err != nil {
		log.Printf("csc3170.dao.GetUserPagesCount Query failed with err: %s\n", err.Error())
		return 0, err
	}
	return userCount[1].UserCount, nil
}
