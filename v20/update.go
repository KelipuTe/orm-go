package v20

import "orm-go/v20/internal"

type S6Update[T any] struct {
	p7Entity       *T
	s5i9Assignment []i9Assignment
	s5where        []S6WhereCondition

	i9Session I9Session
	s6QueryBuilder
}

func (p7this *S6Update[T]) F8SetEntity(p7Entity *T) *S6Update[T] {
	if nil == p7Entity {
		return p7this
	}
	p7this.p7Entity = p7Entity
	return p7this
}

func (p7this *S6Update[T]) F8SetUpdate(s5i9Assignment ...i9Assignment) *S6Update[T] {
	if 0 >= len(s5i9Assignment) {
		return p7this
	}
	if nil == p7this.s5i9Assignment {
		p7this.s5i9Assignment = s5i9Assignment
		return p7this
	}
	p7this.s5i9Assignment = append(p7this.s5i9Assignment, s5i9Assignment...)
	return p7this
}

func (p7this *S6Update[T]) F8Where(s5condition ...S6WhereCondition) *S6Update[T] {
	if 0 >= len(s5condition) {
		return p7this
	}
	if nil == p7this.s5where {
		p7this.s5where = s5condition
		return p7this
	}
	p7this.s5where = append(p7this.s5where, s5condition...)
	return p7this
}

func (p7this *S6Update[T]) F8BuildQuery() (*S6Query, error) {
	var err error

	if nil == p7this.p7Entity {
		p7this.p7Entity = new(T)
	}
	if 0 >= len(p7this.s5i9Assignment) {
		return nil, internal.ErrUpdateWithoutColumn
	}
	if 0 >= len(p7this.s5where) {
		return nil, internal.ErrUpdateWithoutWhere
	}

	p7this.s6QueryBuilder.p7s6Model, err = p7this.s6Monitor.i9Registry.F8Get(p7this.p7Entity)
	if nil != err {
		return nil, err
	}

	p7this.s6QueryBuilder.sqlString.WriteString("UPDATE ")
	p7this.s6QueryBuilder.f8WrapWithQuote(p7this.s6QueryBuilder.p7s6Model.TableName)
	p7this.s6QueryBuilder.sqlString.WriteString(" SET ")

	// 通过反射拿到结构体属性的值
	t4i9result := p7this.f8NewI9Result(p7this.p7Entity, p7this.p7s6Model)
	for i, t4value := range p7this.s5i9Assignment {
		if 0 < i {
			p7this.s6QueryBuilder.sqlString.WriteByte(',')
		}
		switch t4value2 := t4value.(type) {
		case S6Column:
			p7s6ModelField, ok := p7this.s6QueryBuilder.p7s6Model.M3FieldToColumn[t4value2.fieldName]
			if !ok {
				return nil, internal.F8NewErrUnknownField(t4value2.fieldName)
			}
			p7this.s6QueryBuilder.f8WrapWithQuote(p7s6ModelField.ColumnName)
			p7this.s6QueryBuilder.sqlString.WriteString("=?")
			t4EntityValue, err2 := t4i9result.F8GetField(p7s6ModelField.FieldName)
			if err2 != nil {
				return nil, err2
			}
			p7this.f8AddParameter(t4EntityValue)
		case S6Assignment:
			p7s6ModelField, ok := p7this.s6QueryBuilder.p7s6Model.M3FieldToColumn[t4value2.s6Column.fieldName]
			if !ok {
				return nil, internal.F8NewErrUnknownField(t4value2.s6Column.fieldName)
			}
			p7this.s6QueryBuilder.f8WrapWithQuote(p7s6ModelField.ColumnName)
			p7this.s6QueryBuilder.sqlString.WriteByte('=')
			err = p7this.f8BuildExpression(t4value2.i9Expr)
			if nil != err {
				return nil, err
			}
		}
	}

	p7this.sqlString.WriteString(" WHERE ")
	err = p7this.f8BuildWhereCondition(p7this.s5where)
	if nil != err {
		return nil, err
	}

	p7this.s6QueryBuilder.sqlString.WriteByte(';')

	return &S6Query{
		SQLString: p7this.s6QueryBuilder.sqlString.String(),
		S5Value:   p7this.s6QueryBuilder.s5Value,
	}, nil
}

func F8NewS6Update[T any](i9Session I9Session) *S6Update[T] {
	t4p7s6monitor := i9Session.f8GetS6Monitor()
	return &S6Update[T]{
		i9Session: i9Session,
		s6QueryBuilder: s6QueryBuilder{
			s6Monitor: t4p7s6monitor,
			quote:     t4p7s6monitor.i9Dialect.f8GetQuoter(),
		},
	}
}
