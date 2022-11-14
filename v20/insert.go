package v20

import (
	"orm-go/v20/clause"
	"orm-go/v20/metadata"
	"orm-go/v20/result"
)

type S6Insert[T any] struct {
	i9Session I9Session

	s6QueryBuilder
	// p7s6Model orm 映射模型
	p7s6Model *metadata.S6Model

	// 插入的数据
	s5Value []*T
	// s5ColumnName 插入的字段
	s5ColumnName []string
	p7s6Conflict *S6Conflict
}

type s6OnConflict[T any] struct {
	p7s6Insert *S6Insert[T]
}

// 跳到中间builder，处理OnConflict的内容
func (p7this *S6Insert[T]) f8OnConflictBuilder() *s6OnConflict[T] {
	return &s6OnConflict[T]{
		p7s6Insert: p7this,
	}
}

func F8NewS6Insert[T any](i9session I9Session) *S6Insert[T] {
	t4p7s6monitor := i9session.f8GetS6Monitor()
	return &S6Insert[T]{
		i9Session: i9session,
		s6QueryBuilder: s6QueryBuilder{
			s6Monitor: t4p7s6monitor,
			quote:     t4p7s6monitor.i9Dialect.f8GetQuoter(),
		},
	}
}

func (p7this *S6Insert[T]) F8SetValue(s5value ...*T) *S6Insert[T] {
	if 0 >= len(s5value) {
		return p7this
	}

	if nil == p7this.s5Value {
		p7this.s5Value = s5value
		return p7this
	}
	p7this.s5Value = append(p7this.s5Value, s5value...)
	return p7this
}

func (p7this *S6Insert[T]) F8SetColumnName(s5column ...string) *S6Insert[T] {
	if 0 >= len(s5column) {
		return p7this
	}

	if nil == p7this.s5ColumnName {
		p7this.s5ColumnName = s5column
		return p7this
	}
	p7this.s5ColumnName = append(p7this.s5ColumnName, s5column...)
	return p7this
}

func (p7this *S6Insert[T]) F8BuildQuery() (*clause.S6Query, error) {
	if nil == p7this.p7s6Model {
		t4p7s6model, err := p7this.s6Monitor.i9Registry.F8Get(p7this.s5Value[0])
		if nil != err {
			return nil, err
		}
		p7this.p7s6Model = t4p7s6model
	}

	p7this.sqlString.WriteString("INSERT INTO ")
	p7this.f8WrapWithQuote(p7this.p7s6Model.TableName)

	p7this.sqlString.WriteByte('(')
	s5p7s6ModelField := p7this.p7s6Model.S5P7S6ModelField
	if 0 != len(p7this.s5ColumnName) {
		s5p7s6ModelField = make([]*metadata.S6ModelField, 0, len(p7this.s5ColumnName))
		for _, t4ColumnName := range p7this.s5ColumnName {
			t4p7s6ModelField, ok := p7this.p7s6Model.M3StructToField[t4ColumnName]
			if !ok {
				return nil, result.F8NewErrUnknownColumn(t4ColumnName)
			}
			s5p7s6ModelField = append(s5p7s6ModelField, t4p7s6ModelField)
		}
	}

	// UPSERT 语句会传递额外的参数
	p7this.s5value = make([]any, 0, len(s5p7s6ModelField)*(len(p7this.s5Value)+1))
	for i, t4value := range s5p7s6ModelField {
		if 0 < i {
			p7this.sqlString.WriteByte(',')
		}
		p7this.f8WrapWithQuote(t4value.FieldName)
	}

	p7this.sqlString.WriteString(") VALUES")

	for i, t4value := range p7this.s5Value {
		if 0 < i {
			p7this.sqlString.WriteByte(',')
		}
		t4i9result := p7this.f8NewI9Result(t4value, p7this.p7s6Model)
		p7this.sqlString.WriteByte('(')
		for j, t4value2 := range s5p7s6ModelField {
			if 0 < j {
				p7this.sqlString.WriteByte(',')
			}
			p7this.sqlString.WriteByte('?')
			fdVal, err := t4i9result.F8GetField(t4value2.StructName)
			if err != nil {
				return nil, err
			}
			p7this.F8AddParameter(fdVal)
		}

		p7this.sqlString.WriteByte(')')
	}

	if nil != p7this.p7s6Conflict {
		err := p7this.s6Monitor.i9Dialect.f8BuildOnConflict(&p7this.s6QueryBuilder, p7this.p7s6Conflict)
		if err != nil {
			return nil, err
		}
	}

	p7this.sqlString.WriteByte(';')
	return &clause.S6Query{
		SQLString: p7this.sqlString.String(),
		S5Value:   p7this.s5value,
	}, nil
}
