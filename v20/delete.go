package v20

import "orm-go/v20/internal"

type S6Delete[T any] struct {
	p7Entity *T
	s5where  []S6WhereCondition

	i9Session I9Session
	s6QueryBuilder
}

func (p7this *S6Delete[T]) F8SetEntity(p7Entity *T) *S6Delete[T] {
	if nil == p7Entity {
		return p7this
	}
	p7this.p7Entity = p7Entity
	return p7this
}

func (p7this *S6Delete[T]) F8Where(s5condition ...S6WhereCondition) *S6Delete[T] {
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

func (p7this *S6Delete[T]) F8BuildQuery() (*S6Query, error) {
	var err error

	if nil == p7this.p7Entity {
		p7this.p7Entity = new(T)
	}
	if 0 >= len(p7this.s5where) {
		return nil, internal.ErrDeleteWithoutWhere
	}

	p7this.s6QueryBuilder.p7s6Model, err = p7this.s6Monitor.i9Registry.F8Get(p7this.p7Entity)
	if nil != err {
		return nil, err
	}

	p7this.s6QueryBuilder.sqlString.WriteString("DELETE FROM ")
	p7this.s6QueryBuilder.f8WrapWithQuote(p7this.s6QueryBuilder.p7s6Model.TableName)

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

func F8NewS6Delete[T any](i9Session I9Session) *S6Delete[T] {
	t4p7s6monitor := i9Session.f8GetS6Monitor()
	return &S6Delete[T]{
		i9Session: i9Session,
		s6QueryBuilder: s6QueryBuilder{
			s6Monitor: t4p7s6monitor,
			quote:     t4p7s6monitor.i9Dialect.f8GetQuoter(),
		},
	}
}
