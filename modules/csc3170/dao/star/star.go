/*
 * @Author: 		liziwei01
 * @Date: 			2021-04-19 15:00:00
 * @LastEditTime: 	2021-04-19 15:00:00
 * @LastEditors: 	liziwei01
 * @Description: 	搜索明星服务数据库层：在这里访问数据库获取数据
 * @FilePath: 		go-liziwei01-appui/modules/csc3170/dao/star/star.go
 */
package star

import (
	"context"
	"fmt"
	"log"

	"github.com/gogf/gf/util/gconv"

	"go-liziwei01-appui/modules/csc3170/constant"
	baseDao "go-liziwei01-appui/modules/csc3170/dao"
	starModel "go-liziwei01-appui/modules/csc3170/model/star"
	searchModel "go-liziwei01-appui/modules/csc3170/model/search"
)

const (
	// 论文信息表
	PAPER_TABLE_NAME = "tb_gesture_teleoperation_star_info"
)

/**
 * @description: 搜索明星服务后台数据层处理逻辑
 * @param {searchModel.StarSearchParams} params
 * @return {[]starModel.StarInfo}
 */
func GetStarList(ctx context.Context, params searchModel.StarSearchParams) ([]starModel.StarInfo, error) {
	var res []starModel.StarInfo
	var intStart = (params.PageIndex - 1) * params.PageLength
	client, err := baseDao.GetClient(ctx, constant.SERVICE_CONF_DB_NEWAPP_LIZIWEI)
	if err != nil {
		return make([]starModel.StarInfo, 0), err
	}
	where := map[string]interface{}{
		"_orderby":       "point desc",
		"_limit":         []uint{intStart, params.PageLength},
		"title like":     "'%" + params.Title + "%'",
		"author like":    "'%" + params.Authors + "%'",
		"publish_time>=": gconv.String(params.StartTime),
		"publish_time<=": gconv.String(params.EndTime),
		"journal like":   "'%" + params.Journal + "%'",
	}
	columns := []string{"*"}
	err = client.Query(ctx, PAPER_TABLE_NAME, where, columns, &res)
	if err != nil {
		msg := fmt.Sprintf("[GetStarList] -> get star list from db failed")
		log.Fatalln(msg)
		return make([]starModel.StarInfo, 0), err
	}
	return res, nil
}

func GetStarPagesCount(ctx context.Context, params searchModel.StarSearchParams) (int64, error) {
	// to do

	return 0, nil
}
