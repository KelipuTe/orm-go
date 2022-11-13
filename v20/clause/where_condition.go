package clause

// S6WhereCondition 对应 where 子句和 having 子句的查询条件
// 即 SELECT Statement 里的 where_condition
// S6WhereCondition 可以嵌套，组成复杂的查询条件
type S6WhereCondition struct {
	// LeftExpr 操作符左边的查询条件
	LeftExpr I9Expr
	// Operator 操作符
	Operator s6Operator
	// RightExpr 操作符右边的查询条件
	RightExpr I9Expr
}

func (this S6WhereCondition) F8Expr() {}

// And 与，左查询条件 `与` 右查询条件 => (`Id` = 11) AND (S6Column = 'aa')
func (this S6WhereCondition) And(p S6WhereCondition) S6WhereCondition {
	return S6WhereCondition{
		LeftExpr:  this,
		Operator:  c5OperatorAND,
		RightExpr: p,
	}
}

// Or 或，左查询条件 `或` 右查询条件 => (`Id` = 11) OR (S6Column = 'aa')
func (this S6WhereCondition) Or(p S6WhereCondition) S6WhereCondition {
	return S6WhereCondition{
		LeftExpr:  this,
		Operator:  c5OperatorOR,
		RightExpr: p,
	}
}

// Not 非，`非` 右查询条件 => NOT (`id` = 11)
// 注意 Not 条件只有操作符右边的查询条件
func Not(p S6WhereCondition) S6WhereCondition {
	return S6WhereCondition{
		LeftExpr:  nil,
		Operator:  c5OperatorNOT,
		RightExpr: p,
	}
}
