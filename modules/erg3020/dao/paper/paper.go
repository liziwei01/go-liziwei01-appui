/*
 * @Author: 		liziwei01
 * @Date: 			2021-04-19 15:00:00
 * @LastEditTime: 	2021-04-19 15:00:00
 * @LastEditors: 	liziwei01
 * @Description: 	搜索论文服务数据库层：在这里访问数据库获取数据
 * @FilePath: 		github.com/liziwei01/go-liziwei01-appui/modules/erg3020/dao/paper/paper.go
 */
package paper

import (
	"context"

	"github.com/liziwei01/go-liziwei01-appui/modules/erg3020/constant"
	paperModel "github.com/liziwei01/go-liziwei01-appui/modules/erg3020/model/paper"
	searchModel "github.com/liziwei01/go-liziwei01-appui/modules/erg3020/model/search"

	baseDao "github.com/liziwei01/go-liziwei01-library/model/mysql"
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
	client, err := baseDao.GetMysqlClient(ctx, constant.SERVICE_CONF_DB_NEWAPP_LIZIWEI)
	if err != nil {
		return make([]paperModel.PaperInfo, 0), err
	}
	where := map[string]interface{}{
		"_orderby":        "score desc",
		"_limit":          []uint{intStart, params.PageLength},
		"publish_time >=": params.StartTime,
		"publish_time <=": params.EndTime,
	}
	if params.Title != "" {
		where["title like"] = params.Title
	}
	if params.Authors != "" {
		where["author like"] = params.Authors
	}
	if params.Journal != "" {
		where["journal like"] = params.Journal
	}
	columns := []string{"*"}
	err = client.Query(ctx, PAPER_TABLE_NAME, where, columns, &res)
	if err != nil {
		return make([]paperModel.PaperInfo, 0), err
	}
	return res, nil
}

func GetPaperPagesCount(ctx context.Context, params searchModel.PaperSearchParams) (int64, error) {
	var (
		paperCount = make([]struct {
			Count int64 `db:"count"`
		}, 1)
	)
	client, err := baseDao.GetMysqlClient(ctx, constant.SERVICE_CONF_DB_NEWAPP_LIZIWEI)
	if err != nil {
		return 0, err
	}
	where := map[string]interface{}{
		"publish_time >=": params.StartTime,
		"publish_time <=": params.EndTime,
	}
	if params.Title != "" {
		where["title like"] = params.Title
	}
	if params.Authors != "" {
		where["author like"] = params.Authors
	}
	if params.Journal != "" {
		where["journal like"] = params.Journal
	}
	columns := []string{"count(index_number) as count"}
	err = client.Query(ctx, PAPER_TABLE_NAME, where, columns, &paperCount)
	if err != nil {
		return 0, err
	}
	return paperCount[1].Count, nil
}

func AddPaper(ctx context.Context, param paperModel.PaperInfo) error {
	client, err := baseDao.GetMysqlClient(ctx, constant.SERVICE_CONF_DB_NEWAPP_LIZIWEI)
	if err != nil {
		return err
	}
	maps := map[string]interface{}{
		"index_number": param.IndexNumber,
		"title":        param.Title,
		"author":       param.Authors,
		"publish_time": param.PublishTime,
		"journal":      param.Journal,
		"ref":          param.References,
		"total_cites":  param.TotalCites,
		"score":        param.Score,
	}
	err = client.Insert(ctx, PAPER_TABLE_NAME, maps)
	return nil
}
