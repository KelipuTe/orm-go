package v20

// S6WhereCondition 对应 where 子句和 having 子句的查询条件
// 即 SELECT Statement 里的 where_condition
// S6WhereCondition 可以嵌套，组成复杂的查询条件
type S6WhereCondition struct {
	// LeftExpr 操作符左边的查询条件
	LeftExpr I9Expression
	// Operator 操作符
	Operator s6Operator
	// RightExpr 操作符右边的查询条件
	RightExpr I9Expression
}

func (this S6WhereCondition) F8BuildExpression(p7s6qb *s6QueryBuilder) error {
	var err error

	if nil != this.LeftExpr {
		_, lIsP := this.LeftExpr.(S6WhereCondition)
		if lIsP {
			p7s6qb.sqlString.WriteByte('(')
		}
		err = this.LeftExpr.F8BuildExpression(p7s6qb)
		if nil != err {
			return err
		}
		if lIsP {
			p7s6qb.sqlString.WriteByte(')')
		}
	}

	// 处理中间的操作符
	// 如果没有操作符，那么就是原生 sql，没有右边的部分
	if "" == this.Operator.String() {
		return nil
	}
	p7s6qb.sqlString.WriteByte(' ')
	p7s6qb.sqlString.WriteString(this.Operator.String())
	p7s6qb.sqlString.WriteByte(' ')

	// 递归处理右边的部分
	if nil != this.RightExpr {
		_, rIsP := this.RightExpr.(S6WhereCondition)
		if rIsP {
			p7s6qb.sqlString.WriteByte('(')
		}
		err = this.RightExpr.F8BuildExpression(p7s6qb)
		if nil != err {
			return err
		}
		if rIsP {
			p7s6qb.sqlString.WriteByte(')')
		}
	}

	return nil
}

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
