package v20

// I9Assignment 标记接口，对应 INSERT 和 UPDATE 的赋值语句
// 即 INSERT Statement 和 UPDATE Statement 里的 assignment
type I9Assignment interface {
	F8BuildAssignment() error
}

type S6Assignment struct {
	Name string
	Expr I9Expression
}

func (this S6Assignment) F8BuildAssignment() error {
	return nil
}

func NewS6Assignment(name string, input any) S6Assignment {
	expr, ok := input.(I9Expression)
	if !ok {
		expr = S6Value{Value: input}
	}
	return S6Assignment{
		Name: name,
		Expr: expr,
	}
}
