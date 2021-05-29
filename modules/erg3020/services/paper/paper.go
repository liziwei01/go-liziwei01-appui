/*
 * @Author: liziwei01
 * @Date: 2021-04-19 15:00:00
 * @LastEditTime: 2021-05-30 02:27:21
 * @LastEditors: liziwei01
 * @Description: 搜索论文服务服务层：这一层提供完整的服务
 * @FilePath: /github.com/liziwei01/go-liziwei01-appui/modules/erg3020/services/paper/paper.go
 */
package paper

import (
	"context"
	"time"

	"github.com/gogf/gf/util/gconv"

	paperData "github.com/liziwei01/go-liziwei01-appui/modules/erg3020/data/paper"
	searchModel "github.com/liziwei01/go-liziwei01-appui/modules/erg3020/model/search"
)

/**
 * @description: 搜索论文服务后台服务层处理逻辑
 * @param {searchModel.PaperSearchParams} params
 * @return {map[string]interface{}}
 */
func GetPaperList(ctx context.Context, params searchModel.PaperSearchParams) (map[string]interface{}, error) {
	ret, err := paperData.GetPaperList(ctx, params)
	if err != nil {
		return nil, err
	}
	pages, err := paperData.GetPaperPagesCount(ctx, params)
	if err != nil {
		return nil, err
	}
	// res, err := paperData.FormatPaperInfo(ctx, params, ret, pages)
	// if err != nil {
	// 	return nil, err
	// }
	for k, _ := range ret {
		ret[k].PublishTime = gconv.Int64(time.Unix(ret[k].PublishTime, 0).Format("2006"))
	}
	return map[string]interface{}{
		"list":  ret,
		"count": pages,
	}, nil
}
