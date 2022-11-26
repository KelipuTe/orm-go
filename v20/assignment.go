package v20

// i9Assignment 标记接口，对应 INSERT 和 UPDATE 的赋值语句
// 即 INSERT Statement 和 UPDATE Statement 里的 assignment
type i9Assignment interface {
	// 构造赋值语句的 SQL
	f8BuildAssignment() error
}

// S6Assignment 赋值语句
type S6Assignment struct {
	// s6Column 列
	s6Column S6Column
	// i9Expr 表达式
	i9Expr i9Expression
}

// f8BuildAssignment 赋值语句，对应，列 = 表达式，这种
func (this S6Assignment) f8BuildAssignment() error { return nil }
