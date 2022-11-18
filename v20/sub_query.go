package v20

type S6SubQuery struct {
	s5Select []i9SelectExpr
	i9From   i9TableReference
	alias    string

	i9Builder I9QueryBuilder
}

func (this S6SubQuery) F8BuildSubQuery(p7s6Builder *s6QueryBuilder, useAlias bool) error {
	query, err := this.i9Builder.F8BuildQuery()
	if err != nil {
		return err
	}
	p7s6Builder.sqlString.WriteByte('(')
	p7s6Builder.sqlString.WriteString(query.SQLString[:len(query.SQLString)-1])
	if 0 < len(query.S5Value) {
		p7s6Builder.F8AddParameter(query.S5Value...)
	}
	p7s6Builder.sqlString.WriteByte(')')
	if useAlias && "" != this.alias {
		p7s6Builder.sqlString.WriteString(" AS ")
		p7s6Builder.f8WrapWithQuote(this.alias)
	}
	return nil
}

func (this S6SubQuery) f8BuildExpression(p7s6Builder *s6QueryBuilder) error {
	return this.F8BuildSubQuery(p7s6Builder, false)
}

func (this S6SubQuery) f8GetTableReferenceAlies() string {
	return this.alias
}

func (this S6SubQuery) f8BuildTableReference(p7s6Builder *s6QueryBuilder) error {
	return this.F8BuildSubQuery(p7s6Builder, true)
}

func (this S6SubQuery) F8Join(i9reference i9TableReference) *JoinBuilder {
	return &JoinBuilder{
		i9LeftReference:  this,
		operator:         "JOIN",
		i9RightReference: i9reference,
	}
}

func (this S6SubQuery) F8LeftJoin(i9reference i9TableReference) *JoinBuilder {
	return &JoinBuilder{
		i9LeftReference:  this,
		operator:         "LEFT JOIN",
		i9RightReference: i9reference,
	}
}

func (this S6SubQuery) F8Column(name string) S6Column {
	return S6Column{
		i9From: this,
		name:   name,
	}
}
