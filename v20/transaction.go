package v20

import (
	"context"
	"database/sql"
)

type S6Tx struct {
	p7s6SqlTx *sql.Tx
	p7s6DB    *S6DB
}

func (p7this *S6Tx) f8GetS6Monitor() s6Monitor {
	return p7this.p7s6DB.s6Monitor
}

func (p7this *S6Tx) f8DoQueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	return p7this.p7s6SqlTx.QueryContext(ctx, query, args...)
}

func (p7this *S6Tx) f8DoExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return p7this.p7s6SqlTx.ExecContext(ctx, query, args...)
}

func (p7this *S6Tx) Commit() error {
	return p7this.p7s6SqlTx.Commit()
}

func (p7this *S6Tx) Rollback() error {
	return p7this.p7s6SqlTx.Rollback()
}

func (p7this *S6Tx) RollbackIfNotCommit() error {
	err := p7this.p7s6SqlTx.Rollback()
	if err != sql.ErrTxDone {
		return err
	}
	return nil
}
