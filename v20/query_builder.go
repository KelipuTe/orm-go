package v20

import (
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

// F8BuildWhereCondition 处理查询条件
func (p7this *s6QueryBuilder) F8BuildWhereCondition(s5p []S6WhereCondition) error {
	t4p := s5p[0]
	for i := 1; i < len(s5p); i++ {
		t4p = t4p.And(s5p[i])
	}
	//t4p.F8BuildExpression(p7this)
	return p7this.F8BuildExpression(t4p)
}

func (p7this *s6QueryBuilder) F8BuildExpression(expr I9Expression) error {
	//var err error

	if nil == expr {
		return nil
	}

	return expr.F8BuildExpression(p7this)

	//switch t4type := expr.(type) {
	//case S6WhereCondition:
	//	// 处理语句
	//	// 递归处理左边的部分
	//	_, lIsP := t4type.LeftExpr.(S6WhereCondition)
	//	if lIsP {
	//		p7this.sqlString.WriteByte('(')
	//	}
	//	err = p7this.F8BuildExpression(t4type.LeftExpr)
	//	if nil != err {
	//		return err
	//	}
	//	if lIsP {
	//		p7this.sqlString.WriteByte(')')
	//	}
	//
	//	// 处理中间的操作符
	//	// 如果没有操作符，那么就是原生 sql，没有右边的部分
	//	if "" == t4type.Operator.String() {
	//		return nil
	//	}
	//	p7this.sqlString.WriteByte(' ')
	//	p7this.sqlString.WriteString(t4type.Operator.String())
	//	p7this.sqlString.WriteByte(' ')
	//	// 递归处理右边的部分
	//	_, rIsP := t4type.RightExpr.(S6WhereCondition)
	//	if rIsP {
	//		p7this.sqlString.WriteByte('(')
	//	}
	//	err = p7this.F8BuildExpression(t4type.RightExpr)
	//	if nil != err {
	//		return err
	//	}
	//	if rIsP {
	//		p7this.sqlString.WriteByte(')')
	//	}
	//case S6Column:
	//	// 处理列名
	//	err = p7this.F8BuildColumn(t4type)
	//	if nil != err {
	//		return err
	//	}
	//case S6Aggregate:
	//	// 处理聚合函数
	//	err = p7this.F8BuildAggregate(t4type)
	//	if nil != err {
	//		return err
	//	}
	//case S6PartRaw:
	//	// 处理原生 sql
	//	p7this.sqlString.WriteString(t4type.SQLString)
	//	if 0 < len(t4type.S5Value) {
	//		p7this.F8AddParameter(t4type.S5Value...)
	//	}
	//case S6Value:
	//	// 处理占位符对应的参数
	//	p7this.sqlString.WriteByte('?')
	//	p7this.F8AddParameter(t4type.Value)
	//default:
	//	return NewErrUnsupportedExpressionType(expr)
	//}
	//return nil
}

// F8AddParameter 添加占位符对应的参数
func (p7this *s6QueryBuilder) F8AddParameter(s5p ...any) {
	if nil == p7this.s5value {
		p7this.s5value = make([]any, 0, 8)
	}
	p7this.s5value = append(p7this.s5value, s5p...)
}
