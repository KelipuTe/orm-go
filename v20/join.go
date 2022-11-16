package v20

type Join struct {
	i9LeftReference  I9TableReference
	operator         string
	i9RightReference I9TableReference
	s5On             []S6WhereCondition
}

func (this Join) F8GetTableAlies() string {
	return ""
}

func (this Join) F8BuildTableReference(p7s6Builder *s6QueryBuilder) error {
	var err error

	p7s6Builder.sqlString.WriteByte('(')

	if nil != this.i9LeftReference {
		err = this.i9LeftReference.F8BuildTableReference(p7s6Builder)
		if nil != err {
			return err
		}
	}
	p7s6Builder.sqlString.WriteString(" ")
	p7s6Builder.sqlString.WriteString(this.operator)
	p7s6Builder.sqlString.WriteString(" ")
	if nil != this.i9RightReference {
		err = this.i9RightReference.F8BuildTableReference(p7s6Builder)
		if nil != err {
			return err
		}
	}

	if 0 < len(this.s5On) {
		p7s6Builder.sqlString.WriteString(" ON ")
		err = p7s6Builder.F8BuildWhereCondition(this.s5On)
		if nil != err {
			return err
		}
	}

	p7s6Builder.sqlString.WriteByte(')')
	return nil
}

func (this Join) F8Join(i9reference I9TableReference) *JoinBuilder {
	return &JoinBuilder{
		i9LeftReference:  this,
		operator:         "JOIN",
		i9RightReference: i9reference,
	}
}

func (this Join) F8LeftJoin(i9reference I9TableReference) *JoinBuilder {
	return &JoinBuilder{
		i9LeftReference:  this,
		operator:         "LEFT JOIN",
		i9RightReference: i9reference,
	}
}

type JoinBuilder struct {
	i9LeftReference  I9TableReference
	operator         string
	i9RightReference I9TableReference
}

func (this JoinBuilder) F8On(s5condition ...S6WhereCondition) Join {
	return Join{
		i9LeftReference:  this.i9LeftReference,
		operator:         this.operator,
		i9RightReference: this.i9RightReference,
		s5On:             s5condition,
	}
}
