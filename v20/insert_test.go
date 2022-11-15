package v20

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestS6OrmInsertF8Build(p7t *testing.T) {
	p7s6DB := F8NewS6DB(nil)

	s5case := []struct {
		name      string
		i9qb      I9QueryBuilder
		wantQuery *S6Query
		wantErr   error
	}{
		{
			name: "insert",
			i9qb: F8NewS6Insert[S6TestModel](p7s6DB).
				F8SetValue(&S6TestModel{Id: 11, Name: "aa", Age: 22, Sex: 1}),
			wantQuery: &S6Query{
				SQLString: "INSERT INTO `test_model`(`id`,`name`,`age`,`sex`) VALUES(?,?,?,?);",
				S5Value:   []any{11, "aa", int8(22), int8(1)},
			},
			wantErr: nil,
		},
	}

	for _, t4case := range s5case {
		p7t.Run(t4case.name, func(p7t *testing.T) {
			query, err := t4case.i9qb.F8BuildQuery()
			assert.Equal(p7t, t4case.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(p7t, t4case.wantQuery, query)
		})
	}
}
