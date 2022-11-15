package v20

// S6Column 对应列
// 即 SELECT Statement 里的 col_name
type S6Column struct {
	// 列名
	Name string
}

func (this S6Column) F8BuildColumn(p7s6qb *s6QueryBuilder) error {
	p7s6qb.f8WrapWithQuote(this.Name)
	return nil
}

func (this S6Column) F8BuildSelectExpr(p7s6qb *s6QueryBuilder) error {
	return this.F8BuildColumn(p7s6qb)
}

func (this S6Column) F8BuildExpression(p7s6qb *s6QueryBuilder) error {
	return this.F8BuildColumn(p7s6qb)
}

func (this S6Column) F8BuildAssignment() error { return nil }

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
