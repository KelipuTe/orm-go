package clause

// I9Expression 对应表达式（列、聚合函数、查询条件、值）
// 即 SELECT Statement 里的 expr
type I9Expression interface {
	F8Expression()
}

// F8NewI9Expression 把输入转换成表达式
func F8NewI9Expression(input any) I9Expression {
	switch input.(type) {
	case I9Expression:
		// 如果是语句就什么也不做
		return input.(I9Expression)
	default:
		// 如果不是语句就转换成占位符对应的参数
		return F8NewS6Value(input)
	}
}
