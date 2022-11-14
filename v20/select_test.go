package v20

import (
	"github.com/stretchr/testify/assert"
	"orm-go/v20/clause"
	"testing"
)

func TestS6SelectBuildQuery(p7t *testing.T) {
	s5case := []struct {
		name         string
		queryBuilder clause.I9QueryBuilder
		wantQuery    *clause.S6Query
		wantErr      error
	}{
		{
			name:         "all",
			queryBuilder: NewS6Select[S6TestModel](nil),
			wantQuery: &clause.S6Query{
				SQLString: "SELECT * FROM `table_name`;",
			},
		},
	}

	for _, t4value := range s5case {
		p7t.Run(t4value.name, func(p7t2 *testing.T) {
			p7query, err := t4value.queryBuilder.F8BuildQuery()
			assert.Equal(p7t2, t4value.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(p7t2, t4value.wantQuery, p7query)
		})
	}
}

func TestS6SelectOperator(p7t *testing.T) {
	s5case := []struct {
		name         string
		queryBuilder clause.I9QueryBuilder
		wantQuery    *clause.S6Query
		wantErr      error
	}{
		{
			name: "where_equal",
			queryBuilder: NewS6Select[S6TestModel](nil).
				F8Where(clause.F8NewS6Column("Id").F8EQ(11)),
			wantQuery: &clause.S6Query{
				SQLString: "SELECT * FROM `table_name` WHERE `Id` = ?;",
				S5Value:   []any{11},
			},
		},
		{
			name: "where_greater_than",
			queryBuilder: NewS6Select[S6TestModel](nil).
				F8Where(clause.F8NewS6Column("Id").F8GT(11)),
			wantQuery: &clause.S6Query{
				SQLString: "SELECT * FROM `table_name` WHERE `Id` > ?;",
				S5Value:   []any{11},
			},
		},
		{
			name: "where_less_than",
			queryBuilder: NewS6Select[S6TestModel](nil).
				F8Where(clause.F8NewS6Column("Id").F8LT(11)),
			wantQuery: &clause.S6Query{
				SQLString: "SELECT * FROM `table_name` WHERE `Id` < ?;",
				S5Value:   []any{11},
			},
		},
	}

	for _, t4value := range s5case {
		p7t.Run(t4value.name, func(p7t2 *testing.T) {
			p7query, err := t4value.queryBuilder.F8BuildQuery()
			assert.Equal(p7t2, t4value.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(p7t2, t4value.wantQuery, p7query)
		})
	}
}

func TestS6SelectWhere(p7t *testing.T) {
	s5case := []struct {
		name         string
		queryBuilder clause.I9QueryBuilder
		wantQuery    *clause.S6Query
		wantErr      error
	}{
		{
			name: "where_no",
			queryBuilder: NewS6Select[S6TestModel](nil).
				F8Where(),
			wantQuery: &clause.S6Query{
				SQLString: "SELECT * FROM `table_name`;",
				S5Value:   nil,
			},
		},
		{
			name: "where_one",
			queryBuilder: NewS6Select[S6TestModel](nil).
				F8Where(clause.F8NewS6Column("Id").F8EQ(11)),
			wantQuery: &clause.S6Query{
				SQLString: "SELECT * FROM `table_name` WHERE `Id` = ?;",
				S5Value:   []any{11},
			},
		},
		{
			name: "where_two",
			queryBuilder: NewS6Select[S6TestModel](nil).
				F8Where(clause.F8NewS6Column("Id").F8EQ(11)).
				F8Where(clause.F8NewS6Column("S6Column").F8EQ("aa")),
			wantQuery: &clause.S6Query{
				SQLString: "SELECT * FROM `table_name` WHERE (`Id` = ?) AND (`S6Column` = ?);",
				S5Value:   []any{11, "aa"},
			},
		},
		{
			name: "where_one_and_one",
			queryBuilder: NewS6Select[S6TestModel](nil).
				F8Where(clause.F8NewS6Column("Id").F8EQ(11).And(clause.F8NewS6Column("S6Column").F8EQ("aa"))),
			wantQuery: &clause.S6Query{
				SQLString: "SELECT * FROM `table_name` WHERE (`Id` = ?) AND (`S6Column` = ?);",
				S5Value:   []any{11, "aa"},
			},
		},
		{
			name: "where_one_or_one",
			queryBuilder: NewS6Select[S6TestModel](nil).
				F8Where(clause.F8NewS6Column("Id").F8EQ(11).Or(clause.F8NewS6Column("S6Column").F8EQ("aa"))),
			wantQuery: &clause.S6Query{
				SQLString: "SELECT * FROM `table_name` WHERE (`Id` = ?) OR (`S6Column` = ?);",
				S5Value:   []any{11, "aa"},
			},
		},
		{
			name: "where_not_one",
			queryBuilder: NewS6Select[S6TestModel](nil).
				F8Where(clause.Not(clause.F8NewS6Column("Id").F8EQ(11))),
			wantQuery: &clause.S6Query{
				SQLString: "SELECT * FROM `table_name` WHERE  NOT (`Id` = ?);",
				S5Value:   []any{11},
			},
		},
		{
			name: "where_one_and_(one_and_one)",
			queryBuilder: NewS6Select[S6TestModel](nil).
				F8Where(clause.F8NewS6Column("Id").F8EQ(11).And(clause.F8NewS6Column("S6Column").F8EQ("aa").And(clause.F8NewS6Column("Age").F8EQ(22)))),
			wantQuery: &clause.S6Query{
				SQLString: "SELECT * FROM `table_name` WHERE (`Id` = ?) AND ((`S6Column` = ?) AND (`Age` = ?));",
				S5Value:   []any{11, "aa", 22},
			},
		},
		{
			name: "where_one_or_(one_or_one)",
			queryBuilder: NewS6Select[S6TestModel](nil).
				F8Where(clause.F8NewS6Column("Id").F8EQ(11).And(clause.F8NewS6Column("S6Column").F8EQ("aa").Or(clause.F8NewS6Column("Age").F8EQ(22)))),
			wantQuery: &clause.S6Query{
				SQLString: "SELECT * FROM `table_name` WHERE (`Id` = ?) AND ((`S6Column` = ?) OR (`Age` = ?));",
				S5Value:   []any{11, "aa", 22},
			},
		},
		{
			name: "where_one_and_(not_one)",
			queryBuilder: NewS6Select[S6TestModel](nil).
				F8Where(clause.F8NewS6Column("Id").F8EQ(11).And(clause.Not(clause.F8NewS6Column("S6Column").F8EQ("aa")))),
			wantQuery: &clause.S6Query{
				SQLString: "SELECT * FROM `table_name` WHERE (`Id` = ?) AND ( NOT (`S6Column` = ?));",
				S5Value:   []any{11, "aa"},
			},
		},
	}

	for _, t4value := range s5case {
		p7t.Run(t4value.name, func(p7t2 *testing.T) {
			p7query, err := t4value.queryBuilder.F8BuildQuery()
			assert.Equal(p7t2, t4value.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(p7t2, t4value.wantQuery, p7query)
		})
	}
}

func TestS6SelectGroupBy(p7t *testing.T) {
	s5case := []struct {
		name         string
		queryBuilder clause.I9QueryBuilder
		wantQuery    *clause.S6Query
		wantErr      error
	}{
		{
			name: "group_by_no",
			queryBuilder: NewS6Select[S6TestModel](nil).
				F8GroupBy(),
			wantQuery: &clause.S6Query{
				SQLString: "SELECT * FROM `table_name`;",
				S5Value:   nil,
			},
		},
		{
			name: "group_by_one",
			queryBuilder: NewS6Select[S6TestModel](nil).
				F8GroupBy(clause.F8NewS6Column("Age")),
			wantQuery: &clause.S6Query{
				SQLString: "SELECT * FROM `table_name` GROUP BY `Age`;",
				S5Value:   nil,
			},
		},
		{
			name: "group_by_two",
			queryBuilder: NewS6Select[S6TestModel](nil).
				F8GroupBy(clause.F8NewS6Column("Age")).F8GroupBy(clause.F8NewS6Column("Sex")),
			wantQuery: &clause.S6Query{
				SQLString: "SELECT * FROM `table_name` GROUP BY `Age`,`Sex`;",
				S5Value:   nil,
			},
		},
	}

	for _, t4value := range s5case {
		p7t.Run(t4value.name, func(p7t2 *testing.T) {
			p7query, err := t4value.queryBuilder.F8BuildQuery()
			assert.Equal(p7t2, t4value.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(p7t2, t4value.wantQuery, p7query)
		})
	}
}

func TestS6SelectHaving(p7t *testing.T) {
	s5case := []struct {
		name         string
		queryBuilder clause.I9QueryBuilder
		wantQuery    *clause.S6Query
		wantErr      error
	}{
		{
			name:         "having_no",
			queryBuilder: NewS6Select[S6TestModel](nil).F8Having(),
			wantQuery: &clause.S6Query{
				SQLString: "SELECT * FROM `table_name`;",
				S5Value:   nil,
			},
		},
		{
			name: "having_no_group_by",
			queryBuilder: NewS6Select[S6TestModel](nil).
				F8Having(clause.F8NewS6Column("Age").F8GT(22)),
			wantQuery: &clause.S6Query{
				SQLString: "SELECT * FROM `table_name`;",
				S5Value:   nil,
			},
		},
		{
			name: "having_one",
			queryBuilder: NewS6Select[S6TestModel](nil).
				F8GroupBy(clause.F8NewS6Column("Age")).
				F8Having(clause.F8NewS6Column("Id").F8EQ(11)),
			wantQuery: &clause.S6Query{
				SQLString: "SELECT * FROM `table_name` GROUP BY `Age` HAVING `Id` = ?;",
				S5Value:   []any{11},
			},
		},
		{
			name: "group_by_two_having_one",
			queryBuilder: NewS6Select[S6TestModel](nil).
				F8GroupBy(clause.F8NewS6Column("Age")).F8GroupBy(clause.F8NewS6Column("Sex")).
				F8Having(clause.F8NewS6Column("Id").F8EQ(11)),
			wantQuery: &clause.S6Query{
				SQLString: "SELECT * FROM `table_name` GROUP BY `Age`,`Sex` HAVING `Id` = ?;",
				S5Value:   []any{11},
			},
		},
		{
			name: "group_by_two_having_two",
			queryBuilder: NewS6Select[S6TestModel](nil).
				F8GroupBy(clause.F8NewS6Column("Age")).F8GroupBy(clause.F8NewS6Column("Sex")).
				F8Having(clause.F8NewS6Column("Id").F8EQ(11)).F8Having(clause.F8NewS6Column("S6Column").F8EQ("aa")),
			wantQuery: &clause.S6Query{
				SQLString: "SELECT * FROM `table_name` GROUP BY `Age`,`Sex` HAVING (`Id` = ?) AND (`S6Column` = ?);",
				S5Value:   []any{11, "aa"},
			},
		},
	}

	for _, t4value := range s5case {
		p7t.Run(t4value.name, func(p7t2 *testing.T) {
			p7query, err := t4value.queryBuilder.F8BuildQuery()
			assert.Equal(p7t2, t4value.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(p7t2, t4value.wantQuery, p7query)
		})
	}
}

func TestS6SelectOrderBy(p7t *testing.T) {
	s5case := []struct {
		name         string
		queryBuilder clause.I9QueryBuilder
		wantQuery    *clause.S6Query
		wantErr      error
	}{
		{
			name:         "order_by_no",
			queryBuilder: NewS6Select[S6TestModel](nil).F8OrderBy(),
			wantQuery: &clause.S6Query{
				SQLString: "SELECT * FROM `table_name`;",
				S5Value:   nil,
			},
		},
		{
			name: "order_by_one_asc",
			queryBuilder: NewS6Select[S6TestModel](nil).
				F8OrderBy(clause.Asc("S6Column")),
			wantQuery: &clause.S6Query{
				SQLString: "SELECT * FROM `table_name` ORDER BY `S6Column` ASC;",
				S5Value:   nil,
			},
		},
		{
			name: "order_by_one_desc",
			queryBuilder: NewS6Select[S6TestModel](nil).
				F8OrderBy(clause.Desc("S6Column")),
			wantQuery: &clause.S6Query{
				SQLString: "SELECT * FROM `table_name` ORDER BY `S6Column` DESC;",
				S5Value:   nil,
			},
		},
		{
			name: "order_by_two_asc_desc",
			queryBuilder: NewS6Select[S6TestModel](nil).
				F8OrderBy(clause.Asc("S6Column")).F8OrderBy(clause.Desc("Age")),
			wantQuery: &clause.S6Query{
				SQLString: "SELECT * FROM `table_name` ORDER BY `S6Column` ASC,`Age` DESC;",
				S5Value:   nil,
			},
		},
	}

	for _, t4value := range s5case {
		p7t.Run(t4value.name, func(p7t *testing.T) {
			p7query, err := t4value.queryBuilder.F8BuildQuery()
			assert.Equal(p7t, t4value.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(p7t, t4value.wantQuery, p7query)
		})
	}
}

func TestS6SelectOffsetLimit(p7t *testing.T) {
	s5case := []struct {
		name         string
		queryBuilder clause.I9QueryBuilder
		wantQuery    *clause.S6Query
		wantErr      error
	}{
		{
			name:         "limit",
			queryBuilder: NewS6Select[S6TestModel](nil).F8Limit(11),
			wantQuery: &clause.S6Query{
				SQLString: "SELECT * FROM `table_name` LIMIT ?;",
				S5Value:   []any{11},
			},
		},
		{
			name:         "offset",
			queryBuilder: NewS6Select[S6TestModel](nil).F8Offset(111),
			wantQuery: &clause.S6Query{
				SQLString: "SELECT * FROM `table_name` OFFSET ?;",
				S5Value:   []any{111},
			},
		},
		{
			name:         "limit_offset",
			queryBuilder: NewS6Select[S6TestModel](nil).F8Limit(11).F8Offset(111),
			wantQuery: &clause.S6Query{
				SQLString: "SELECT * FROM `table_name` LIMIT ? OFFSET ?;",
				S5Value:   []any{11, 111},
			},
		},
	}

	for _, t4value := range s5case {
		p7t.Run(t4value.name, func(p7t *testing.T) {
			p7query, err := t4value.queryBuilder.F8BuildQuery()
			assert.Equal(p7t, t4value.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(p7t, t4value.wantQuery, p7query)
		})
	}
}

func TestS6SelectSelect(p7t *testing.T) {
	s5case := []struct {
		name         string
		queryBuilder clause.I9QueryBuilder
		wantQuery    *clause.S6Query
		wantErr      error
	}{
		{
			name: "select_one_column",
			queryBuilder: NewS6Select[S6TestModel](nil).
				F8Select(clause.F8NewS6Column("Id")),
			wantQuery: &clause.S6Query{
				SQLString: "SELECT `Id` FROM `table_name`;",
				S5Value:   nil,
			},
		},
		{
			name: "select_two_column",
			queryBuilder: NewS6Select[S6TestModel](nil).
				F8Select(clause.F8NewS6Column("Id")).F8Select(clause.F8NewS6Column("S6Column")),
			wantQuery: &clause.S6Query{
				SQLString: "SELECT `Id`,`S6Column` FROM `table_name`;",
				S5Value:   nil,
			},
		},
	}

	for _, t4value := range s5case {
		p7t.Run(t4value.name, func(p7t2 *testing.T) {
			p7query, err := t4value.queryBuilder.F8BuildQuery()
			assert.Equal(p7t2, t4value.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(p7t2, t4value.wantQuery, p7query)
		})
	}
}

func TestS6SelectAggregate(p7t *testing.T) {
	s5case := []struct {
		name         string
		queryBuilder clause.I9QueryBuilder
		wantQuery    *clause.S6Query
		wantErr      error
	}{
		{
			name: "select_one_aggregate",
			queryBuilder: NewS6Select[S6TestModel](nil).
				F8Select(clause.Count("Id")),
			wantQuery: &clause.S6Query{
				SQLString: "SELECT COUNT(`Id`) FROM `table_name`;",
				S5Value:   nil,
			},
		},
		{
			name: "select_two_aggregate",
			queryBuilder: NewS6Select[S6TestModel](nil).
				F8Select(clause.Count("Id")).F8Select(clause.Avg("Age")),
			wantQuery: &clause.S6Query{
				SQLString: "SELECT COUNT(`Id`),AVG(`Age`) FROM `table_name`;",
				S5Value:   nil,
			},
		},
		{
			name: "having_one_aggregate",
			queryBuilder: NewS6Select[S6TestModel](nil).
				F8GroupBy(clause.F8NewS6Column("Age")).
				F8Having(clause.Count("Id").GreaterThan(5)),
			wantQuery: &clause.S6Query{
				SQLString: "SELECT * FROM `table_name` GROUP BY `Age` HAVING COUNT(`Id`) > ?;",
				S5Value:   []any{5},
			},
		},
	}

	for _, t4value := range s5case {
		p7t.Run(t4value.name, func(p7t2 *testing.T) {
			p7query, err := t4value.queryBuilder.F8BuildQuery()
			assert.Equal(p7t2, t4value.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(p7t2, t4value.wantQuery, p7query)
		})
	}
}

func TestS6SelectPartRaw(p7t *testing.T) {
	s5case := []struct {
		name         string
		queryBuilder clause.I9QueryBuilder
		wantQuery    *clause.S6Query
		wantErr      error
	}{
		{
			name: "select_part_raw",
			queryBuilder: NewS6Select[S6TestModel](nil).
				F8Select(clause.NewS6PartRaw("DISTINCT(Id)")),
			wantQuery: &clause.S6Query{
				SQLString: "SELECT DISTINCT(Id) FROM `table_name`;",
				S5Value:   nil,
			},
		},
		{
			name: "where_part_raw",
			queryBuilder: NewS6Select[S6TestModel](nil).
				F8Where(clause.NewS6PartRaw("Id > ?", 11).ToPredicate()),
			wantQuery: &clause.S6Query{
				SQLString: "SELECT * FROM `table_name` WHERE Id > ?;",
				S5Value:   []any{11},
			},
		},
		{
			name: "where_part_raw_and_one",
			queryBuilder: NewS6Select[S6TestModel](nil).
				F8Where(clause.NewS6PartRaw("Id > ?", 11).ToPredicate().And(clause.F8NewS6Column("S6Column").F8EQ("aa"))),
			wantQuery: &clause.S6Query{
				SQLString: "SELECT * FROM `table_name` WHERE (Id > ?) AND (`S6Column` = ?);",
				S5Value:   []any{11, "aa"},
			},
		},
		{
			name: "having_part_raw",
			queryBuilder: NewS6Select[S6TestModel](nil).
				F8GroupBy(clause.F8NewS6Column("Age")).
				F8Having(clause.NewS6PartRaw("COUNT(Id) > ?", 5).ToPredicate()),
			wantQuery: &clause.S6Query{
				SQLString: "SELECT * FROM `table_name` GROUP BY `Age` HAVING COUNT(Id) > ?;",
				S5Value:   []any{5},
			},
		},
	}

	for _, t4value := range s5case {
		p7t.Run(t4value.name, func(p7t2 *testing.T) {
			p7query, err := t4value.queryBuilder.F8BuildQuery()
			assert.Equal(p7t2, t4value.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(p7t2, t4value.wantQuery, p7query)
		})
	}
}

//func TestS6SelectorF8Get(p7s6t *testing.T) {
//	// 构造 mock 数据库连接
//	p7s6dbMock, sqlMock, err := sqlmock.New()
//	if nil != err {
//		p7s6t.Fatal(err)
//	}
//	defer func() {
//		_ = p7s6dbMock.Close()
//	}()
//
//	p7s6OrmDB := F8NewS6DB(p7s6dbMock)
//
//	s5case := []struct {
//		name      string
//		sqlString string
//		rowsMock  *sqlmock.Rows
//		errMock   error
//		valueWant *S6TestModel
//		errWant   error
//	}{
//		{
//			name:      "normal_sql",
//			sqlString: "SELECT .*",
//			rowsMock: func() *sqlmock.Rows {
//				rows := sqlmock.NewRows([]string{"id", "name", "age", "sex"})
//				rows.AddRow([]byte("11"), []byte("aa"), []byte("22"), []byte("1"))
//				return rows
//			}(),
//			valueWant: &S6TestModel{
//				Id:   11,
//				Name: "aa",
//				Age:  22,
//				Sex:  1,
//			},
//		},
//	}
//	// 把预设的查询结果装进 mock
//	for _, t4case := range s5case {
//		t4p7eq := sqlMock.ExpectQuery(t4case.sqlString)
//		if nil != t4case.errMock {
//			t4p7eq.WillReturnError(t4case.errMock)
//		} else {
//			t4p7eq.WillReturnRows(t4case.rowsMock)
//		}
//	}
//
//	for _, t4case := range s5case {
//		p7s6t.Run(t4case.name, func(p7s6t *testing.T) {
//			s6query := clause.S6Query{SQLString: t4case.sqlString, S5Value: []any{}}
//			res, err := F8NewS6OrmSelect[S6TestModel](p7s6OrmDB, s6query).F4Get(context.Background())
//			assert.Equal(p7s6t, t4case.errWant, err)
//			if err != nil {
//				return
//			}
//			assert.Equal(p7s6t, t4case.valueWant, res)
//		})
//	}
//}
