package mysql

import (
	"context"
	"fmt"
	"log"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gogf/gf/util/gconv"
	"github.com/jmoiron/sqlx"
)

const (
	SERVICE_CONF_DB_NEWAPP_LIZIWEI = "db_liziwei01"
	DB_DRIVER_NAME_MYSQL           = "mysql"
)

// SelectBuilder 默认的select sql builder
type SelectBuilder struct {
	table  string
	where  map[string]interface{}
	fields []string
}

// InsertBuilder 默认的select sql builder
type InsertBuilder struct {
	table string
}

type Client struct {
	Name string
}

var clients []*Client

func InitClients() {
	clients = append(clients, &Client{
		Name: SERVICE_CONF_DB_NEWAPP_LIZIWEI,
	})
}

// GetMysqlClient 获取创建
func GetMysqlClient(ctx context.Context, serviceName string) (*Client, error) {
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

func (dao *Client) Insert(ctx context.Context, tableName string, data map[string]interface{}) error {
	builder := NewInsertBuilder(tableName)
	err := InsertWithBuilder(ctx, dao, builder, data)
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

func NewInsertBuilder(table string) *InsertBuilder {
	return &InsertBuilder{
		table: table,
	}
}

// QueryWithBuilder 传入一个 SQLBuilder 并执行 QueryContext
func QueryWithBuilder(ctx context.Context, client *Client, builder *SelectBuilder, data interface{}) error {
	query := QueryCompiler(ctx, client, builder)
	db, err := sqlx.Connect(DB_DRIVER_NAME_MYSQL, "work:liziwei01@tcp(localhost:3306)/"+client.Name)
	if err != nil {
		return err
	}
	err = db.Select(data, query)
	if err != nil {
		return err
	}
	return nil
}

// InsertWithBuilder 传入一个 SQLBuilder 并执行 QueryContext
func InsertWithBuilder(ctx context.Context, client *Client, builder *InsertBuilder, data map[string]interface{}) error {
	query := InsertCompiler(ctx, client, builder, data)
	db, err := sqlx.Connect(DB_DRIVER_NAME_MYSQL, "work:liziwei01@tcp(localhost:3306)/"+client.Name)
	if err != nil {
		return err
	}
	_, err = db.Queryx(query)
	if err != nil {
		return err
	}
	return nil
}

func beforeCompiler(ctx context.Context, builder *SelectBuilder) *SelectBuilder {
	var (
		equalSign = false
	)
	for k, v := range builder.where {
		if k[0:1] == "_" || len(gconv.String(v)) == 0 {
			continue
		}
		if len(k) > 4 && k[len(k)-4:] == "like" {
			builder.where[k] = "%" + gconv.String(v) + "%"
		} else if k[len(k)-1:] == "=" || k[len(k)-1:] == ">" || k[len(k)-1:] == "<" {
		} else {
			equalSign = true
		}
		if reflect.TypeOf(v) == reflect.TypeOf("") {
			builder.where[k] = "'" + gconv.String(builder.where[k]) + "'"
		}
		if equalSign {
			builder.where[k] = "= " + gconv.String(builder.where[k])
		}
	}
	return builder
}

func QueryCompiler(ctx context.Context, client *Client, builder *SelectBuilder) string {
	var (
		limitPar   []uint
		orderbyPar string
		query      = "SELECT"
	)
	builder = beforeCompiler(ctx, builder)
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
		// _特殊处理
		if k[0:1] == "_" {
			if k[1:] == "limit" {
				limitPar = v.([]uint)
			} else if k[1:] == "orderby" {
				orderbyPar = gconv.String(v)
			}
		} else {
			if gconv.String(v) == "" || gconv.String(v) == "''" || gconv.String(v) == "'%%'" {
				continue
			}
			if count == 0 {
				query = query + k + " " + gconv.String(v)
				count++
			} else {
				query = query + " and " + k + " " + gconv.String(v)
			}
		}
	}
	if orderbyPar != "" {
		query = query + " ORDER BY " + orderbyPar
	}
	if len(limitPar) != 0 {
		query = query + " LIMIT " + gconv.String(limitPar[0]) + "," + gconv.String(limitPar[1])
	}
	log.Printf("query: %s\n", query)
	return query
}

func InsertCompiler(ctx context.Context, client *Client, builder *InsertBuilder, data map[string]interface{}) string {
	var (
		query     = "INSERT INTO " + builder.table + " ("
		prefixLen = len(query)
		keysLen   = 0
	)

	for k, v := range data {
		query = query[0:prefixLen+keysLen] + k + ", " + query[prefixLen+keysLen:]
		keysLen = keysLen + len(k) + len(", ")
		query = query + gconv.String(v) + ", "
	}

	query = query[0:prefixLen+keysLen-2] + ") VALUES (" + query[prefixLen+keysLen:len(query)-2] + ")"
	log.Printf("query: %s\n", query)
	return query
}
