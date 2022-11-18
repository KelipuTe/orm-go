package v20

// i9SelectExpr 对应查询语句里 select 子句的列或者聚合函数
type i9SelectExpr interface {
	f8BuildSelectExpr(p7s6Builder *s6QueryBuilder) error
}
