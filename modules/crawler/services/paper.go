/*
 * @Author: liziwei01
 * @Date: 2021-04-19 15:00:00
 * @LastEditTime: 2022-02-26 17:02:43
 * @LastEditors: liziwei01
 * @Description: 搜索论文服务服务层：这一层提供完整的服务
 * @FilePath: /github.com/liziwei01/go-liziwei01-appui/modules/erg3020/services/paper/paper.go
 */
package paper

import (
	"context"

	crawlData "github.com/liziwei01/go-liziwei01-appui/modules/crawler/data"
	searchModel "github.com/liziwei01/go-liziwei01-appui/modules/crawler/model/search"
)

/**
 * @description: 搜索论文服务后台服务层处理逻辑
 * @param {searchModel.PaperSearchParams} params
 * @return {map[string]interface{}}
 */
func CrawlPaper(ctx context.Context, params searchModel.PaperSearchParams) error {
	err := crawlData.CrawlPaper(ctx, params)
	if err != nil {
		return err
	}
	return nil
}
