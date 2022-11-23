package v20

type s6BinaryExpression struct {
	// i9LeftExpr 操作符左边的查询条件
	i9LeftExpr i9Expression
	// s6Operator 操作符
	s6Operator s6Operator
	// i9RightExpr 操作符右边的查询条件
	i9RightExpr i9Expression
}

func (this s6BinaryExpression) f8BuildExpression(p7s6Builder *s6QueryBuilder) error {
	var err error

	if nil != this.i9LeftExpr {
		_, lIsP := this.i9LeftExpr.(S6WhereCondition)
		if lIsP {
			p7s6Builder.sqlString.WriteByte('(')
		}
		err = this.i9LeftExpr.f8BuildExpression(p7s6Builder)
		if nil != err {
			return err
		}
		if lIsP {
			p7s6Builder.sqlString.WriteByte(')')
		}
	}

	// 处理中间的操作符
	// 如果没有操作符，那么就是原生 sql，没有右边的部分
	if "" == this.s6Operator.String() {
		return nil
	}
	p7s6Builder.sqlString.WriteByte(' ')
	p7s6Builder.sqlString.WriteString(this.s6Operator.String())
	p7s6Builder.sqlString.WriteByte(' ')

	// 递归处理右边的部分
	if nil != this.i9RightExpr {
		_, rIsP := this.i9RightExpr.(S6WhereCondition)
		if rIsP {
			p7s6Builder.sqlString.WriteByte('(')
		}
		err = this.i9RightExpr.f8BuildExpression(p7s6Builder)
		if nil != err {
			return err
		}
		if rIsP {
			p7s6Builder.sqlString.WriteByte(')')
		}
	}

	return nil
}
