package v20

// S6Table 表，对应 table_references
type S6Table struct {
	// p7struct 用构造器的那个泛型 new(T) 出来的，用于获取表名
	p7struct any
	// alias 别名
	alias string
}

// F8As 给表设置别名
func (this S6Table) F8As(alias string) S6Table {
	return S6Table{
		p7struct: this.p7struct,
		alias:    alias,
	}
}

// F8Column 创建带表的列 `table_name`.`column_name`
func (this S6Table) F8Column(name string) S6Column {
	return S6Column{
		i9From:    this,
		fieldName: name,
		alias:     "",
	}
}

func (this S6Table) f8BuildTableReference(p7s6Builder *s6QueryBuilder) error {
	p7s6Model, err := p7s6Builder.s6Monitor.i9Registry.F8Get(this.p7struct)
	if nil != err {
		return err
	}
	p7s6Builder.f8WrapWithQuote(p7s6Model.TableName)
	if "" != this.alias {
		p7s6Builder.sqlString.WriteString(" AS ")
		p7s6Builder.f8WrapWithQuote(this.alias)
	}
	return nil
}

func (this S6Table) f8GetTableReferenceAlies() string {
	return this.alias
}

func (this S6Table) f8GetTableReferenceEntity() []any {
	return nil
}

func (this S6Table) F8Join(i9reference i9TableReference) *JoinBuilder {
	return &JoinBuilder{
		i9LeftReference:  this,
		operator:         "JOIN",
		i9RightReference: i9reference,
	}
}

func (this S6Table) F8LeftJoin(i9reference i9TableReference) *JoinBuilder {
	return &JoinBuilder{
		i9LeftReference:  this,
		operator:         "LEFT JOIN",
		i9RightReference: i9reference,
	}
}

func F8NewS6Table(input any) S6Table {
	return S6Table{
		p7struct: input,
		alias:    "",
	}
}
