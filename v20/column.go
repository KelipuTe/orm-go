package v20

// S6Column 对应 col_name
// 即语句中表示 表、JOIN、子查询 的列的部分
type S6Column struct {
	// 列对应的：表、JOIN、子查询
	i9From i9TableReference
	// name 列名
	name string
	// alias 别名
	alias string
}

func (this S6Column) f8BuildColumn(p7s6Builder *s6QueryBuilder) error {
	if nil != this.i9From {
		alies := this.i9From.f8GetTableReferenceAlies()
		if "" != alies {
			p7s6Builder.f8WrapWithQuote(alies)
			p7s6Builder.sqlString.WriteByte('.')
		}
	}
	p7s6Builder.f8WrapWithQuote(this.name)
	return nil
}

func (this S6Column) f8BuildSelectExpr(p7s6qb *s6QueryBuilder) error {
	return this.f8BuildColumn(p7s6qb)
}

func (this S6Column) f8BuildExpression(p7s6Builder *s6QueryBuilder) error {
	return this.f8BuildColumn(p7s6Builder)
}

func (this S6Column) F8BuildAssignment() error { return nil }

func (this S6Column) F8Equal(p any) S6WhereCondition {
	return S6WhereCondition{
		i9LeftExpr:  this,
		s6Operator:  c5OperatorEqual,
		i9RightExpr: f8NewI9Expression(p),
	}
}

func (this S6Column) F8GreaterThan(p any) S6WhereCondition {
	return S6WhereCondition{
		i9LeftExpr:  this,
		s6Operator:  c5OperatorGreaterThan,
		i9RightExpr: f8NewI9Expression(p),
	}
}

func (this S6Column) F8LessThan(p any) S6WhereCondition {
	return S6WhereCondition{
		i9LeftExpr:  this,
		s6Operator:  c5OperatorLessThan,
		i9RightExpr: f8NewI9Expression(p),
	}
}

func (this S6Column) F8InQuery(sub S6SubQuery) S6WhereCondition {
	return S6WhereCondition{
		i9LeftExpr:  this,
		s6Operator:  c5OperatorIn,
		i9RightExpr: sub,
	}
}

func F8NewS6Column(name string) S6Column {
	return S6Column{
		name:  name,
		alias: "",
	}
}
