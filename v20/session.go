package v20

import (
	"context"
	"database/sql"
)

type I9Session interface {
	f8GetS6Monitor() s6Monitor
	f8DoQueryContext(ctx context.Context, query string, args []any) (*sql.Rows, error)
	f8DoExecContext(ctx context.Context, query string, args []any) (sql.Result, error)
}
