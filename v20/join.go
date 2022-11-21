package v20

// S6Join 对应 JOIN，可以嵌套
type S6Join struct {
	// i9LeftReference JOIN 左边的部分
	i9LeftReference i9TableReference
	// operator JOIN、LEFT JOIN、RIGHT JOIN
	operator string
	// i9RightReference JOIN 右边的部分
	i9RightReference i9TableReference
	// s5On JOIN ... ON 后面的
	s5On []S6WhereCondition
}

func (this S6Join) f8GetTableReferenceAlies() string {
	return ""
}

func (this S6Join) f8BuildTableReference(p7s6Builder *s6QueryBuilder) error {
	var err error

	p7s6Builder.sqlString.WriteByte('(')

	// 递归处理左边的
	if nil != this.i9LeftReference {
		err = this.i9LeftReference.f8BuildTableReference(p7s6Builder)
		if nil != err {
			return err
		}
	}
	// 处理中间的操作符
	p7s6Builder.sqlString.WriteString(" ")
	p7s6Builder.sqlString.WriteString(this.operator)
	p7s6Builder.sqlString.WriteString(" ")

	// 递归处理右边的
	if nil != this.i9RightReference {
		err = this.i9RightReference.f8BuildTableReference(p7s6Builder)
		if nil != err {
			return err
		}
	}

	if 0 < len(this.s5On) {
		p7s6Builder.sqlString.WriteString(" ON ")
		err = p7s6Builder.f8BuildWhereCondition(this.s5On)
		if nil != err {
			return err
		}
	}

	p7s6Builder.sqlString.WriteByte(')')
	return nil
}

func (this S6Join) f8GetTableReferenceEntity() []any {
	return nil
}

func (this S6Join) F8Join(i9reference i9TableReference) *JoinBuilder {
	return &JoinBuilder{
		i9LeftReference:  this,
		operator:         "JOIN",
		i9RightReference: i9reference,
	}
}

func (this S6Join) F8LeftJoin(i9reference i9TableReference) *JoinBuilder {
	return &JoinBuilder{
		i9LeftReference:  this,
		operator:         "LEFT JOIN",
		i9RightReference: i9reference,
	}
}

type JoinBuilder struct {
	i9LeftReference  i9TableReference
	operator         string
	i9RightReference i9TableReference
}

func (this JoinBuilder) F8On(s5condition ...S6WhereCondition) S6Join {
	return S6Join{
		i9LeftReference:  this.i9LeftReference,
		operator:         this.operator,
		i9RightReference: this.i9RightReference,
		s5On:             s5condition,
	}
}
