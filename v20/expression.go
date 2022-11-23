package v20

// i9Expression 对应 i9Expr
// SELECT Statement 里的 i9Expr
// 可以是：列、聚合函数、查询条件、值
type i9Expression interface {
	// f8BuildExpression 构造 i9Expr
	f8BuildExpression(p7s6Builder *s6QueryBuilder) error
}

// f8NewI9Expression 把输入转换成 i9Expr
func f8NewI9Expression(input any) i9Expression {
	switch input.(type) {
	case i9Expression:
		// 如果是 i9Expr 就断言一下丢回去
		return input.(i9Expression)
	default:
		// 如果不是 i9Expr 就转换成 value
		return F8NewS6Value(input)
	}
}
