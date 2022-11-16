package v20

import (
	"context"
	"database/sql"
	"orm-go/v20/metadata"
	"orm-go/v20/result"
)

type F8S6DBOption func(*S6DB)

func F8DBWithDialect(i9Dialect I9Dialect) F8S6DBOption {
	return func(p7s6DB *S6DB) {
		p7s6DB.s6Monitor.i9Dialect = i9Dialect
	}
}

// S6DB orm 框架的数据库对象
type S6DB struct {
	// p7s6SqlDB 真正的数据库对象
	p7s6SqlDB *sql.DB
	s6Monitor
}

func (p7this *S6DB) f8GetS6Monitor() s6Monitor {
	return p7this.s6Monitor
}

func (p7this *S6DB) f8DoQueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	return p7this.p7s6SqlDB.QueryContext(ctx, query, args...)
}

func (p7this *S6DB) f8DoExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return p7this.p7s6SqlDB.ExecContext(ctx, query, args...)
}

func (p7this *S6DB) F8BeginTx(ctx context.Context, opts *sql.TxOptions) (*S6Tx, error) {
	p7s6Tx, err := p7this.p7s6SqlDB.BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}
	return &S6Tx{p7s6SqlTx: p7s6Tx, p7s6DB: p7this}, nil
}

// F8NewS6DB 构造 S6DB
func F8NewS6DB(p7s6SqlDB *sql.DB, s5Option ...F8S6DBOption) *S6DB {
	p7s6DB := &S6DB{
		p7s6SqlDB: p7s6SqlDB,
		s6Monitor: s6Monitor{
			i9Registry:    metadata.F8NewI9Registry(),
			f8NewI9Result: result.F8NewS6ResultUseUnsafe,
			i9Dialect:     S6MySQLDialect,
		},
	}

	for _, t4value := range s5Option {
		t4value(p7s6DB)
	}

	return p7s6DB
}
