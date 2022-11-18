package v20

// #### Part Raw ####

// S6PartRaw 对应一段原生 SQL
type S6PartRaw struct {
	// SQLString 带有占位符的 SQL 语句
	SQLString string
	// S5Value SQL 语句中占位符对应的值
	S5Value []any
}

func (this S6PartRaw) F8PartRaw(p7s6qb *s6QueryBuilder) error {
	p7s6qb.sqlString.WriteString(this.SQLString)
	if 0 < len(this.S5Value) {
		p7s6qb.F8AddParameter(this.S5Value...)
	}
	return nil
}

func (this S6PartRaw) f8BuildSelectExpr(p7s6qb *s6QueryBuilder) error {
	return this.F8PartRaw(p7s6qb)
}

func (this S6PartRaw) F8BuildExpression(p7s6qb *s6QueryBuilder) error {
	return this.F8PartRaw(p7s6qb)
}

func (this S6PartRaw) ToPredicate() S6WhereCondition {
	return S6WhereCondition{
		LeftExpr:  this,
		Operator:  "",
		RightExpr: nil,
	}
}

func NewS6PartRaw(sql string, s5Value ...any) S6PartRaw {
	return S6PartRaw{
		SQLString: sql,
		S5Value:   s5Value,
	}
}
