package v20

// Assignable 标记接口，
// 实现该接口意味着可以用于赋值语句，
// 用于在 UPDATE 和 UPSERT 中
type I9Assignment interface {
	isAssignment()
}

type S6Assignment struct {
	column string
	val    Expression
}

func (a S6Assignment) isAssignment() {}

func NewS6Assignment(column string, val any) S6Assignment {
	v, ok := val.(Expression)
	if !ok {
		v = s6value{value: val}
	}
	return S6Assignment{
		column: column,
		val:    v,
	}
}
