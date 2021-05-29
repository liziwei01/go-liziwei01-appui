/*
 * @Author: liziwei01
 * @Date: 2021-04-29 15:14:24
 * @LastEditors: liziwei01
 * @LastEditTime: 2021-05-30 02:27:31
 * @Description: file content
 * @FilePath: /github.com/liziwei01/go-liziwei01-appui/script/erg3020/readcsv/readcsv.go
 */
package readcsv

import (
	"bufio"
	"context"
	"encoding/csv"
	"io"
	"os"
	"time"

	"github.com/gogf/gf/util/gconv"

	paperDao "github.com/liziwei01/go-liziwei01-appui/modules/erg3020/dao/paper"
	paperModel "github.com/liziwei01/go-liziwei01-appui/modules/erg3020/model/paper"
)

const (
	STD_TIME         = "2006-01-02 15:04:05"
	MAXIMUM_CSV_SIZE = 99999
)

func AddBatchAsync(ctx context.Context, params []paperModel.PaperInfo) error {
	for _, param := range params {
		err := paperDao.AddPaper(ctx, param)
		if err != nil {
			return err
		}
	}
	return nil
}

func ParseBatchCsv(ctx context.Context, csvName string) ([]paperModel.PaperInfo, error) {
	var (
		params []paperModel.PaperInfo
	)
	csvFile, err := os.Open(csvName)
	if err != nil {
		return make([]paperModel.PaperInfo, 0), err
	}
	// err = checkCsvRows(ctx, csvFile)
	if err != nil {
		return make([]paperModel.PaperInfo, 0), err
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return make([]paperModel.PaperInfo, 0), err
		}
		if err != nil {
			return make([]paperModel.PaperInfo, 0), err
		}
		params = append(params, paperModel.PaperInfo{
			IndexNumber: gconv.Int64(line[0]),
			Title:       line[1],
			Authors:     line[2],
			PublishTime: year2TimeStamp(line[3]),
			Journal:     line[4],
			References:  line[5],
			TotalCites:  gconv.Int64(line[6]),
			Score:       gconv.Int64(line[7]),
		})
	}
	return params, nil
}

func timeStr2Int64(layout string, timeStr string) int64 {
	time, _ := time.Parse(layout, timeStr)
	return time.Local().Unix()
}

func year2TimeStamp(timeStr string) int64 {
	return timeStr2Int64("2006", timeStr)
}
