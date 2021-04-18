/*
 * @Author: 		liziwei01
 * @Date: 			2021-04-19 15:00:00
 * @LastEditTime: 	2021-04-19 15:00:00
 * @LastEditors: 	liziwei01
 * @Description: 	搜索论文服务服务层
 * @FilePath: 		go-liziwei01-appui/modules/erg3020/services/paper/paper.go
 */
package paper

import(
	"fmt"
	"context"

	searchModel "go-liziwei01-appui/modules/erg3020/model/search"
	paperData "go-liziwei01-appui/modules/erg3020/data/paper"
)

/**
 * @description: 搜索论文服务后台服务层处理逻辑
 * @param {searchModel.PaperSearchParams} params
 * @return {map[string]interface{}}
 */
func GetPaperList(ctx context.Context, params searchModel.PaperSearchParams) (map[string]interface{}, error) {
	fmt.Println("service->GetPaperList")
	ret, err := paperData.GetPaperList(ctx, params)
	if err != nil {
		return nil, err
	}
	pages, err := paperData.GetPaperPagesCount(ctx, params)
	if err != nil {
		return nil, err
	}
	res, err := paperData.FormatPaperInfo(ctx, params, ret, pages)
	if err != nil {
		return nil, err
	}
	return res, nil
}