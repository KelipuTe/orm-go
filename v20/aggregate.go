package v20

// S6Aggregate 对应聚合函数
type S6Aggregate struct {
	// Name 聚合函数的函数名
	Name string
	// S6Column 聚合函数操作的列
	S6Column S6Column
}

func (this S6Aggregate) f8BuildAggregate(p7s6Builder *s6QueryBuilder) error {
	p7s6Builder.sqlString.WriteString(this.Name)
	p7s6Builder.sqlString.WriteByte('(')
	p7s6Builder.f8WrapWithQuote(this.S6Column.Name)
	p7s6Builder.sqlString.WriteByte(')')
	return nil
}

func (this S6Aggregate) f8BuildSelectExpr(p7s6Builder *s6QueryBuilder) error {
	return this.f8BuildAggregate(p7s6Builder)
}

func (this S6Aggregate) F8BuildExpression(p7s6Builder *s6QueryBuilder) error {
	return this.f8BuildAggregate(p7s6Builder)
}

func (this S6Aggregate) Equal(input any) S6WhereCondition {
	return S6WhereCondition{
		LeftExpr:  this,
		Operator:  c5OperatorEqual,
		RightExpr: F8NewI9Expression(input),
	}
}

func (this S6Aggregate) GreaterThan(input any) S6WhereCondition {
	return S6WhereCondition{
		LeftExpr:  this,
		Operator:  c5OperatorGreaterThan,
		RightExpr: F8NewI9Expression(input),
	}
}

func (this S6Aggregate) LessThan(input any) S6WhereCondition {
	return S6WhereCondition{
		LeftExpr:  this,
		Operator:  c5OperatorLessThan,
		RightExpr: F8NewI9Expression(input),
	}
}

// Count 对列求和
func Count(name string) S6Aggregate {
	return S6Aggregate{
		Name:     "COUNT",
		S6Column: S6Column{Name: name},
	}
}

// Avg 对列求平均
func Avg(name string) S6Aggregate {
	return S6Aggregate{
		Name:     "AVG",
		S6Column: S6Column{Name: name},
	}
}
