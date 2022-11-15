package v20

// I9Assignment 标记接口，对应 INSERT 和 UPDATE 的赋值语句
// 即 INSERT Statement 和 UPDATE Statement 里的 assignment
type I9Assignment interface {
	F8Assignment()
}

type S6Assignment struct {
	Name string
	Expr I9Expression
}

func (this S6Assignment) F8Assignment() {}

func NewS6Assignment(name string, expr any) S6Assignment {
	t4Expr, ok := expr.(I9Expression)
	if !ok {
		t4Expr = S6Value{Value: expr}
	}
	return S6Assignment{
		Name: name,
		Expr: t4Expr,
	}
}
