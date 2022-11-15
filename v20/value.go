package v20

// S6Value 查询语句里占位符对应的值
type S6Value struct {
	Value any
}

func (this S6Value) F8BuildExpression(p7s6qb *s6QueryBuilder) error {
	p7s6qb.sqlString.WriteByte('?')
	p7s6qb.F8AddParameter(this.Value)
	return nil
}

// F8NewS6Value 把输入转换成查询语句里占位符对应的值
func F8NewS6Value(input any) S6Value {
	return S6Value{
		Value: input,
	}
}
