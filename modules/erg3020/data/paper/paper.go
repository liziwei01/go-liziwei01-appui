/*
 * @Author: liziwei01
 * @Date: 2021-04-19 15:00:00
 * @LastEditTime: 2022-02-27 17:20:03
 * @LastEditors: liziwei01
 * @Description: 搜索论文服务数据层：要从数据库获取或者要写入的数据在这里处理
 * @FilePath: /github.com/liziwei01/go-liziwei01-appui/modules/erg3020/data/paper/paper.go
 */
package paper

import (
	"context"
	"sort"
	"time"

	LD "github.com/liziwei01/go-liziwei01-appui/library/string_distance"
	paperDao "github.com/liziwei01/go-liziwei01-appui/modules/erg3020/dao/paper"
	paperModel "github.com/liziwei01/go-liziwei01-appui/modules/erg3020/model/paper"
	searchModel "github.com/liziwei01/go-liziwei01-appui/modules/erg3020/model/search"

	"github.com/gogf/gf/util/gconv"
)

/**
 * @description: 搜索论文服务后台数据层处理逻辑
 * @param {searchModel.PaperSearchParams} params
 * @return {[]paperModel.PaperInfo}
 */
func GetPaperList(ctx context.Context, params searchModel.PaperSearchParams) ([]paperModel.PaperInfo, error) {
	params.PageLength = params.PageLength * 2
	res, err := paperDao.GetPaperList(ctx, params)
	if err != nil {
		return make([]paperModel.PaperInfo, 0), err
	}
	return ScoreSimilarity(ctx, params, res), nil
}

func GetPaper(ctx context.Context, params searchModel.PaperSearchParams) (string, error) {
	res, err := paperDao.GetPaper(ctx, params)
	if err != nil {
		return "", err
	}
	return res, nil
}

// 排序
func ScoreSimilarity(ctx context.Context, params searchModel.PaperSearchParams, papers []paperModel.PaperInfo) []paperModel.PaperInfo {
	var (
		key string
		val string
	)
	for k, v := range papers {
		if params.Authors != "" {
			key = params.Authors
			val = v.Authors
		} else if params.Title != "" {
			key = params.Title
			val = v.Title
		} else if params.Journal != "" {
			key = params.Journal
			val = v.Journal
		} else {
			if gconv.Int(params.PageLength/2) < len(papers) {
				return papers[:params.PageLength/2]
			}
			return papers
		}
		// 计算编辑距离
		papers[k].ScoreSimilarity = LD.Ld(key, val, true)
	}
	// 按照引用分数减去编辑距离排序，两者比重各占约50%
	sort.SliceStable(papers, func(i, j int) bool {
		if (int(papers[i].Score) - papers[i].ScoreSimilarity*100) >
			(int(papers[i].Score) - papers[i].ScoreSimilarity*100) {
			return true
		}
		return false
	})
	if gconv.Int(params.PageLength/2) < len(papers) {
		return papers[:params.PageLength/2]
	}
	return papers
}

/**
 * @description: 计数处理
 * @param {searchModel.PaperSearchParams} params
 * @return {map[string]interface{}}
 */
func GetPaperPagesCount(ctx context.Context, params searchModel.PaperSearchParams) (int64, error) {
	count, err := paperDao.GetPaperPagesCount(ctx, params)
	if err != nil {
		return 0, err
	}
	return count, nil
}

/**
 * @description: 处理数据库获取的数据
 * @param {searchModel.PaperSearchParams} params
 * @param {[]paperModel.PaperInfo} papersInfo
 * @param {int64} count
 * @return {map[string]interface{}}
 */
func FormatPaperInfo(ctx context.Context, params searchModel.PaperSearchParams, papersInfo []paperModel.PaperInfo, count int64) (map[string]interface{}, error) {
	var (
		res []map[string]interface{}
	)
	for _, v := range papersInfo {
		timeStr := time.Unix(v.PublishTime, 0).Format("2006")
		res = append(res, map[string]interface{}{
			"index_number": v.IndexNumber,
			"title":        v.Title,
			"authors":      v.Authors,
			"journal":      v.Journal,
			"publish_time": timeStr,
			"references":   v.References,
			"total_cites":  v.TotalCites,
			"score":        int(v.Score) - v.ScoreSimilarity*100,
		})
	}
	return map[string]interface{}{
		"list":  res,
		"count": count,
	}, nil
}
