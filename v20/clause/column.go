package clause

// S6Column 对应列
// 即 SELECT Statement 里的 col_name
type S6Column struct {
	// 列名
	Name string
}

func (this S6Column) F8Expression() {}

func (this S6Column) F8SelectExpr() {}

func (this S6Column) F8Assignment() {}

func (this S6Column) F8EQ(p any) S6WhereCondition {
	return S6WhereCondition{
		LeftExpr:  this,
		Operator:  c5OperatorEqual,
		RightExpr: F8NewI9Expression(p),
	}
}

func (this S6Column) F8GT(p any) S6WhereCondition {
	return S6WhereCondition{
		LeftExpr:  this,
		Operator:  c5OperatorGreaterThan,
		RightExpr: F8NewI9Expression(p),
	}
}

func (this S6Column) F8LT(p any) S6WhereCondition {
	return S6WhereCondition{
		LeftExpr:  this,
		Operator:  c5OperatorLessThan,
		RightExpr: F8NewI9Expression(p),
	}
}

func F8NewS6Column(n string) S6Column {
	return S6Column{
		Name: n,
	}
}
