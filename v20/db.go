package v20

import (
	"context"
	"database/sql"
	"orm-go/v20/metadata"
	"orm-go/v20/result"
)

type S6DBOption func(*S6DB)

// S6DB orm 框架的数据库对象
type S6DB struct {
	// p7s6SqlDB 真正的数据库对象
	p7s6SqlDB *sql.DB
	s6Monitor
}

func (p7this *S6DB) f8GetS6Monitor() s6Monitor {
	return p7this.s6Monitor
}

func (p7this *S6DB) f8DoQueryContext(ctx context.Context, query string, args []any) (*sql.Rows, error) {
	return p7this.p7s6SqlDB.QueryContext(ctx, query, args)
}

func (p7this *S6DB) f8DoExecContext(ctx context.Context, query string, args []any) (sql.Result, error) {
	return p7this.p7s6SqlDB.ExecContext(ctx, query, args)
}

// F8NewS6DB 构造 S6DB
func F8NewS6DB(p7s6SqlDB *sql.DB, s5option ...S6DBOption) *S6DB {
	p7s6db := &S6DB{
		p7s6SqlDB: p7s6SqlDB,
		s6Monitor: s6Monitor{
			i9Registry:    metadata.F8NewI9Registry(),
			f8NewI9Result: result.F8NewS6ResultUseUnsafe,
			i9Dialect:     S6MySQLDialect,
		},
	}

	for _, t4value := range s5option {
		t4value(p7s6db)
	}

	return p7s6db
}