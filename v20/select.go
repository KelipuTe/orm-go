package v20

import (
	"context"
	"orm-go/v20/result"
)

// S6Select 用于构造 F8Select 语句
type S6Select[T any] struct {
	// s5select 查询子句，对应 select_expr
	s5select []I9SelectExpr
	// tableName 表名
	tableName string
	// s5where WHERE 子句
	s5where []S6WhereCondition
	// s5GroupBy GROUP BY 子句
	s5GroupBy []S6Column
	// s5having GROUP BY 的 HAVING 子句
	s5having []S6WhereCondition
	// s5OrderBy ORDER BY 子句
	s5OrderBy []S6OrderBy
	// limit LIMIT 行数
	limit int
	// offset OFFSET 行数
	offset int

	i9session I9Session
	s6QueryBuilder
}

// F8Select 添加查询子句
func (p7this *S6Select[T]) F8Select(s5expr ...I9SelectExpr) *S6Select[T] {
	if 0 >= len(s5expr) {
		return p7this
	}
	if nil == p7this.s5select {
		p7this.s5select = s5expr
		return p7this
	}
	p7this.s5select = append(p7this.s5select, s5expr...)
	return p7this
}

// F8Where 添加 where 子句
func (p7this *S6Select[T]) F8Where(s5condition ...S6WhereCondition) *S6Select[T] {
	if 0 >= len(s5condition) {
		return p7this
	}
	if nil == p7this.s5where {
		p7this.s5where = s5condition
		return p7this
	}
	p7this.s5where = append(p7this.s5where, s5condition...)
	return p7this
}

// F8GroupBy 添加 group by 子句
func (p7this *S6Select[T]) F8GroupBy(s5column ...S6Column) *S6Select[T] {
	if 0 >= len(s5column) {
		return p7this
	}
	if nil == p7this.s5GroupBy {
		p7this.s5GroupBy = s5column
		return p7this
	}
	p7this.s5GroupBy = append(p7this.s5GroupBy, s5column...)
	return p7this
}

// F8Having 添加 having 子句
func (p7this *S6Select[T]) F8Having(s5condition ...S6WhereCondition) *S6Select[T] {
	if 0 >= len(s5condition) {
		return p7this
	}
	if nil == p7this.s5having {
		p7this.s5having = s5condition
		return p7this
	}
	p7this.s5having = append(p7this.s5having, s5condition...)
	return p7this
}

// F8OrderBy 添加 order by 子句
func (p7this *S6Select[T]) F8OrderBy(s5OrderBy ...S6OrderBy) *S6Select[T] {
	if 0 >= len(s5OrderBy) {
		return p7this
	}
	if nil == p7this.s5OrderBy {
		p7this.s5OrderBy = s5OrderBy
		return p7this
	}
	p7this.s5OrderBy = append(p7this.s5OrderBy, s5OrderBy...)
	return p7this
}

// F8Limit 添加 LIMIT 行数
func (p7this *S6Select[T]) F8Limit(rowCount int) *S6Select[T] {
	p7this.limit = rowCount
	return p7this
}

// F8Offset 添加 OFFSET 行数
func (p7this *S6Select[T]) F8Offset(rowCount int) *S6Select[T] {
	p7this.offset = rowCount
	return p7this
}

func (p7this *S6Select[T]) F8BuildQuery() (*S6Query, error) {
	var err error

	p7this.sqlString.WriteString("SELECT ")

	// 处理查询的列
	err = p7this.f8BuildSelect()
	if nil != err {
		return nil, err
	}

	p7this.sqlString.WriteString(" FROM ")

	// 处理表名
	p7this.sqlString.WriteByte('`')
	p7this.sqlString.WriteString(p7this.tableName)
	p7this.sqlString.WriteByte('`')

	// 处理 where
	if 0 < len(p7this.s5where) {
		p7this.sqlString.WriteString(" WHERE ")
		err = p7this.F8BuildWhereCondition(p7this.s5where)
		if nil != err {
			return nil, err
		}
	}

	// 处理 group by
	if 0 < len(p7this.s5GroupBy) {
		p7this.sqlString.WriteString(" GROUP BY ")
		for i, t4value := range p7this.s5GroupBy {
			if i > 0 {
				p7this.sqlString.WriteByte(',')
			}
			err = t4value.F8BuildColumn(&p7this.s6QueryBuilder)
			if nil != err {
				return nil, err
			}
		}

		// 在有 group by 的情况下，才处理 having
		if 0 < len(p7this.s5having) {
			p7this.sqlString.WriteString(" HAVING ")
			err = p7this.F8BuildWhereCondition(p7this.s5having)
			if nil != err {
				return nil, err
			}
		}
	}

	// 处理 order by
	if 0 < len(p7this.s5OrderBy) {
		p7this.sqlString.WriteString(" ORDER BY ")
		for i, t4value := range p7this.s5OrderBy {
			if i > 0 {
				p7this.sqlString.WriteByte(',')
			}
			err = t4value.F8BuildOrderBy(&p7this.s6QueryBuilder)
			if nil != err {
				return nil, err
			}
		}
	}

	// 处理 limit offset
	if 0 < p7this.limit {
		p7this.sqlString.WriteString(" LIMIT ?")
		p7this.F8AddParameter(p7this.limit)
	}
	if 0 < p7this.offset {
		p7this.sqlString.WriteString(" OFFSET ?")
		p7this.F8AddParameter(p7this.offset)
	}

	p7this.sqlString.WriteString(";")

	p7s6query := &S6Query{
		SQLString: p7this.sqlString.String(),
		S5Value:   p7this.s5value,
	}

	return p7s6query, nil
}

func (p7this *S6Select[T]) f8BuildSelect() error {
	if 0 >= len(p7this.s5select) {
		p7this.sqlString.WriteByte('*')
		return nil
	}

	for i, t4value := range p7this.s5select {
		if i > 0 {
			p7this.sqlString.WriteByte(',')
		}
		err := t4value.F8BuildSelectExpr(&p7this.s6QueryBuilder)
		if nil != err {
			return err
		}

		//switch t4Type := t4value.(type) {
		//case S6Column:
		//	err := p7this.F8BuildColumn(t4Type)
		//	if nil != err {
		//		return err
		//	}
		//case S6Aggregate:
		//	err := p7this.F8BuildAggregate(t4Type)
		//	if nil != err {
		//		return err
		//	}
		//case S6PartRaw:
		//	p7this.sqlString.WriteString(t4Type.SQLString)
		//	if 0 > len(t4Type.S5Value) {
		//		p7this.F8AddParameter(t4Type.S5Value...)
		//	}
		//}
	}
	return nil
}

func F8NewS6Select[T any](i9session I9Session) *S6Select[T] {
	t4p7s6monitor := i9session.f8GetS6Monitor()
	return &S6Select[T]{
		i9session: i9session,
		s6QueryBuilder: s6QueryBuilder{
			s6Monitor: t4p7s6monitor,
			quote:     t4p7s6monitor.i9Dialect.f8GetQuoter(),
		},
		tableName: "table_name",
	}
}

// F4Get 执行查询
func (p7this *S6Select[T]) F4Get(i9ctx context.Context) (*T, error) {
	p7s6query, _ := p7this.F8BuildQuery()

	// 执行查询
	rows, err := p7this.i9session.f8DoQueryContext(i9ctx, p7s6query.SQLString, p7s6query.S5Value...)
	if nil != err {
		return nil, err
	}

	// 处理数据库返回的查询结果
	if !rows.Next() {
		return nil, result.ErrNoRows
	}

	// new 一个类型 T 的变量
	t4p7T := new(T)
	// 获取类型 T 对应的 orm 映射模型
	t4s6model, err2 := p7this.i9session.f8GetS6Monitor().i9Registry.F8Get(t4p7T)
	if nil != err2 {
		return nil, err2
	}
	// 用数据库返回的查询结果构造结构体
	t4result := p7this.i9session.f8GetS6Monitor().f8NewI9Result(t4p7T, t4s6model)
	err4 := t4result.F8SetField(rows)
	return t4p7T, err4
}
