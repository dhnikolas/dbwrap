package dbwrap

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/opentracing/opentracing-go"
)
type Wrapper struct {
	db *sql.DB
}

func New(db *sql.DB) *Wrapper {
	wrapper := &Wrapper{db: db}
	return wrapper
}
func (w *Wrapper) Query(ctx context.Context, sql string, args ...interface{}) (*sql.Rows, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, sql)
	defer span.Finish()
	if len(args) > 0 {
		span.SetTag("args", args)
	}

	return w.db.QueryContext(ctx, sql, args...)
}

func (w *Wrapper) QueryRow(ctx context.Context, sql string, args ...interface{}) *sql.Row {
	span, ctx := opentracing.StartSpanFromContext(ctx, sql)
	defer span.Finish()
	if len(args) > 0 {
		span.SetTag("args", args)
	}

	return w.db.QueryRowContext(ctx, sql, args...)
}

func (w *Wrapper) Exec(ctx context.Context, sql string, args ...interface{}) (sql.Result, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, sql)
	defer span.Finish()
	if len(args) > 0 {
		span.SetTag("args", args)
	}

	return w.db.ExecContext(ctx, sql, args...)
}

func (w *Wrapper) Db() *sql.DB {
	return w.db
}


