package v20

// i9SelectExpr 对应 SELECT 语句 SELECT 后面 FROM 前面的 select_expr
type i9SelectExpr interface {
	// f8BuildSelectExpr 构造 select_expr
	f8BuildSelectExpr(p7s6Builder *s6QueryBuilder) error
}
