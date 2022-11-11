package v20

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrmSelect_BuildQuery(t *testing.T) {
	s5case := []struct {
		name      string
		i9qb      I9QueryBuilder
		wantQuery *S6Query
		wantErr   error
	}{
		{
			name: "all",
			i9qb: NewS6OrmSelect[S6TestModel](),
			wantQuery: &S6Query{
				SQLString: "SELECT * FROM `table_name`;",
			},
		},
	}
	for _, t4case := range s5case {
		t.Run(t4case.name, func(t *testing.T) {
			p7query, err := t4case.i9qb.F8BuildQuery()
			assert.Equal(t, t4case.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, t4case.wantQuery, p7query)
		})
	}
}

func TestOrmSelect_Operator(t *testing.T) {
	s5case := []struct {
		name      string
		i9qb      I9QueryBuilder
		wantQuery *S6Query
		wantErr   error
	}{
		{
			name: "where_eq",
			i9qb: NewS6OrmSelect[S6TestModel]().
				Where(NewField("Id").EQ(11)),
			wantQuery: &S6Query{
				SQLString:   "SELECT * FROM `table_name` WHERE `Id` = ?;",
				S5parameter: []any{11},
			},
		},
		{
			name: "where_gt",
			i9qb: NewS6OrmSelect[S6TestModel]().
				Where(NewField("Id").GT(11)),
			wantQuery: &S6Query{
				SQLString:   "SELECT * FROM `table_name` WHERE `Id` > ?;",
				S5parameter: []any{11},
			},
		},
		{
			name: "where_lt",
			i9qb: NewS6OrmSelect[S6TestModel]().
				Where(NewField("Id").LT(11)),
			wantQuery: &S6Query{
				SQLString:   "SELECT * FROM `table_name` WHERE `Id` < ?;",
				S5parameter: []any{11},
			},
		},
	}
	for _, t4case := range s5case {
		t.Run(t4case.name, func(t *testing.T) {
			p7query, err := t4case.i9qb.F8BuildQuery()
			assert.Equal(t, t4case.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, t4case.wantQuery, p7query)
		})
	}
}

func TestOrmSelect_Where(t *testing.T) {
	s5case := []struct {
		name      string
		i9qb      I9QueryBuilder
		wantQuery *S6Query
		wantErr   error
	}{
		{
			name: "where_no",
			i9qb: NewS6OrmSelect[S6TestModel]().Where(),
			wantQuery: &S6Query{
				SQLString:   "SELECT * FROM `table_name`;",
				S5parameter: nil,
			},
		},
		{
			name: "where_one",
			i9qb: NewS6OrmSelect[S6TestModel]().
				Where(NewField("Id").EQ(11)),
			wantQuery: &S6Query{
				SQLString:   "SELECT * FROM `table_name` WHERE `Id` = ?;",
				S5parameter: []any{11},
			},
		},
		{
			name: "where_two",
			i9qb: NewS6OrmSelect[S6TestModel]().
				Where(NewField("Id").EQ(11)).
				Where(NewField("Name").EQ("aa")),
			wantQuery: &S6Query{
				SQLString:   "SELECT * FROM `table_name` WHERE (`Id` = ?) AND (`Name` = ?);",
				S5parameter: []any{11, "aa"},
			},
		},
		{
			name: "where_one_and_one",
			i9qb: NewS6OrmSelect[S6TestModel]().
				Where(NewField("Id").EQ(11).And(NewField("Name").EQ("aa"))),
			wantQuery: &S6Query{
				SQLString:   "SELECT * FROM `table_name` WHERE (`Id` = ?) AND (`Name` = ?);",
				S5parameter: []any{11, "aa"},
			},
		},
		{
			name: "where_one_or_one",
			i9qb: NewS6OrmSelect[S6TestModel]().
				Where(NewField("Id").EQ(11).Or(NewField("Name").EQ("aa"))),
			wantQuery: &S6Query{
				SQLString:   "SELECT * FROM `table_name` WHERE (`Id` = ?) OR (`Name` = ?);",
				S5parameter: []any{11, "aa"},
			},
		},
		{
			name: "where_not_one",
			i9qb: NewS6OrmSelect[S6TestModel]().
				Where(Not(NewField("Id").EQ(11))),
			wantQuery: &S6Query{
				SQLString:   "SELECT * FROM `table_name` WHERE  NOT (`Id` = ?);",
				S5parameter: []any{11},
			},
		},
		{
			name: "where_one_and_(one_and_one)",
			i9qb: NewS6OrmSelect[S6TestModel]().
				Where(NewField("Id").EQ(11).And(NewField("Name").EQ("aa").And(NewField("Age").EQ(22)))),
			wantQuery: &S6Query{
				SQLString:   "SELECT * FROM `table_name` WHERE (`Id` = ?) AND ((`Name` = ?) AND (`Age` = ?));",
				S5parameter: []any{11, "aa", 22},
			},
		},
		{
			name: "where_one_or_(one_or_one)",
			i9qb: NewS6OrmSelect[S6TestModel]().
				Where(NewField("Id").EQ(11).And(NewField("Name").EQ("aa").Or(NewField("Age").EQ(22)))),
			wantQuery: &S6Query{
				SQLString:   "SELECT * FROM `table_name` WHERE (`Id` = ?) AND ((`Name` = ?) OR (`Age` = ?));",
				S5parameter: []any{11, "aa", 22},
			},
		},
		{
			name: "where_one_and_(not_one)",
			i9qb: NewS6OrmSelect[S6TestModel]().
				Where(NewField("Id").EQ(11).And(Not(NewField("Name").EQ("aa")))),
			wantQuery: &S6Query{
				SQLString:   "SELECT * FROM `table_name` WHERE (`Id` = ?) AND ( NOT (`Name` = ?));",
				S5parameter: []any{11, "aa"},
			},
		},
	}
	for _, t4case := range s5case {
		t.Run(t4case.name, func(t *testing.T) {
			p7query, err := t4case.i9qb.F8BuildQuery()
			assert.Equal(t, t4case.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, t4case.wantQuery, p7query)
		})
	}
}

func TestOrmSelect_GroupBy(t *testing.T) {
	s5case := []struct {
		name      string
		i9qb      I9QueryBuilder
		wantQuery *S6Query
		wantErr   error
	}{
		{
			name: "group_by_no",
			i9qb: NewS6OrmSelect[S6TestModel]().GroupBy(),
			wantQuery: &S6Query{
				SQLString:   "SELECT * FROM `table_name`;",
				S5parameter: nil,
			},
		},
		{
			name: "group_by_one",
			i9qb: NewS6OrmSelect[S6TestModel]().
				GroupBy(NewField("Age")),
			wantQuery: &S6Query{
				SQLString:   "SELECT * FROM `table_name` GROUP BY `Age`;",
				S5parameter: nil,
			},
		},
		{
			name: "group_by_two",
			i9qb: NewS6OrmSelect[S6TestModel]().
				GroupBy(NewField("Age")).GroupBy(NewField("Sex")),
			wantQuery: &S6Query{
				SQLString:   "SELECT * FROM `table_name` GROUP BY `Age`,`Sex`;",
				S5parameter: nil,
			},
		},
	}
	for _, t4case := range s5case {
		t.Run(t4case.name, func(t *testing.T) {
			p7query, err := t4case.i9qb.F8BuildQuery()
			assert.Equal(t, t4case.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, t4case.wantQuery, p7query)
		})
	}
}

func TestOrmSelect_Having(t *testing.T) {
	s5case := []struct {
		name      string
		i9qb      I9QueryBuilder
		wantQuery *S6Query
		wantErr   error
	}{
		{
			name: "having_no",
			i9qb: NewS6OrmSelect[S6TestModel]().Having(),
			wantQuery: &S6Query{
				SQLString:   "SELECT * FROM `table_name`;",
				S5parameter: nil,
			},
		},
		{
			name: "having_no_group_by",
			i9qb: NewS6OrmSelect[S6TestModel]().
				Having(NewField("Age").GT(22)),
			wantQuery: &S6Query{
				SQLString:   "SELECT * FROM `table_name`;",
				S5parameter: nil,
			},
		},
		{
			name: "having_one",
			i9qb: NewS6OrmSelect[S6TestModel]().
				GroupBy(NewField("Age")).
				Having(NewField("Id").EQ(11)),
			wantQuery: &S6Query{
				SQLString:   "SELECT * FROM `table_name` GROUP BY `Age` HAVING `Id` = ?;",
				S5parameter: []any{11},
			},
		},
		{
			name: "group_by_two_having_one",
			i9qb: NewS6OrmSelect[S6TestModel]().
				GroupBy(NewField("Age")).GroupBy(NewField("Sex")).
				Having(NewField("Id").EQ(11)),
			wantQuery: &S6Query{
				SQLString:   "SELECT * FROM `table_name` GROUP BY `Age`,`Sex` HAVING `Id` = ?;",
				S5parameter: []any{11},
			},
		},
		{
			name: "group_by_two_having_two",
			i9qb: NewS6OrmSelect[S6TestModel]().
				GroupBy(NewField("Age")).GroupBy(NewField("Sex")).
				Having(NewField("Id").EQ(11)).Having(NewField("Name").EQ("aa")),
			wantQuery: &S6Query{
				SQLString:   "SELECT * FROM `table_name` GROUP BY `Age`,`Sex` HAVING (`Id` = ?) AND (`Name` = ?);",
				S5parameter: []any{11, "aa"},
			},
		},
	}
	for _, t4case := range s5case {
		t.Run(t4case.name, func(t *testing.T) {
			p7query, err := t4case.i9qb.F8BuildQuery()
			assert.Equal(t, t4case.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, t4case.wantQuery, p7query)
		})
	}
}

func TestOrmSelect_OrderBy(t *testing.T) {
	s5case := []struct {
		name      string
		i9qb      I9QueryBuilder
		wantQuery *S6Query
		wantErr   error
	}{
		{
			name: "order_by_no",
			i9qb: NewS6OrmSelect[S6TestModel]().OrderBy(),
			wantQuery: &S6Query{
				SQLString:   "SELECT * FROM `table_name`;",
				S5parameter: nil,
			},
		},
		{
			name: "order_by_one_asc",
			i9qb: NewS6OrmSelect[S6TestModel]().
				OrderBy(Asc("Name")),
			wantQuery: &S6Query{
				SQLString:   "SELECT * FROM `table_name` ORDER BY `Name` ASC;",
				S5parameter: nil,
			},
		},
		{
			name: "order_by_one_desc",
			i9qb: NewS6OrmSelect[S6TestModel]().
				OrderBy(Desc("Name")),
			wantQuery: &S6Query{
				SQLString:   "SELECT * FROM `table_name` ORDER BY `Name` DESC;",
				S5parameter: nil,
			},
		},
		{
			name: "order_by_two_asc_desc",
			i9qb: NewS6OrmSelect[S6TestModel]().
				OrderBy(Asc("Name")).OrderBy(Desc("Age")),
			wantQuery: &S6Query{
				SQLString:   "SELECT * FROM `table_name` ORDER BY `Name` ASC,`Age` DESC;",
				S5parameter: nil,
			},
		},
	}
	for _, t4case := range s5case {
		t.Run(t4case.name, func(t *testing.T) {
			p7query, err := t4case.i9qb.F8BuildQuery()
			assert.Equal(t, t4case.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, t4case.wantQuery, p7query)
		})
	}
}

func TestOrmSelect_OffsetLimit(t *testing.T) {
	s5case := []struct {
		name      string
		i9qb      I9QueryBuilder
		wantQuery *S6Query
		wantErr   error
	}{
		{
			name: "limit",
			i9qb: NewS6OrmSelect[S6TestModel]().Limit(11),
			wantQuery: &S6Query{
				SQLString:   "SELECT * FROM `table_name` LIMIT ?;",
				S5parameter: []any{11},
			},
		},
		{
			name: "offset",
			i9qb: NewS6OrmSelect[S6TestModel]().Offset(111),
			wantQuery: &S6Query{
				SQLString:   "SELECT * FROM `table_name` OFFSET ?;",
				S5parameter: []any{111},
			},
		},
		{
			name: "limit_offset",
			i9qb: NewS6OrmSelect[S6TestModel]().Limit(11).Offset(111),
			wantQuery: &S6Query{
				SQLString:   "SELECT * FROM `table_name` LIMIT ? OFFSET ?;",
				S5parameter: []any{11, 111},
			},
		},
	}
	for _, t4case := range s5case {
		t.Run(t4case.name, func(t *testing.T) {
			p7query, err := t4case.i9qb.F8BuildQuery()
			assert.Equal(t, t4case.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, t4case.wantQuery, p7query)
		})
	}
}

func TestOrmSelect_Select(t *testing.T) {
	s5case := []struct {
		name      string
		i9qb      I9QueryBuilder
		wantQuery *S6Query
		wantErr   error
	}{
		{
			name: "select_one_column",
			i9qb: NewS6OrmSelect[S6TestModel]().
				Select(NewField("Id")),
			wantQuery: &S6Query{
				SQLString:   "SELECT `Id` FROM `table_name`;",
				S5parameter: nil,
			},
		},
		{
			name: "select_two_column",
			i9qb: NewS6OrmSelect[S6TestModel]().
				Select(NewField("Id")).Select(NewField("Name")),
			wantQuery: &S6Query{
				SQLString:   "SELECT `Id`,`Name` FROM `table_name`;",
				S5parameter: nil,
			},
		},
	}
	for _, t4case := range s5case {
		t.Run(t4case.name, func(t *testing.T) {
			p7query, err := t4case.i9qb.F8BuildQuery()
			assert.Equal(t, t4case.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, t4case.wantQuery, p7query)
		})
	}
}

func TestOrmSelect_Aggregate(t *testing.T) {
	s5case := []struct {
		name      string
		i9qb      I9QueryBuilder
		wantQuery *S6Query
		wantErr   error
	}{
		{
			name: "select_one_aggregate",
			i9qb: NewS6OrmSelect[S6TestModel]().
				Select(Count("Id")),
			wantQuery: &S6Query{
				SQLString:   "SELECT COUNT(`Id`) FROM `table_name`;",
				S5parameter: nil,
			},
		},
		{
			name: "select_two_aggregate",
			i9qb: NewS6OrmSelect[S6TestModel]().
				Select(Count("Id")).Select(Avg("Age")),
			wantQuery: &S6Query{
				SQLString:   "SELECT COUNT(`Id`),AVG(`Age`) FROM `table_name`;",
				S5parameter: nil,
			},
		},
		{
			name: "having_one_aggregate",
			i9qb: NewS6OrmSelect[S6TestModel]().
				GroupBy(NewField("Age")).
				Having(Count("Id").GT(5)),
			wantQuery: &S6Query{
				SQLString:   "SELECT * FROM `table_name` GROUP BY `Age` HAVING COUNT(`Id`) > ?;",
				S5parameter: []any{5},
			},
		},
	}
	for _, t4case := range s5case {
		t.Run(t4case.name, func(t *testing.T) {
			p7query, err := t4case.i9qb.F8BuildQuery()
			assert.Equal(t, t4case.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, t4case.wantQuery, p7query)
		})
	}
}

func TestOrmSelect_Raw(t *testing.T) {
	s5case := []struct {
		name      string
		i9qb      I9QueryBuilder
		wantQuery *S6Query
		wantErr   error
	}{
		{
			name: "select_raw",
			i9qb: NewS6OrmSelect[S6TestModel]().
				Select(NewRaw("DISTINCT(Id)")),
			wantQuery: &S6Query{
				SQLString:   "SELECT DISTINCT(Id) FROM `table_name`;",
				S5parameter: nil,
			},
		},
		{
			name: "where_raw",
			i9qb: NewS6OrmSelect[S6TestModel]().
				Where(NewRaw("Id > ?", 11).toPredicate()),
			wantQuery: &S6Query{
				SQLString:   "SELECT * FROM `table_name` WHERE Id > ?;",
				S5parameter: []any{11},
			},
		},
		{
			name: "where_raw_and_one",
			i9qb: NewS6OrmSelect[S6TestModel]().
				Where(NewRaw("Id > ?", 11).toPredicate().And(NewField("Name").EQ("aa"))),
			wantQuery: &S6Query{
				SQLString:   "SELECT * FROM `table_name` WHERE (Id > ?) AND (`Name` = ?);",
				S5parameter: []any{11, "aa"},
			},
		},
		{
			name: "having_raw",
			i9qb: NewS6OrmSelect[S6TestModel]().
				GroupBy(NewField("Age")).
				Having(NewRaw("COUNT(Id) > ?", 5).toPredicate()),
			wantQuery: &S6Query{
				SQLString:   "SELECT * FROM `table_name` GROUP BY `Age` HAVING COUNT(Id) > ?;",
				S5parameter: []any{5},
			},
		},
	}
	for _, t4case := range s5case {
		t.Run(t4case.name, func(t *testing.T) {
			p7query, err := t4case.i9qb.F8BuildQuery()
			assert.Equal(t, t4case.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, t4case.wantQuery, p7query)
		})
	}
}

func TestS6SelectorF8Get(p7s6t *testing.T) {
	// 构造 mock 数据库连接
	p7s6dbMock, sqlMock, err := sqlmock.New()
	if nil != err {
		p7s6t.Fatal(err)
	}
	defer func() {
		_ = p7s6dbMock.Close()
	}()

	p7s6OrmDB := F8NewS6DB(p7s6dbMock)

	s5case := []struct {
		name      string
		sqlString string
		rowsMock  *sqlmock.Rows
		errMock   error
		valueWant *S6TestModel
		errWant   error
	}{
		{
			name:      "normal_sql",
			sqlString: "SELECT .*",
			rowsMock: func() *sqlmock.Rows {
				rows := sqlmock.NewRows([]string{"id", "name", "age", "sex"})
				rows.AddRow([]byte("11"), []byte("aa"), []byte("22"), []byte("1"))
				return rows
			}(),
			valueWant: &S6TestModel{
				Id:   11,
				Name: "aa",
				Age:  22,
				Sex:  1,
			},
		},
	}
	// 把预设的查询结果装进 mock
	for _, t4case := range s5case {
		t4p7eq := sqlMock.ExpectQuery(t4case.sqlString)
		if nil != t4case.errMock {
			t4p7eq.WillReturnError(t4case.errMock)
		} else {
			t4p7eq.WillReturnRows(t4case.rowsMock)
		}
	}

	for _, t4case := range s5case {
		p7s6t.Run(t4case.name, func(p7s6t *testing.T) {
			s6query := S6Query{SQLString: t4case.sqlString, S5parameter: []any{}}
			res, err := F8NewS6OrmSelect[S6TestModel](p7s6OrmDB, s6query).F4Get(context.Background())
			assert.Equal(p7s6t, t4case.errWant, err)
			if err != nil {
				return
			}
			assert.Equal(p7s6t, t4case.valueWant, res)
		})
	}
}
