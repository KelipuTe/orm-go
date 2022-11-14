package v20

import "orm-go/v20/clause"

type S6ConflictBuilder[T any] struct {
	p7s6Insert       *S6Insert[T]
	S5ConflictColumn []string
}

func (p7this *S6ConflictBuilder[T]) F8Column(s5column ...string) *S6ConflictBuilder[T] {
	p7this.S5ConflictColumn = s5column
	return p7this
}

func (p7this *S6ConflictBuilder[T]) F8Update(s5i9assignment ...clause.I9Assignment) *S6Insert[T] {
	p7this.p7s6Insert.p7s6Conflict = &S6Conflict{
		S5ConflictColumn: p7this.S5ConflictColumn,
		S5I9Assignment:   s5i9assignment,
	}
	return p7this.p7s6Insert
}