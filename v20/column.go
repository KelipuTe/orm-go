package v20

// S6Column 对应查询条件里的列
type S6Column struct {
	// 列名
	name string
}

func (this S6Column) doExpression() {}

func (this S6Column) canSelect() {}

func (this S6Column) EQ(p any) Predicate {
	return Predicate{
		left:  this,
		op:    opEQ,
		right: toExpression(p),
	}
}

func (this S6Column) GT(p any) Predicate {
	return Predicate{
		left:  this,
		op:    opGT,
		right: toExpression(p),
	}
}

func (this S6Column) LT(p any) Predicate {
	return Predicate{
		left:  this,
		op:    opLT,
		right: toExpression(p),
	}
}

func NewField(n string) S6Column {
	return S6Column{
		name: n,
	}
}
