package dbwrap

import (
	"context"
	"database/sql"
)

//go:generate mockgen -package ${GOPACKAGE}_mocks -source $GOFILE -destination gen/mocks.go
type Db interface {
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
}
