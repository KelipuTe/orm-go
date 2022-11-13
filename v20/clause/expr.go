package clause

// I9Expr 对应查询语句里的语句或者语句的一部分
// 即 SELECT Statement 里的 expr
type I9Expr interface {
	F8Expr()
}

// NewI9Expr 把输入转换成查询语句里语句
func NewI9Expr(input any) I9Expr {
	switch input.(type) {
	case I9Expr:
		// 如果是语句就什么也不做
		return input.(I9Expr)
	default:
		// 如果不是语句就转换成占位符对应的参数
		return NewS6Value(input)
	}
}
