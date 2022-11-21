package v20

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestS6InsertF8Build(p7s6t *testing.T) {
	p7s6DB := F8NewS6DB(nil)

	s5case := []struct {
		name         string
		queryBuilder I9QueryBuilder
		wantQuery    *S6Query
		wantErr      error
	}{
		{
			name: "insert_one",
			queryBuilder: F8NewS6Insert[S6APPUserModel](p7s6DB).
				F8SetValue(&S6APPUserModel{Id: 11, Name: "aa", Age: 22, Sex: 1}),
			wantQuery: &S6Query{
				SQLString: "INSERT INTO `app_user`(`id`,`fieldName`,`age`,`sex`) VALUES(?,?,?,?);",
				S5Value:   []any{11, "aa", int8(22), int8(1)},
			},
			wantErr: nil,
		},
		{
			name: "insert_two",
			queryBuilder: F8NewS6Insert[S6APPUserModel](p7s6DB).
				F8SetValue(
					&S6APPUserModel{Id: 11, Name: "aa", Age: 22, Sex: 1},
					&S6APPUserModel{Id: 22, Name: "bb", Age: 33, Sex: 2},
				),
			wantQuery: &S6Query{
				SQLString: "INSERT INTO `app_user`(`id`,`fieldName`,`age`,`sex`) VALUES(?,?,?,?),(?,?,?,?);",
				S5Value:   []any{11, "aa", int8(22), int8(1), 22, "bb", int8(33), int8(2)},
			},
			wantErr: nil,
		},
		{
			name: "insert_one_set_column",
			queryBuilder: F8NewS6Insert[S6APPUserModel](p7s6DB).
				F8SetValue(&S6APPUserModel{Id: 11, Name: "aa", Age: 22, Sex: 1}).
				F8SetColumn("Id", "Name"),
			wantQuery: &S6Query{
				SQLString: "INSERT INTO `app_user`(`id`,`fieldName`) VALUES(?,?);",
				S5Value:   []any{11, "aa"},
			},
			wantErr: nil,
		},
		{
			name: "insert_one_on_conflict_update_with_column",
			queryBuilder: F8NewS6Insert[S6APPUserModel](p7s6DB).
				F8SetValue(&S6APPUserModel{Id: 11, Name: "aa", Age: 22, Sex: 1}).
				f8OnConflictBuilder().
				F8Update(S6Column{fieldName: "id"}, S6Column{fieldName: "fieldName"}),
			wantQuery: &S6Query{
				SQLString: "INSERT INTO `app_user`(`id`,`fieldName`,`age`,`sex`) VALUES(?,?,?,?) " +
					"ON DUPLICATE KEY UPDATE `id`=VALUES(`id`),`fieldName`=VALUES(`fieldName`);",
				S5Value: []any{11, "aa", int8(22), int8(1)},
			},
			wantErr: nil,
		},
		{
			name: "insert_one_on_conflict_update_with_value",
			queryBuilder: F8NewS6Insert[S6APPUserModel](p7s6DB).
				F8SetValue(&S6APPUserModel{Id: 11, Name: "aa", Age: 22, Sex: 1}).
				f8OnConflictBuilder().
				F8Update(NewS6Assignment("fieldName", "bb")),
			wantQuery: &S6Query{
				SQLString: "INSERT INTO `app_user`(`id`,`fieldName`,`age`,`sex`) VALUES(?,?,?,?) " +
					"ON DUPLICATE KEY UPDATE `fieldName`=?;",
				S5Value: []any{11, "aa", int8(22), int8(1), "bb"},
			},
			wantErr: nil,
		},
	}

	for _, t4value := range s5case {
		p7s6t.Run(t4value.name, func(p7s6t2 *testing.T) {
			query, err := t4value.queryBuilder.F8BuildQuery()
			assert.Equal(p7s6t2, t4value.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(p7s6t2, t4value.wantQuery, query)
		})
	}
}

func TestS6InsertF8BuildSQLite3(p7s6t *testing.T) {
	p7s6DB := F8NewS6DB(nil, F8DBWithDialect(S6SQLite3Dialect))

	s5case := []struct {
		name         string
		queryBuilder I9QueryBuilder
		wantQuery    *S6Query
		wantErr      error
	}{
		{
			name: "insert_one_on_conflict_update_with_column",
			queryBuilder: F8NewS6Insert[S6APPUserModel](p7s6DB).
				F8SetValue(&S6APPUserModel{Id: 11, Name: "aa", Age: 22, Sex: 1}).
				f8OnConflictBuilder().
				F8ConflictColumn(F8NewS6Column("id")).
				F8Update(S6Column{fieldName: "fieldName"}, S6Column{fieldName: "age"}),
			wantQuery: &S6Query{
				SQLString: "INSERT INTO `app_user`(`id`,`fieldName`,`age`,`sex`) VALUES(?,?,?,?) " +
					"ON CONFLICT (`id`) DO UPDATE SET `fieldName`=excluded.`fieldName`,`age`=excluded.`age`;",
				S5Value: []any{11, "aa", int8(22), int8(1)},
			},
			wantErr: nil,
		},
		{
			name: "insert_one_on_conflict_update_with_value",
			queryBuilder: F8NewS6Insert[S6APPUserModel](p7s6DB).
				F8SetValue(&S6APPUserModel{Id: 11, Name: "aa", Age: 22, Sex: 1}).
				f8OnConflictBuilder().
				F8ConflictColumn(F8NewS6Column("id")).
				F8Update(NewS6Assignment("fieldName", "bb")),
			wantQuery: &S6Query{
				SQLString: "INSERT INTO `app_user`(`id`,`fieldName`,`age`,`sex`) VALUES(?,?,?,?) " +
					"ON CONFLICT (`id`) DO UPDATE SET `fieldName`=?;",
				S5Value: []any{11, "aa", int8(22), int8(1), "bb"},
			},
			wantErr: nil,
		},
	}

	for _, t4value := range s5case {
		p7s6t.Run(t4value.name, func(p7s6t2 *testing.T) {
			query, err := t4value.queryBuilder.F8BuildQuery()
			assert.Equal(p7s6t2, t4value.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(p7s6t2, t4value.wantQuery, query)
		})
	}
}
