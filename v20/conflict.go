package v20

type S6Conflict struct {
	S5ConflictColumn []S6Column
	S5Assignment     []i9Assignment
}

type S6ConflictBuilder[T any] struct {
	p7s6Insert       *S6Insert[T]
	S5ConflictColumn []S6Column
}

func (p7this *S6ConflictBuilder[T]) F8ConflictColumn(s5column ...S6Column) *S6ConflictBuilder[T] {
	p7this.S5ConflictColumn = s5column
	return p7this
}

func (p7this *S6ConflictBuilder[T]) F8SetUpdate(s5i9assignment ...i9Assignment) *S6Insert[T] {
	p7this.p7s6Insert.p7s6Conflict = &S6Conflict{
		S5ConflictColumn: p7this.S5ConflictColumn,
		S5Assignment:     s5i9assignment,
	}
	return p7this.p7s6Insert
}
