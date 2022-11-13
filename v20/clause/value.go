package clause

// S6Value 查询语句里占位符对应的值
type S6Value struct {
	Value any
}

func (this S6Value) F8Expr() {}

// NewS6Value 把输入转换成查询语句里占位符对应的值
func NewS6Value(input any) S6Value {
	return S6Value{
		Value: input,
	}
}
