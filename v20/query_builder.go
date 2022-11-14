package v20

import (
	"orm-go/v20/clause"
	"strings"
)

type s6QueryBuilder struct {
	s6Monitor
	// quote 这个东西 s6Monitor 里面有，拿到这里方便操作
	quote byte
	// sqlString 带有占位符的 SQL 语句
	sqlString strings.Builder
	// s5value SQL 语句中占位符对应的参数
	s5value []any
}

func (p7this *s6QueryBuilder) f8WrapWithQuote(name string) {
	p7this.sqlString.WriteByte(p7this.quote)
	p7this.sqlString.WriteString(name)
	p7this.sqlString.WriteByte(p7this.quote)
}

func (p7this *s6QueryBuilder) F8BuildExpression(e clause.I9Expression) error {
	var err error

	if nil == e {
		return nil
	}

	switch e.(type) {
	case clause.S6WhereCondition:
		// 处理语句
		t4predicate := e.(clause.S6WhereCondition)
		// 递归处理左边的部分
		_, lIsP := t4predicate.LeftExpr.(clause.S6WhereCondition)
		if lIsP {
			p7this.sqlString.WriteByte('(')
		}
		err = p7this.F8BuildExpression(t4predicate.LeftExpr)
		if nil != err {
			return err
		}
		if lIsP {
			p7this.sqlString.WriteByte(')')
		}

		// 处理中间的操作符
		// 如果没有操作符，那么就是原生 sql，没有右边的部分
		if "" == t4predicate.Operator.String() {
			return nil
		}
		p7this.sqlString.WriteByte(' ')
		p7this.sqlString.WriteString(t4predicate.Operator.String())
		p7this.sqlString.WriteByte(' ')
		// 递归处理右边的部分
		_, rIsP := t4predicate.RightExpr.(clause.S6WhereCondition)
		if rIsP {
			p7this.sqlString.WriteByte('(')
		}
		err = p7this.F8BuildExpression(t4predicate.RightExpr)
		if nil != err {
			return err
		}
		if rIsP {
			p7this.sqlString.WriteByte(')')
		}
	case clause.S6Column:
		// 处理列名
		t4c := e.(clause.S6Column)
		err = p7this.F8BuildColumn(t4c)
		if nil != err {
			return err
		}
	case clause.S6Aggregate:
		// 处理聚合函数
		t4a := e.(clause.S6Aggregate)
		err = p7this.F8BuildAggregate(t4a)
		if nil != err {
			return err
		}
	case clause.S6PartRaw:
		// 处理原生 sql
		t4r := e.(clause.S6PartRaw)
		p7this.sqlString.WriteString(t4r.SQLString)
		if 0 < len(t4r.S5Value) {
			p7this.F8AddParameter(t4r.S5Value...)
		}
	case clause.S6Value:
		// 处理占位符对应的参数
		t4parameter := e.(clause.S6Value)
		p7this.sqlString.WriteByte('?')
		p7this.F8AddParameter(t4parameter.Value)
	default:
		return NewErrUnsupportedExpressionType(e)
	}
	return nil
}

// F8BuildColumn 处理列
func (p7this *s6QueryBuilder) F8BuildColumn(column clause.S6Column) error {
	p7this.sqlString.WriteByte('`')
	p7this.sqlString.WriteString(column.Name)
	p7this.sqlString.WriteByte('`')
	return nil
}

// F8BuildWhereCondition 处理查询条件
func (p7this *s6QueryBuilder) F8BuildWhereCondition(s5p []clause.S6WhereCondition) error {
	t4p := s5p[0]
	for i := 1; i < len(s5p); i++ {
		t4p = t4p.And(s5p[i])
	}
	return p7this.F8BuildExpression(t4p)
}

// F8BuildAggregate 处理聚合函数
func (p7this *s6QueryBuilder) F8BuildAggregate(s6aggregate clause.S6Aggregate) error {
	p7this.sqlString.WriteString(s6aggregate.Name)
	p7this.sqlString.WriteString("(`")
	p7this.sqlString.WriteString(s6aggregate.S6Column.Name)
	p7this.sqlString.WriteString("`)")
	return nil
}

// F8BuildExpression 处理语句
func (p7this *S6Select[T]) F8BuildExpression(e clause.I9Expression) error {
	var err error

	if nil == e {
		return nil
	}

	switch e.(type) {
	case clause.S6WhereCondition:
		// 处理语句
		t4predicate := e.(clause.S6WhereCondition)
		// 递归处理左边的部分
		_, lIsP := t4predicate.LeftExpr.(clause.S6WhereCondition)
		if lIsP {
			p7this.sqlString.WriteByte('(')
		}
		err = p7this.F8BuildExpression(t4predicate.LeftExpr)
		if nil != err {
			return err
		}
		if lIsP {
			p7this.sqlString.WriteByte(')')
		}

		// 处理中间的操作符
		// 如果没有操作符，那么就是原生 sql，没有右边的部分
		if "" == t4predicate.Operator.String() {
			return nil
		}
		p7this.sqlString.WriteByte(' ')
		p7this.sqlString.WriteString(t4predicate.Operator.String())
		p7this.sqlString.WriteByte(' ')
		// 递归处理右边的部分
		_, rIsP := t4predicate.RightExpr.(clause.S6WhereCondition)
		if rIsP {
			p7this.sqlString.WriteByte('(')
		}
		err = p7this.F8BuildExpression(t4predicate.RightExpr)
		if nil != err {
			return err
		}
		if rIsP {
			p7this.sqlString.WriteByte(')')
		}
	case clause.S6Column:
		// 处理列名
		t4c := e.(clause.S6Column)
		err = p7this.F8BuildColumn(t4c)
		if nil != err {
			return err
		}
	case clause.S6Aggregate:
		// 处理聚合函数
		t4a := e.(clause.S6Aggregate)
		err = p7this.F8BuildAggregate(t4a)
		if nil != err {
			return err
		}
	case clause.S6PartRaw:
		// 处理原生 sql
		t4r := e.(clause.S6PartRaw)
		p7this.sqlString.WriteString(t4r.SQLString)
		if 0 < len(t4r.S5Value) {
			p7this.F8AddParameter(t4r.S5Value...)
		}
	case clause.S6Value:
		// 处理占位符对应的参数
		t4parameter := e.(clause.S6Value)
		p7this.sqlString.WriteByte('?')
		p7this.F8AddParameter(t4parameter.Value)
	default:
		return NewErrUnsupportedExpressionType(e)
	}
	return nil
}

// F8AddParameter 添加占位符对应的参数
func (p7this *s6QueryBuilder) F8AddParameter(s5p ...any) {
	if nil == p7this.s5value {
		p7this.s5value = make([]any, 0, 2)
	}
	p7this.s5value = append(p7this.s5value, s5p...)
}
