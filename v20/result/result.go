package result

import (
	"database/sql"
	"orm-go/v20/metadata"
)

// I9Result 接口抽象：用数据库返回的查询结果构造结构体
type I9Result interface {
	// F8SetField 方法抽象：将数据库返回的查询结果放到结构体对应的字段上去
	F8SetField(rows *sql.Rows) error
	// Field 返回字段对应的值
	F8GetField(name string) (any, error)
}

// F8NewI9Result 方法抽象：创建一个 I9Result 接口的实例
type F8NewI9Result func(value interface{}, p7s6Model *metadata.S6Model) I9Result
