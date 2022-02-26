/*
 * @Author: liziwei01
 * @Date: 2021-04-19 15:00:00
 * @LastEditTime: 2022-02-26 19:49:58
 * @LastEditors: liziwei01
 * @Description: 搜索论文服务数据层：要从数据库获取或者要写入的数据在这里处理
 * @FilePath: /github.com/liziwei01/go-liziwei01-appui/modules/erg3020/data/paper/paper.go
 */
package paper

import (
	"context"
	"io"
	"net/http"
	"regexp"

	"github.com/gogf/gf/util/gconv"
	paperModel "github.com/liziwei01/go-liziwei01-appui/modules/crawler/model/paper"
	searchModel "github.com/liziwei01/go-liziwei01-appui/modules/crawler/model/search"

	crawlerDao "github.com/liziwei01/go-liziwei01-appui/modules/crawler/dao"

	"github.com/liziwei01/go-liziwei01-library/model/logit"
)

var (
	size   = 25
	regpdf = "<a href=\"https://arxiv.org/pdf/[0-9]*.[0-9]*\">"
)

/**
 * @description: 搜索论文服务后台数据层处理逻辑
 * @param {searchModel.PaperSearchParams} params
 * @return {[]paperModel.PaperInfo}
 */
func CrawlPaper(ctx context.Context, params searchModel.PaperSearchParams) error {
	go crawlPaper(context.Background(), params)
	return nil
}

func crawlPaper(ctx context.Context, params searchModel.PaperSearchParams) {
	start := 0
	continueCrawl := true
	for {
		if !continueCrawl {
			break
		}
		resp, err := http.Get("https://arxiv.org/?searchtype=all&size=" + gconv.String(size) + "&query=" + params.Title + "&start=" + gconv.String(start))
		if err != nil {
			logit.Logger.Error(err)
			break
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			logit.Logger.Error(err)
			break
		}
		pdfUrls := regexp.MustCompile(regpdf).FindAllStringSubmatch(string(body), -1)

		if len(pdfUrls) == 0 {
			break
		}
		start += size

		for _, url := range pdfUrls {
			url := url[0][9:len(url[0])-2] + ".pdf"
			resp2, err := http.Get(url)
			body2, err := io.ReadAll(resp2.Body)
			if err != nil {
				logit.Logger.Error(err)
				break
			}
			paper := paperModel.PaperInfo{
				Title:   url[22:],
				Ref:     params.Title,
				Content: gconv.String(body2),
			}
			crawlerDao.AddPaper(ctx, paper)
			resp2.Body.Close()
		}
	}
}
