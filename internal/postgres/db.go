package postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/sashaaro/gophkeeper/internal/log"
)

var ErrDBConnection = errors.New("connection db fails")

// Conn - обёртка над sql.DB
type Conn struct {
	db *sql.DB
}

func NewConn(dsn string) (*Conn, error) {
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
	if err := initDatabase(db); err != nil {
		return nil, err
	}
	return &Conn{db: db}, nil
}

func (c *Conn) Close() error {
	return c.db.Close()
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

const sqlCreateTableUser = `CREATE TABLE IF NOT EXISTS "user"
    (
		id uuid not null primary key,
		login varchar(255) not null unique,
		pass varchar(255) not null
	)`

func initDatabase(db *sql.DB) (err error) {
	log.Info("Initialize database")
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	defer func(tx *sql.Tx) {
		if er := tx.Rollback(); er != nil && !errors.Is(er, sql.ErrTxDone) {
			err = errors.Join(err, er)
		}
	}(tx)
	if _, err := tx.ExecContext(ctx, sqlCreateTableUser); err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
