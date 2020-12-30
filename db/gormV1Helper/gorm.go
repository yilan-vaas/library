package gormV1Helper

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"strings"
	"time"
)

type Base struct {
	ID        uint       `gorm:"primary_key" json:"id" form:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

type QueryOptions struct {
	Conditions map[string]interface{}
	Having     map[string]interface{}
	Columns    string
	Order      string
	Group      string
	Preload    []string
	Offset     int
	Limit      int
}

func (q *QueryOptions) GetWhere() (query string, args []interface{}) {
	glue := ""
	for key, value := range q.Conditions {
		query += glue + key
		if values, assert := value.([]interface{}); assert {
			args = append(args, values...)
		} else {
			args = append(args, value)
		}
		glue = " AND "
	}
	return
}

func (q *QueryOptions) GetHaving() (query string, arg interface{}) {
	for key, value := range q.Having {
		query = key
		arg = value
		break
	}
	return
}

type Query struct {
	Fn func(*QueryOptions)
}

func Columns(cols ...string) Query {
	return Query{func(ops *QueryOptions) {
		ops.Columns = strings.Join(cols, ",")
	}}
}

func Conditions(c map[string]interface{}) Query {
	return Query{func(ops *QueryOptions) {
		ops.Conditions = c
	}}
}

func Order(o string) Query {
	return Query{func(ops *QueryOptions) {
		ops.Order = o
	}}
}

func Group(g string) Query {
	return Query{func(ops *QueryOptions) {
		ops.Group = g
	}}
}

func Having(h map[string]interface{}) Query {
	return Query{func(ops *QueryOptions) {
		ops.Having = h
	}}
}

func Preload(pload ...string) Query {
	return Query{func(ops *QueryOptions) {
		ops.Preload = pload
	}}
}

func Paginator(page, size int) Query {
	return Query{func(ops *QueryOptions) {
		ops.Offset = (page - 1) * size
		if ops.Offset < 0 {
			ops.Offset = 0
		}
		ops.Limit = size
	}}
}

type Table interface {
	DB() *gorm.DB
}

func QueryList(inf Table, result interface{}, total *int64, args ...Query) error {
	queryOpts := &QueryOptions{Offset: 0, Limit: -1}
	db := inf.DB().Model(inf)
	if len(args) > 0 {
		for _, opt := range args {
			opt.Fn(queryOpts)
		}
		if queryOpts.Columns != "" {
			db = db.Select(queryOpts.Columns)
		}
		if len(queryOpts.Conditions) > 0 {
			q, ags := queryOpts.GetWhere()
			db = db.Where(q, ags...)
		}
		if len(queryOpts.Preload) > 0 {
			for _, pload := range queryOpts.Preload {
				db = db.Preload(pload)
			}
		}
		if total != nil {
			db.Group(queryOpts.Group).Count(total)
		}
		db = db.Group(queryOpts.Group).Order(queryOpts.Order).Offset(queryOpts.Offset).Limit(queryOpts.Limit)
	}
	return db.Find(result).Error
}

type MYSQL interface {
	Addr() string
	TablePrefix() string
	SessionOn() bool
	LogOn() bool
	Hooks(*gorm.DB)
}

func MysqlConnect(config MYSQL) *gorm.DB {
	conn, err := gorm.Open("mysql", config.Addr())
	if err != nil {
		panic(err.Error())
	}
	conn.LogMode(config.LogOn())
	conn.SingularTable(true)
	conn.InstantSet("table_prefix", config.TablePrefix())
	if config.SessionOn() {
		conn.Set("gorm:table_options", "ENGINE=MyISAM DEFAULT CHARSET=utf8").AutoMigrate(&Session{})
	}
	config.Hooks(conn) //诸如连接池之类的设定可以放在这里，helper只提供基础和必要设置
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		prefix, ok := db.Get("table_prefix")
		if ok && strings.HasPrefix(defaultTableName, prefix.(string)) {
			return defaultTableName
		}
		return prefix.(string) + defaultTableName
	}
	return conn
}
