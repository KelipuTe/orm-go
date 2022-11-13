package clause

// S6Query I9QueryBuilder.F8BuildQuery 的结果
type S6Query struct {
	// SQLString 带有占位符的 SQL 语句
	SQLString string
	// S5Value SQL 语句中占位符对应的值
	S5Value []any
}

// I9QueryBuilder 接口抽象：查询构造器
// Builder 设计模式
type I9QueryBuilder interface {
	// F8BuildQuery 方法抽象：构造 S6Query
	F8BuildQuery() (*S6Query, error)
}
