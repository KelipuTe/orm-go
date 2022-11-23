package v20

// i9Assignment 标记接口，对应 INSERT 和 UPDATE 的赋值语句
// 即 INSERT Statement 和 UPDATE Statement 里的 assignment
type i9Assignment interface {
	f8BuildAssignment() error
}

type S6Assignment struct {
	s6Column S6Column
	i9Expr   i9Expression
}

func (this S6Assignment) f8BuildAssignment() error {
	return nil
}
