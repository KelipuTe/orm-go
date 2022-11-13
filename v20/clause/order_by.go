package clause

// S6OrderBy 对应 SELECT 语句的 ORDER BY
// 即 SELECT Statement 里的 ORDER BY {col_name} [ASC | DESC]
type S6OrderBy struct {
	// S6Column 排序列
	S6Column S6Column
	// OrderString 排序规则：ASC；DESC
	OrderString string
}

// Asc 升序
func Asc(name string) S6OrderBy {
	return S6OrderBy{
		S6Column:    S6Column{Name: name},
		OrderString: "ASC",
	}
}

// Desc 降序
func Desc(name string) S6OrderBy {
	return S6OrderBy{
		S6Column:    S6Column{Name: name},
		OrderString: "DESC",
	}
}
