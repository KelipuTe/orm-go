package v20

// S6Aggregate 对应 aggregate 即聚合函数
type S6Aggregate struct {
	// name 聚合函数的函数名
	name string
	// s6Column 聚合函数操作的列
	s6Column S6Column
}

func (this S6Aggregate) f8BuildAggregate(p7s6Builder *s6QueryBuilder) error {
	p7s6Builder.sqlString.WriteString(this.name)
	p7s6Builder.sqlString.WriteByte('(')
	p7s6Builder.f8WrapWithQuote(this.s6Column.fieldName)
	p7s6Builder.sqlString.WriteByte(')')
	return nil
}

func (this S6Aggregate) f8BuildSelectExpr(p7s6Builder *s6QueryBuilder) error {
	return this.f8BuildAggregate(p7s6Builder)
}

func (this S6Aggregate) f8BuildExpression(p7s6Builder *s6QueryBuilder) error {
	return this.f8BuildAggregate(p7s6Builder)
}

// Equal 等于，this = input
func (this S6Aggregate) Equal(input any) S6WhereCondition {
	return S6WhereCondition{
		i9LeftExpr:  this,
		s6Operator:  c5OperatorEqual,
		i9RightExpr: f8NewI9Expression(input),
	}
}

// GreaterThan 大于，this > input
func (this S6Aggregate) GreaterThan(input any) S6WhereCondition {
	return S6WhereCondition{
		i9LeftExpr:  this,
		s6Operator:  c5OperatorGreaterThan,
		i9RightExpr: f8NewI9Expression(input),
	}
}

// LessThan 小于，this < input
func (this S6Aggregate) LessThan(input any) S6WhereCondition {
	return S6WhereCondition{
		i9LeftExpr:  this,
		s6Operator:  c5OperatorLessThan,
		i9RightExpr: f8NewI9Expression(input),
	}
}

// F8Count 对列求和
func F8Count(name string) S6Aggregate {
	return S6Aggregate{
		name:     "COUNT",
		s6Column: S6Column{fieldName: name},
	}
}

// F8Avg 对列求平均
func F8Avg(name string) S6Aggregate {
	return S6Aggregate{
		name:     "AVG",
		s6Column: S6Column{fieldName: name},
	}
}
