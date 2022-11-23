package v20

// S6WhereCondition 对应 where_condition
// SELECT Statement 里的 where_condition
// SELECT ... WHERE 后面的部分
// SELECT ... GROUP BY ... HAVING 后面的部分
// 可以通过嵌套组成复杂的查询条件
type S6WhereCondition s6BinaryExpression

func (this S6WhereCondition) f8BuildExpression(p7s6Builder *s6QueryBuilder) error {
	expr := s6BinaryExpression(this)
	return expr.f8BuildExpression(p7s6Builder)
}

//type S6WhereCondition struct {
//	// i9LeftExpr 操作符左边的查询条件
//	i9LeftExpr i9Expression
//	// s6Operator 操作符
//	s6Operator s6Operator
//	// i9RightExpr 操作符右边的查询条件
//	i9RightExpr i9Expression
//}

//func (this S6WhereCondition) f8BuildExpression(p7s6Builder *s6QueryBuilder) error {
//	var err error
//
//	if nil != this.i9LeftExpr {
//		_, lIsP := this.i9LeftExpr.(S6WhereCondition)
//		if lIsP {
//			p7s6Builder.sqlString.WriteByte('(')
//		}
//		err = this.i9LeftExpr.f8BuildExpression(p7s6Builder)
//		if nil != err {
//			return err
//		}
//		if lIsP {
//			p7s6Builder.sqlString.WriteByte(')')
//		}
//	}
//
//	// 处理中间的操作符
//	// 如果没有操作符，那么就是原生 sql，没有右边的部分
//	if "" == this.s6Operator.String() {
//		return nil
//	}
//	p7s6Builder.sqlString.WriteByte(' ')
//	p7s6Builder.sqlString.WriteString(this.s6Operator.String())
//	p7s6Builder.sqlString.WriteByte(' ')
//
//	// 递归处理右边的部分
//	if nil != this.i9RightExpr {
//		_, rIsP := this.i9RightExpr.(S6WhereCondition)
//		if rIsP {
//			p7s6Builder.sqlString.WriteByte('(')
//		}
//		err = this.i9RightExpr.f8BuildExpression(p7s6Builder)
//		if nil != err {
//			return err
//		}
//		if rIsP {
//			p7s6Builder.sqlString.WriteByte(')')
//		}
//	}
//
//	return nil
//}

// F8And 与，左查询条件 `与` 右查询条件 => (`Id` = 11) AND (S6Column = 'aa')
func (this S6WhereCondition) F8And(p S6WhereCondition) S6WhereCondition {
	return S6WhereCondition{
		i9LeftExpr:  this,
		s6Operator:  c5OperatorAND,
		i9RightExpr: p,
	}
}

// F8Or 或，左查询条件 `或` 右查询条件 => (`Id` = 11) OR (S6Column = 'aa')
func (this S6WhereCondition) F8Or(p S6WhereCondition) S6WhereCondition {
	return S6WhereCondition{
		i9LeftExpr:  this,
		s6Operator:  c5OperatorOR,
		i9RightExpr: p,
	}
}

// F8Not 非，`非` 右查询条件 => NOT (`id` = 11)
// 注意 F8Not 条件只有操作符右边的查询条件
func F8Not(p S6WhereCondition) S6WhereCondition {
	return S6WhereCondition{
		i9LeftExpr:  nil,
		s6Operator:  c5OperatorNOT,
		i9RightExpr: p,
	}
}
