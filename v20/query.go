package v20

import (
	"orm-go/v20/metadata"
	"strings"
)

// S6Query I9QueryBuilder.F8BuildQuery 的结果
type S6Query struct {
	// SQLString 带有占位符的 SQL 语句
	SQLString string
	// S5Value SQL 语句中占位符对应的值
	S5Value []any
}

// I9QueryBuilder 接口抽象：查询构造器
// Builder 设计模式
type I9QueryBuilder interface {
	// F8BuildQuery 方法抽象：构造 S6Query
	F8BuildQuery() (*S6Query, error)
}

type s6QueryBuilder struct {
	s6Monitor
	p7s6Model *metadata.S6Model
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
		t4p = t4p.F8And(s5p[i])
	}
	//t4p.f8BuildExpression(p7this)
	return p7this.F8BuildExpression(t4p)
}

func (p7this *s6QueryBuilder) F8BuildExpression(expr i9Expression) error {
	//var err error

	if nil == expr {
		return nil
	}

	return expr.f8BuildExpression(p7this)

	//switch t4type := expr.(type) {
	//case S6WhereCondition:
	//	// 处理语句
	//	// 递归处理左边的部分
	//	_, lIsP := t4type.i9LeftExpr.(S6WhereCondition)
	//	if lIsP {
	//		p7this.sqlString.WriteByte('(')
	//	}
	//	err = p7this.f8BuildExpression(t4type.i9LeftExpr)
	//	if nil != err {
	//		return err
	//	}
	//	if lIsP {
	//		p7this.sqlString.WriteByte(')')
	//	}
	//
	//	// 处理中间的操作符
	//	// 如果没有操作符，那么就是原生 sql，没有右边的部分
	//	if "" == t4type.operator.String() {
	//		return nil
	//	}
	//	p7this.sqlString.WriteByte(' ')
	//	p7this.sqlString.WriteString(t4type.operator.String())
	//	p7this.sqlString.WriteByte(' ')
	//	// 递归处理右边的部分
	//	_, rIsP := t4type.i9RightExpr.(S6WhereCondition)
	//	if rIsP {
	//		p7this.sqlString.WriteByte('(')
	//	}
	//	err = p7this.f8BuildExpression(t4type.i9RightExpr)
	//	if nil != err {
	//		return err
	//	}
	//	if rIsP {
	//		p7this.sqlString.WriteByte(')')
	//	}
	//case s6Column:
	//	// 处理列名
	//	err = p7this.f8BuildColumn(t4type)
	//	if nil != err {
	//		return err
	//	}
	//case S6Aggregate:
	//	// 处理聚合函数
	//	err = p7this.f8BuildAggregate(t4type)
	//	if nil != err {
	//		return err
	//	}
	//case S6PartRaw:
	//	// 处理原生 sql
	//	p7this.sqlString.WriteString(t4type.sqlString)
	//	if 0 < len(t4type.s5Value) {
	//		p7this.F8AddParameter(t4type.s5Value...)
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
