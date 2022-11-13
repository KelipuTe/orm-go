package clause

// S6Column 对应 SELECT 语句的查询表达式
// 即 SELECT Statement 里的 col_name
type S6Column struct {
	// 列名
	Name string
}

func (this S6Column) F8Expr() {}

func (this S6Column) F8SelectExpr() {}

func (this S6Column) F8Assignment() {}

func (this S6Column) EQ(p any) S6WhereCondition {
	return S6WhereCondition{
		LeftExpr:  this,
		Operator:  c5OperatorEqual,
		RightExpr: NewI9Expr(p),
	}
}

func (this S6Column) GT(p any) S6WhereCondition {
	return S6WhereCondition{
		LeftExpr:  this,
		Operator:  c5OperatorGreaterThan,
		RightExpr: NewI9Expr(p),
	}
}

func (this S6Column) LT(p any) S6WhereCondition {
	return S6WhereCondition{
		LeftExpr:  this,
		Operator:  c5OperatorLessThan,
		RightExpr: NewI9Expr(p),
	}
}

func NewS6Column(n string) S6Column {
	return S6Column{
		Name: n,
	}
}
