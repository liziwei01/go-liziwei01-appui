/*
 * @Author: 		liziwei01
 * @Date: 			2021-04-19 15:00:00
 * @LastEditTime: 	2021-04-19 15:00:00
 * @LastEditors: 	liziwei01
 * @Description: 	搜索论文服务数据库层
 * @FilePath: 		go-liziwei01-appui/modules/erg3020/dao/paper/paper.go
 */
package paper

import (
	"context"

	searchModel "go-liziwei01-appui/modules/erg3020/model/search"
	paperModel "go-liziwei01-appui/modules/erg3020/model/paper"
)

/**
 * @description: 搜索论文服务后台数据层处理逻辑
 * @param {searchModel.PaperSearchParams} params
 * @return {[]paperModel.PaperInfo}
 */
func GetPaperList(ctx context.Context, params searchModel.PaperSearchParams) ([]paperModel.PaperInfo, error) {
	// to do

	return make([]paperModel.PaperInfo, 0), nil
}