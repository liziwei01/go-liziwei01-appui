/*
 * @Author: 		liziwei01
 * @Date: 			2021-04-19 15:00:00
 * @LastEditTime: 	2021-04-19 15:00:00
 * @LastEditors: 	liziwei01
 * @Description: 	搜索论文服务数据库层：在这里访问数据库获取数据
 * @FilePath: 		go-liziwei01-appui/modules/erg3020/dao/paper/paper.go
 */
package paper

import (
	"context"
	"fmt"
	"log"

	"github.com/gogf/gf/util/gconv"

	"go-liziwei01-appui/modules/erg3020/constant"
	baseDao "go-liziwei01-appui/modules/erg3020/dao"
	paperModel "go-liziwei01-appui/modules/erg3020/model/paper"
	searchModel "go-liziwei01-appui/modules/erg3020/model/search"
)

const (
	// 论文信息表
	PAPER_TABLE_NAME = "tb_gesture_teleoperation_paper_info"
)

/**
 * @description: 搜索论文服务后台数据层处理逻辑
 * @param {searchModel.PaperSearchParams} params
 * @return {[]paperModel.PaperInfo}
 */
func GetPaperList(ctx context.Context, params searchModel.PaperSearchParams) ([]paperModel.PaperInfo, error) {
	var res []paperModel.PaperInfo
	var intStart = (params.PageIndex - 1) * params.PageLength
	client, err := baseDao.GetClient(ctx, constant.SERVICE_CONF_DB_NEWAPP_LIZIWEI)
	if err != nil {
		return make([]paperModel.PaperInfo, 0), err
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
		msg := fmt.Sprintf("[GetPaperList] -> get paper list from db failed")
		log.Fatalln(msg)
		return make([]paperModel.PaperInfo, 0), err
	}
	return res, nil
}

func GetPaperPagesCount(ctx context.Context, params searchModel.PaperSearchParams) (int64, error) {
	// to do

	return 0, nil
}
