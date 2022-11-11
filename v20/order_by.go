package v20

// OrderBy 对应查询语句里的 order by
type OrderBy struct {
	// 字段
	field S6Column
	// order 排序规则：ASC；DESC
	order string
}

// Asc 升序
func Asc(n string) OrderBy {
	return OrderBy{
		field: S6Column{name: n},
		order: "ASC",
	}
}

// Desc 降序
func Desc(n string) OrderBy {
	return OrderBy{
		field: S6Column{name: n},
		order: "DESC",
	}
}
