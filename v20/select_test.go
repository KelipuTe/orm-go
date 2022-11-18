package v20

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelectBuildQuery(p7s6t *testing.T) {
	p7s6DB := F8NewS6DB(nil)

	s5case := []struct {
		name         string
		queryBuilder I9QueryBuilder
		wantQuery    *S6Query
		wantErr      error
	}{
		{
			name:         "all",
			queryBuilder: F8NewS6Select[S6APPUserModel](p7s6DB),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user`;",
			},
		},
	}

	for _, t4value := range s5case {
		p7s6t.Run(t4value.name, func(p7s6t2 *testing.T) {
			p7query, err := t4value.queryBuilder.F8BuildQuery()
			assert.Equal(p7s6t2, t4value.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(p7s6t2, t4value.wantQuery, p7query)
		})
	}
}

func TestSelectOperator(p7s6t *testing.T) {
	p7s6DB := F8NewS6DB(nil)

	s5case := []struct {
		name         string
		queryBuilder I9QueryBuilder
		wantQuery    *S6Query
		wantErr      error
	}{
		{
			name: "where_equal",
			queryBuilder: F8NewS6Select[S6APPUserModel](p7s6DB).
				F8Where(F8NewS6Column("Id").F8Equal(11)),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` WHERE `Id` = ?;",
				S5Value:   []any{11},
			},
		},
		{
			name: "where_greater_than",
			queryBuilder: F8NewS6Select[S6APPUserModel](p7s6DB).
				F8Where(F8NewS6Column("Id").F8GreaterThan(11)),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` WHERE `Id` > ?;",
				S5Value:   []any{11},
			},
		},
		{
			name: "where_less_than",
			queryBuilder: F8NewS6Select[S6APPUserModel](p7s6DB).
				F8Where(F8NewS6Column("Id").F8LessThan(11)),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` WHERE `Id` < ?;",
				S5Value:   []any{11},
			},
		},
	}

	for _, t4value := range s5case {
		p7s6t.Run(t4value.name, func(p7s6t2 *testing.T) {
			p7query, err := t4value.queryBuilder.F8BuildQuery()
			assert.Equal(p7s6t2, t4value.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(p7s6t2, t4value.wantQuery, p7query)
		})
	}
}

func TestSelectWhere(p7s6t *testing.T) {
	p7s6DB := F8NewS6DB(nil)

	s5case := []struct {
		name         string
		queryBuilder I9QueryBuilder
		wantQuery    *S6Query
		wantErr      error
	}{
		{
			name: "where_no",
			queryBuilder: F8NewS6Select[S6APPUserModel](p7s6DB).
				F8Where(),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user`;",
				S5Value:   nil,
			},
		},
		{
			name: "where_one",
			queryBuilder: F8NewS6Select[S6APPUserModel](p7s6DB).
				F8Where(F8NewS6Column("Id").F8Equal(11)),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` WHERE `Id` = ?;",
				S5Value:   []any{11},
			},
		},
		{
			name: "where_two",
			queryBuilder: F8NewS6Select[S6APPUserModel](p7s6DB).
				F8Where(F8NewS6Column("Id").F8Equal(11)).
				F8Where(F8NewS6Column("S6Column").F8Equal("aa")),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` WHERE (`Id` = ?) AND (`S6Column` = ?);",
				S5Value:   []any{11, "aa"},
			},
		},
		{
			name: "where_one_and_one",
			queryBuilder: F8NewS6Select[S6APPUserModel](p7s6DB).
				F8Where(F8NewS6Column("Id").F8Equal(11).And(F8NewS6Column("S6Column").F8Equal("aa"))),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` WHERE (`Id` = ?) AND (`S6Column` = ?);",
				S5Value:   []any{11, "aa"},
			},
		},
		{
			name: "where_one_or_one",
			queryBuilder: F8NewS6Select[S6APPUserModel](p7s6DB).
				F8Where(F8NewS6Column("Id").F8Equal(11).Or(F8NewS6Column("S6Column").F8Equal("aa"))),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` WHERE (`Id` = ?) OR (`S6Column` = ?);",
				S5Value:   []any{11, "aa"},
			},
		},
		{
			name: "where_not_one",
			queryBuilder: F8NewS6Select[S6APPUserModel](p7s6DB).
				F8Where(Not(F8NewS6Column("Id").F8Equal(11))),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` WHERE  NOT (`Id` = ?);",
				S5Value:   []any{11},
			},
		},
		{
			name: "where_one_and_(one_and_one)",
			queryBuilder: F8NewS6Select[S6APPUserModel](p7s6DB).
				F8Where(F8NewS6Column("Id").F8Equal(11).And(F8NewS6Column("S6Column").F8Equal("aa").And(F8NewS6Column("Age").F8Equal(22)))),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` WHERE (`Id` = ?) AND ((`S6Column` = ?) AND (`Age` = ?));",
				S5Value:   []any{11, "aa", 22},
			},
		},
		{
			name: "where_one_or_(one_or_one)",
			queryBuilder: F8NewS6Select[S6APPUserModel](p7s6DB).
				F8Where(F8NewS6Column("Id").F8Equal(11).And(F8NewS6Column("S6Column").F8Equal("aa").Or(F8NewS6Column("Age").F8Equal(22)))),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` WHERE (`Id` = ?) AND ((`S6Column` = ?) OR (`Age` = ?));",
				S5Value:   []any{11, "aa", 22},
			},
		},
		{
			name: "where_one_and_(not_one)",
			queryBuilder: F8NewS6Select[S6APPUserModel](p7s6DB).
				F8Where(F8NewS6Column("Id").F8Equal(11).And(Not(F8NewS6Column("S6Column").F8Equal("aa")))),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` WHERE (`Id` = ?) AND ( NOT (`S6Column` = ?));",
				S5Value:   []any{11, "aa"},
			},
		},
	}

	for _, t4value := range s5case {
		p7s6t.Run(t4value.name, func(p7s6t2 *testing.T) {
			p7query, err := t4value.queryBuilder.F8BuildQuery()
			assert.Equal(p7s6t2, t4value.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(p7s6t2, t4value.wantQuery, p7query)
		})
	}
}

func TestSelectGroupBy(p7s6t *testing.T) {
	p7s6DB := F8NewS6DB(nil)

	s5case := []struct {
		name         string
		queryBuilder I9QueryBuilder
		wantQuery    *S6Query
		wantErr      error
	}{
		{
			name: "group_by_no",
			queryBuilder: F8NewS6Select[S6APPUserModel](p7s6DB).
				F8GroupBy(),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user`;",
				S5Value:   nil,
			},
		},
		{
			name: "group_by_one",
			queryBuilder: F8NewS6Select[S6APPUserModel](p7s6DB).
				F8GroupBy(F8NewS6Column("Age")),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` GROUP BY `Age`;",
				S5Value:   nil,
			},
		},
		{
			name: "group_by_two",
			queryBuilder: F8NewS6Select[S6APPUserModel](p7s6DB).
				F8GroupBy(F8NewS6Column("Age")).F8GroupBy(F8NewS6Column("Sex")),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` GROUP BY `Age`,`Sex`;",
				S5Value:   nil,
			},
		},
	}

	for _, t4value := range s5case {
		p7s6t.Run(t4value.name, func(p7s6t2 *testing.T) {
			p7query, err := t4value.queryBuilder.F8BuildQuery()
			assert.Equal(p7s6t2, t4value.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(p7s6t2, t4value.wantQuery, p7query)
		})
	}
}

func TestSelectHaving(p7s6t *testing.T) {
	p7s6DB := F8NewS6DB(nil)

	s5case := []struct {
		name         string
		queryBuilder I9QueryBuilder
		wantQuery    *S6Query
		wantErr      error
	}{
		{
			name: "having_no",
			queryBuilder: F8NewS6Select[S6APPUserModel](p7s6DB).
				F8Having(),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user`;",
				S5Value:   nil,
			},
		},
		{
			name: "having_no_group_by",
			queryBuilder: F8NewS6Select[S6APPUserModel](p7s6DB).
				F8Having(F8NewS6Column("Age").F8GreaterThan(22)),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user`;",
				S5Value:   nil,
			},
		},
		{
			name: "having_one",
			queryBuilder: F8NewS6Select[S6APPUserModel](p7s6DB).
				F8GroupBy(F8NewS6Column("Age")).
				F8Having(F8NewS6Column("Id").F8Equal(11)),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` GROUP BY `Age` HAVING `Id` = ?;",
				S5Value:   []any{11},
			},
		},
		{
			name: "group_by_two_having_one",
			queryBuilder: F8NewS6Select[S6APPUserModel](p7s6DB).
				F8GroupBy(F8NewS6Column("Age")).F8GroupBy(F8NewS6Column("Sex")).
				F8Having(F8NewS6Column("Id").F8Equal(11)),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` GROUP BY `Age`,`Sex` HAVING `Id` = ?;",
				S5Value:   []any{11},
			},
		},
		{
			name: "group_by_two_having_two",
			queryBuilder: F8NewS6Select[S6APPUserModel](p7s6DB).
				F8GroupBy(F8NewS6Column("Age")).F8GroupBy(F8NewS6Column("Sex")).
				F8Having(F8NewS6Column("Id").F8Equal(11)).F8Having(F8NewS6Column("S6Column").F8Equal("aa")),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` GROUP BY `Age`,`Sex` HAVING (`Id` = ?) AND (`S6Column` = ?);",
				S5Value:   []any{11, "aa"},
			},
		},
	}

	for _, t4value := range s5case {
		p7s6t.Run(t4value.name, func(p7s6t2 *testing.T) {
			p7query, err := t4value.queryBuilder.F8BuildQuery()
			assert.Equal(p7s6t2, t4value.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(p7s6t2, t4value.wantQuery, p7query)
		})
	}
}

func TestSelectOrderBy(p7s6t *testing.T) {
	p7s6DB := F8NewS6DB(nil)

	s5case := []struct {
		name         string
		queryBuilder I9QueryBuilder
		wantQuery    *S6Query
		wantErr      error
	}{
		{
			name:         "order_by_no",
			queryBuilder: F8NewS6Select[S6APPUserModel](p7s6DB).F8OrderBy(),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user`;",
				S5Value:   nil,
			},
		},
		{
			name: "order_by_one_asc",
			queryBuilder: F8NewS6Select[S6APPUserModel](p7s6DB).
				F8OrderBy(F8Asc("S6Column")),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` ORDER BY `S6Column` ASC;",
				S5Value:   nil,
			},
		},
		{
			name: "order_by_one_desc",
			queryBuilder: F8NewS6Select[S6APPUserModel](p7s6DB).
				F8OrderBy(F8Desc("S6Column")),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` ORDER BY `S6Column` DESC;",
				S5Value:   nil,
			},
		},
		{
			name: "order_by_two_asc_desc",
			queryBuilder: F8NewS6Select[S6APPUserModel](p7s6DB).
				F8OrderBy(F8Asc("S6Column")).F8OrderBy(F8Desc("Age")),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` ORDER BY `S6Column` ASC,`Age` DESC;",
				S5Value:   nil,
			},
		},
	}

	for _, t4value := range s5case {
		p7s6t.Run(t4value.name, func(p7s6t *testing.T) {
			p7query, err := t4value.queryBuilder.F8BuildQuery()
			assert.Equal(p7s6t, t4value.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(p7s6t, t4value.wantQuery, p7query)
		})
	}
}

func TestSelectOffsetLimit(p7s6t *testing.T) {
	p7s6DB := F8NewS6DB(nil)

	s5case := []struct {
		name         string
		queryBuilder I9QueryBuilder
		wantQuery    *S6Query
		wantErr      error
	}{
		{
			name: "limit",
			queryBuilder: F8NewS6Select[S6APPUserModel](p7s6DB).
				F8Limit(11),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` LIMIT ?;",
				S5Value:   []any{11},
			},
		},
		{
			name: "offset",
			queryBuilder: F8NewS6Select[S6APPUserModel](p7s6DB).
				F8Offset(111),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` OFFSET ?;",
				S5Value:   []any{111},
			},
		},
		{
			name: "limit_offset",
			queryBuilder: F8NewS6Select[S6APPUserModel](p7s6DB).
				F8Limit(11).F8Offset(111),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` LIMIT ? OFFSET ?;",
				S5Value:   []any{11, 111},
			},
		},
	}

	for _, t4value := range s5case {
		p7s6t.Run(t4value.name, func(p7s6t *testing.T) {
			p7query, err := t4value.queryBuilder.F8BuildQuery()
			assert.Equal(p7s6t, t4value.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(p7s6t, t4value.wantQuery, p7query)
		})
	}
}

func TestSelectSelect(p7s6t *testing.T) {
	p7s6DB := F8NewS6DB(nil)

	s5case := []struct {
		name         string
		queryBuilder I9QueryBuilder
		wantQuery    *S6Query
		wantErr      error
	}{
		{
			name: "select_one_column",
			queryBuilder: F8NewS6Select[S6APPUserModel](p7s6DB).
				F8Select(F8NewS6Column("Id")),
			wantQuery: &S6Query{
				SQLString: "SELECT `Id` FROM `app_user`;",
				S5Value:   nil,
			},
		},
		{
			name: "select_two_column",
			queryBuilder: F8NewS6Select[S6APPUserModel](p7s6DB).
				F8Select(F8NewS6Column("Id")).F8Select(F8NewS6Column("S6Column")),
			wantQuery: &S6Query{
				SQLString: "SELECT `Id`,`S6Column` FROM `app_user`;",
				S5Value:   nil,
			},
		},
	}

	for _, t4value := range s5case {
		p7s6t.Run(t4value.name, func(p7s6t2 *testing.T) {
			p7query, err := t4value.queryBuilder.F8BuildQuery()
			assert.Equal(p7s6t2, t4value.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(p7s6t2, t4value.wantQuery, p7query)
		})
	}
}

func TestSelectAggregate(p7s6t *testing.T) {
	p7s6DB := F8NewS6DB(nil)

	s5case := []struct {
		name         string
		queryBuilder I9QueryBuilder
		wantQuery    *S6Query
		wantErr      error
	}{
		{
			name: "select_one_aggregate",
			queryBuilder: F8NewS6Select[S6APPUserModel](p7s6DB).
				F8Select(Count("Id")),
			wantQuery: &S6Query{
				SQLString: "SELECT COUNT(`Id`) FROM `app_user`;",
				S5Value:   nil,
			},
		},
		{
			name: "select_two_aggregate",
			queryBuilder: F8NewS6Select[S6APPUserModel](p7s6DB).
				F8Select(Count("Id")).F8Select(Avg("Age")),
			wantQuery: &S6Query{
				SQLString: "SELECT COUNT(`Id`),AVG(`Age`) FROM `app_user`;",
				S5Value:   nil,
			},
		},
		{
			name: "having_one_aggregate",
			queryBuilder: F8NewS6Select[S6APPUserModel](p7s6DB).
				F8GroupBy(F8NewS6Column("Age")).
				F8Having(Count("Id").GreaterThan(5)),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` GROUP BY `Age` HAVING COUNT(`Id`) > ?;",
				S5Value:   []any{5},
			},
		},
	}

	for _, t4value := range s5case {
		p7s6t.Run(t4value.name, func(p7s6t2 *testing.T) {
			p7query, err := t4value.queryBuilder.F8BuildQuery()
			assert.Equal(p7s6t2, t4value.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(p7s6t2, t4value.wantQuery, p7query)
		})
	}
}

func TestSelectJoin(p7s6t *testing.T) {
	p7s6DB := F8NewS6DB(nil)

	s5case := []struct {
		name         string
		queryBuilder I9QueryBuilder
		wantQuery    *S6Query
		wantErr      error
	}{
		{
			name: "a_join_b",
			queryBuilder: func() I9QueryBuilder {
				t1 := F8NewS6Table(&S6APPUserModel{}).F8As("t1")
				t2 := F8NewS6Table(&S6APPUserInfoModel{})
				return F8NewS6Select[S6APPUserModel](p7s6DB).
					F8From(t1.F8Join(t2).F8On(t1.F8Column("id").F8Equal(t2.F8Column("user_id"))))
			}(),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM (`app_user` AS `t1` JOIN `app_user_info` ON `t1`.`id` = `user_id`);",
				S5Value:   nil,
			},
		},
		{
			name: "(a_join_b)_join_c",
			queryBuilder: func() I9QueryBuilder {
				t1 := F8NewS6Table(&S6APPUserModel{}).F8As("t1")
				t2 := F8NewS6Table(&S6APPUserInfoModel{}).F8As("t2")
				t3 := F8NewS6Table(&S6APPUserOrder{}).F8As("t3")
				return F8NewS6Select[S6APPUserModel](p7s6DB).
					F8From(t1.F8Join(t2).F8On(t1.F8Column("id").F8Equal(t2.F8Column("user_id"))).
						F8Join(t3).F8On(t1.F8Column("id").F8Equal(t3.F8Column("user_id"))))
			}(),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM ((`app_user` AS `t1` JOIN `app_user_info` AS `t2` ON `t1`.`id` = `t2`.`user_id`) " +
					"JOIN `app_user_order` AS `t3` ON `t1`.`id` = `t3`.`user_id`);",
				S5Value: nil,
			},
		},
		{
			name: "a_leftjoin_b",
			queryBuilder: func() I9QueryBuilder {
				t1 := F8NewS6Table(&S6APPUserModel{}).F8As("t1")
				t2 := F8NewS6Table(&S6APPUserInfoModel{})
				return F8NewS6Select[S6APPUserModel](p7s6DB).
					F8From(t1.F8LeftJoin(t2).F8On(t1.F8Column("id").F8Equal(t2.F8Column("user_id"))))
			}(),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM (`app_user` AS `t1` LEFT JOIN `app_user_info` ON `t1`.`id` = `user_id`);",
				S5Value:   nil,
			},
		},
		{
			name: "(a_leftjoin_b)_join_c",
			queryBuilder: func() I9QueryBuilder {
				t1 := F8NewS6Table(&S6APPUserModel{}).F8As("t1")
				t2 := F8NewS6Table(&S6APPUserInfoModel{}).F8As("t2")
				t3 := F8NewS6Table(&S6APPUserOrder{}).F8As("t3")
				return F8NewS6Select[S6APPUserModel](p7s6DB).
					F8From(t1.F8LeftJoin(t2).F8On(t1.F8Column("id").F8Equal(t2.F8Column("user_id"))).
						F8Join(t3).F8On(t1.F8Column("id").F8Equal(t3.F8Column("user_id"))))
			}(),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM ((`app_user` AS `t1` LEFT JOIN `app_user_info` AS `t2` ON `t1`.`id` = `t2`.`user_id`) " +
					"JOIN `app_user_order` AS `t3` ON `t1`.`id` = `t3`.`user_id`);",
				S5Value: nil,
			},
		},
	}

	for _, t4value := range s5case {
		p7s6t.Run(t4value.name, func(p7s6t2 *testing.T) {
			p7query, err := t4value.queryBuilder.F8BuildQuery()
			assert.Equal(p7s6t2, t4value.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(p7s6t2, t4value.wantQuery, p7query)
		})
	}
}

func TestSelectSubQuery(p7s6t *testing.T) {
	p7s6DB := F8NewS6DB(nil)

	s5case := []struct {
		name         string
		queryBuilder I9QueryBuilder
		wantQuery    *S6Query
		wantErr      error
	}{
		{
			name: "a_from_b",
			queryBuilder: func() I9QueryBuilder {
				sub := F8NewS6Select[S6APPUserInfoModel](p7s6DB).
					F8AsSubQuery("sub")
				return F8NewS6Select[S6APPUserModel](p7s6DB).
					F8From(sub)
			}(),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM (SELECT * FROM `app_user_info`) AS `sub`;",
				S5Value:   nil,
			},
		},
		{
			name: "a_in_b",
			queryBuilder: func() I9QueryBuilder {
				sub := F8NewS6Select[S6APPUserInfoModel](p7s6DB).
					F8AsSubQuery("sub")
				return F8NewS6Select[S6APPUserModel](p7s6DB).
					F8Where(F8NewS6Column("id").F8InQuery(sub))
			}(),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` WHERE `id` IN (SELECT * FROM `app_user_info`);",
				S5Value:   nil,
			},
		},
	}

	for _, t4value := range s5case {
		p7s6t.Run(t4value.name, func(p7s6t2 *testing.T) {
			p7query, err := t4value.queryBuilder.F8BuildQuery()
			assert.Equal(p7s6t2, t4value.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(p7s6t2, t4value.wantQuery, p7query)
		})
	}
}

func TestSelectPartRaw(p7s6t *testing.T) {
	p7s6DB := F8NewS6DB(nil)

	s5case := []struct {
		name         string
		queryBuilder I9QueryBuilder
		wantQuery    *S6Query
		wantErr      error
	}{
		{
			name: "select_part_raw",
			queryBuilder: F8NewS6Select[S6APPUserModel](p7s6DB).
				F8Select(NewS6PartRaw("DISTINCT(Id)")),
			wantQuery: &S6Query{
				SQLString: "SELECT DISTINCT(Id) FROM `app_user`;",
				S5Value:   nil,
			},
		},
		{
			name: "where_part_raw",
			queryBuilder: F8NewS6Select[S6APPUserModel](p7s6DB).
				F8Where(NewS6PartRaw("Id > ?", 11).ToPredicate()),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` WHERE Id > ?;",
				S5Value:   []any{11},
			},
		},
		{
			name: "where_part_raw_and_one",
			queryBuilder: F8NewS6Select[S6APPUserModel](p7s6DB).
				F8Where(NewS6PartRaw("Id > ?", 11).ToPredicate().And(F8NewS6Column("S6Column").F8Equal("aa"))),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` WHERE (Id > ?) AND (`S6Column` = ?);",
				S5Value:   []any{11, "aa"},
			},
		},
		{
			name: "having_part_raw",
			queryBuilder: F8NewS6Select[S6APPUserModel](p7s6DB).
				F8GroupBy(F8NewS6Column("Age")).
				F8Having(NewS6PartRaw("COUNT(Id) > ?", 5).ToPredicate()),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `app_user` GROUP BY `Age` HAVING COUNT(Id) > ?;",
				S5Value:   []any{5},
			},
		},
	}

	for _, t4value := range s5case {
		p7s6t.Run(t4value.name, func(p7s6t2 *testing.T) {
			p7query, err := t4value.queryBuilder.F8BuildQuery()
			assert.Equal(p7s6t2, t4value.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(p7s6t2, t4value.wantQuery, p7query)
		})
	}
}

func TestSelectorF8Get(p7s6t *testing.T) {
	// 构造 mock 数据库连接
	p7s6MockDB, sqlMock, err := sqlmock.New()
	if nil != err {
		p7s6t.Fatal(err)
	}
	defer func() {
		_ = p7s6MockDB.Close()
	}()
	p7s6DB := F8NewS6DB(p7s6MockDB)

	s5case := []struct {
		name      string
		sqlString string
		mockRows  *sqlmock.Rows
		mockErr   error
		wantRes   *S6APPUserModel
		wantErr   error
	}{
		{
			name:      "normal_sql",
			sqlString: "SELECT .*",
			mockRows: func() *sqlmock.Rows {
				rows := sqlmock.NewRows([]string{"id", "name", "age", "sex"})
				rows.AddRow([]byte("11"), []byte("aa"), []byte("22"), []byte("1"))
				return rows
			}(),
			wantRes: &S6APPUserModel{Id: 11, Name: "aa", Age: 22, Sex: 1},
		},
	}

	// 把预设的查询结果装进 mock
	for _, t4value := range s5case {
		t4p7eq := sqlMock.ExpectQuery(t4value.sqlString)
		if nil != t4value.mockErr {
			t4p7eq.WillReturnError(t4value.mockErr)
		} else {
			t4p7eq.WillReturnRows(t4value.mockRows)
		}
	}

	for _, t4value := range s5case {
		p7s6t.Run(t4value.name, func(p7s6t2 *testing.T) {
			t4p7s6select := F8NewS6Select[S6APPUserModel](p7s6DB)
			res, err2 := t4p7s6select.F4Get(context.Background())
			assert.Equal(p7s6t2, t4value.wantErr, err2)
			if err2 != nil {
				return
			}
			assert.Equal(p7s6t2, t4value.wantRes, res)
		})
	}
}
