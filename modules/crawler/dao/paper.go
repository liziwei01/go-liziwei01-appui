/*
 * @Author: liziwei01
 * @Date: 2021-04-19 15:00:00
 * @LastEditTime: 2022-02-26 20:04:20
 * @LastEditors: liziwei01
 * @Description: 搜索论文服务数据库层：在这里访问数据库获取数据
 * @FilePath: /github.com/liziwei01/go-liziwei01-appui/modules/erg3020/dao/paper/paper.go
 */
package paper

import (
	"context"

	"github.com/liziwei01/go-liziwei01-appui/modules/crawler/constant"
	paperModel "github.com/liziwei01/go-liziwei01-appui/modules/crawler/model/paper"

	"github.com/liziwei01/go-liziwei01-library/model/logit"
	baseDao "github.com/liziwei01/go-liziwei01-library/model/mysql"
)

const (
	// 论文信息表
	PAPER_TABLE_NAME = "tb_gesture_teleoperation_paper_info"
)

func AddPaper(ctx context.Context, param paperModel.PaperInfo) error {
	client, err := baseDao.GetMysqlClient(ctx, constant.SERVICE_CONF_DB_NEWAPP_LIZIWEI)
	if err != nil {
		return err
	}
	maps := map[string]interface{}{
		"title":   param.Title,
		"ref":     param.Ref,
		"content": param.Content,
	}
	err = client.Insert(ctx, PAPER_TABLE_NAME, maps)
	if err != nil {
		logit.Logger.Error("AddPaper", "err", err)
		return err
	}
	logit.Logger.Info("AddPaper", "success", param)
	return nil
}
