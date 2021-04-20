package dao

import (
	"context"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gogf/gf/util/gconv"
	"github.com/jmoiron/sqlx"

	"go-liziwei01-appui/modules/erg3020/constant"
)

// SelectBuilder 默认的select sql builder
type SelectBuilder struct {
	table  string
	where  map[string]interface{}
	fields []string
}

type Client struct {
	Name string
}

var clients []*Client

func InitClients() {
	clients = append(clients, &Client{
		Name: constant.SERVICE_CONF_DB_NEWAPP_LIZIWEI,
	})
}

// GetClient 获取创建
func GetClient(ctx context.Context, serviceName string) (*Client, error) {
	for _, v := range clients {
		if v.Name == serviceName {
			return v, nil
		}
	}
	return &Client{}, fmt.Errorf("cannot find db")
}

func (dao *Client) Query(ctx context.Context, tableName string, where map[string]interface{}, columns []string, data interface{}) error {
	builder := NewSelectBuilder(tableName, where, columns)
	err := QueryWithBuilder(ctx, dao, builder, data)
	if err != nil {
		return err
	}
	return nil
}

func NewSelectBuilder(table string, where map[string]interface{}, fields []string) *SelectBuilder {
	return &SelectBuilder{
		table:  table,
		where:  where,
		fields: fields,
	}
}

// QueryWithBuilder 传入一个 SQLBuilder 并执行 QueryContext
func QueryWithBuilder(ctx context.Context, client *Client, builder *SelectBuilder, data interface{}) error {
	query := QueryCompiler(ctx, client, builder)
	db, err := sqlx.Connect(constant.DB_DRIVER_NAME_MYSQL, "erg3020:liziwei01@tcp(10.30.202.213:3306)/"+client.Name)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	err = db.Select(data, query)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func QueryCompiler(ctx context.Context, client *Client, builder *SelectBuilder) string {
	var (
		limitPar   []uint
		orderbyPar string
	)

	query := "SELECT"
	for k, v := range builder.fields {
		if k == 0 {
			query = query + " " + v
		} else {
			query = ", " + query + " " + v
		}
	}
	query = query + " FROM " + builder.table + " WHERE "
	count := 0
	for k, v := range builder.where {
		if k[0:1] == "_" {
			if k[1:] == "limit" {
				limitPar = v.([]uint)
			} else if k[1:] == "orderby" {
				orderbyPar = v.(string)
			}
		} else {
			if v.(string) == "" || v.(string) == "''" || v.(string) == "'%%'" {
				continue
			}
			if count == 0 {
				query = query + k + v.(string)
				count++
			} else {
				query = query + " and " + k + v.(string)
			}
		}
	}
	if orderbyPar != "" {
		query = query + " ORDER BY " + orderbyPar
	}
	if len(limitPar) != 0 {
		query = query + " LIMIT " + gconv.String(limitPar[0]) + "," + gconv.String(limitPar[1])
	}
	// query = query + ";"
	log.Printf("query: %s\n", query)
	return query
}
