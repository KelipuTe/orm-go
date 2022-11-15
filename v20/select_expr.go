package v20

// I9SelectExpr 对应查询语句里 select 子句的列或者聚合函数
type I9SelectExpr interface {
	F8BuildSelectExpr(p7s6qb *s6QueryBuilder) error
}
