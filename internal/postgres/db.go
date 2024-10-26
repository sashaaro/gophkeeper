package postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var (
	ErrDBConnection     = errors.New("connection db fails")
	ErrModelIsNil       = errors.New("model should not be nil")
	ErrEmptyResultSlice = errors.New("res slice should have positive capacity")
)

// Conn - обёртка над sql.DB
type Conn struct {
	db *sql.DB
}

func NewConn(dsn string) (conn *Conn, err error) {
	db, err := sql.Open("pgx", dsn)
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
	conn = &Conn{db: db}
	if err := RunMigrations(ctx, db); err != nil {
		return nil, err
	}
	return conn, nil
}

func (c *Conn) Close() error {
	return c.db.Close()
}

func (c *Conn) InTransaction(ctx context.Context, fn func(ctx context.Context, tx *sql.Tx) error) (err error) {
	tx, err := c.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}
	defer func(tx *sql.Tx) {
		if er := tx.Rollback(); er != nil && !errors.Is(er, sql.ErrTxDone) {
			err = errors.Join(err, er)
		}
	}(tx)
	if err := fn(ctx, tx); err != nil {
		return err
	}
	return tx.Commit()
}

func (c *Conn) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return c.db.ExecContext(ctx, query, args...)
}

func (c *Conn) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	return c.db.QueryContext(ctx, query, args...)
}

func (c *Conn) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	return c.db.QueryRowContext(ctx, query, args...)
}

func (d *Conn) Ping(ctx context.Context) error {
	return d.db.PingContext(ctx)
}
