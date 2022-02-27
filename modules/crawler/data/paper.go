/*
 * @Author: liziwei01
 * @Date: 2021-04-19 15:00:00
 * @LastEditTime: 2022-02-27 18:13:20
 * @LastEditors: liziwei01
 * @Description: 搜索论文服务数据层：要从数据库获取或者要写入的数据在这里处理
 * @FilePath: /github.com/liziwei01/go-liziwei01-appui/modules/erg3020/data/paper/paper.go
 */
package paper

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"time"

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
		link := "https://arxiv.org/search/?searchtype=all&size=" + gconv.String(size) + "&query=" + params.Title + "&start=" + gconv.String(start)
		logit.Logger.Info("start search paper " + link)
		resp, err := http.Get(link)
		if err != nil {
			logit.Logger.Error(err)
			break
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			logit.Logger.Error(err)
			break
		}
		pdfUrls := regexp.MustCompile(regpdf).FindAllStringSubmatch(string(body), -1)
		logit.Logger.Info("get " + gconv.String(len(pdfUrls)) + " papers")
		if len(pdfUrls) == 0 {
			break
		}
		start += size

		for _, url := range pdfUrls {
			time.Sleep(time.Second * 60)
			url := url[0][9:len(url[0])-2] + ".pdf"
			logit.Logger.Info("start get paper " + url)
			resp2, err := http.Get(url)
			body2, err := ioutil.ReadAll(resp2.Body)
			if err != nil {
				logit.Logger.Error(err)
				continueCrawl = false
				break
			}
			title := url[22:]
			pwd, err := os.Getwd()
			if err != nil {
				logit.Logger.Error(err)
				continueCrawl = false
				break
			}
			path := pwd + "/data/" + title
			err = ioutil.WriteFile(path, body2, 0644)
			if err != nil {
				logit.Logger.Error(err)
				continueCrawl = false
				break
			}
			paper := paperModel.PaperInfo{
				Title:   title,
				Ref:     params.Title,
				Content: path,
			}
			err = crawlerDao.AddPaper(ctx, paper)
			if err != nil {
				logit.Logger.Error(err)
				continueCrawl = false
				break
			}
			resp2.Body.Close()
			logit.Logger.Info("finish get paper " + url)
		}
		resp.Body.Close()
	}
}
