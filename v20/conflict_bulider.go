package v20

type S6ConflictBuilder[T any] struct {
	p7s6insert       *S6Insert[T]
	S5ConflictColumn []string
}

func (p7this *S6ConflictBuilder[T]) F8Column(s5column ...string) *S6ConflictBuilder[T] {
	p7this.S5ConflictColumn = s5column
	return p7this
}

func (p7this *S6ConflictBuilder[T]) F8Update(assign ...Assignable) *S6ConflictBuilder[T] {
	return p7this
}
