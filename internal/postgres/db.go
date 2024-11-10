package postgres

import (
	"context"
	"errors"
	"github.com/jmoiron/sqlx"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var (
	ErrDBConnection     = errors.New("connection db fails")
	ErrModelIsNil       = errors.New("model should not be nil")
	ErrEmptyResultSlice = errors.New("res slice should have positive capacity")
)

// Conn - обёртка над sql.DB

func NewConn(dsn string) (conn *sqlx.DB, err error) {
	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	var ok bool
	row := db.QueryRowContext(ctx, "SELECT true AS ok")
	if err := row.Scan(&ok); err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrDBConnection
	}
	if err := RunMigrations(ctx, db); err != nil {
		return nil, err
	}
	return db, nil
}
