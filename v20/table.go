package v20

type S6Table struct {
	// 传到构造器的结构体
	p7struct any
	// 别名
	alies string
}

func (this S6Table) F8As(name string) S6Table {
	return S6Table{
		p7struct: this.p7struct,
		alies:    name,
	}
}

func (this S6Table) F8Column(name string) S6Column {
	return S6Column{
		i9From: this,
		Name:   name,
	}
}

func (this S6Table) F8BuildTableReference(p7s6Builder *s6QueryBuilder) error {
	p7s6Model, err := p7s6Builder.s6Monitor.i9Registry.F8Get(this.p7struct)
	if nil != err {
		return err
	}
	p7s6Builder.f8WrapWithQuote(p7s6Model.TableName)
	if "" != this.alies {
		p7s6Builder.sqlString.WriteString(" AS ")
		p7s6Builder.f8WrapWithQuote(this.alies)
	}
	return nil
}

func (this S6Table) F8GetTableAlies() string {
	return this.alies
}

func (this S6Table) F8Join(i9reference I9TableReference) *JoinBuilder {
	return &JoinBuilder{
		i9LeftReference:  this,
		operator:         "JOIN",
		i9RightReference: i9reference,
	}
}

func (this S6Table) F8LeftJoin(i9reference I9TableReference) *JoinBuilder {
	return &JoinBuilder{
		i9LeftReference:  this,
		operator:         "LEFT JOIN",
		i9RightReference: i9reference,
	}
}

func F8NewS6Table(input any) S6Table {
	return S6Table{
		p7struct: input,
	}
}
