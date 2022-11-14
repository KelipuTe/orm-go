package clause

// S6Aggregate 对应聚合函数
type S6Aggregate struct {
	// S6Column 聚合函数的函数名
	Name string
	// S6Column 聚合函数操作的列
	S6Column S6Column
}

func (this S6Aggregate) F8SelectExpr() {}

func (this S6Aggregate) F8Expression() {}

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
