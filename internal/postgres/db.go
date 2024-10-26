package postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/sashaaro/gophkeeper/internal/log"
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
	err = conn.InTransaction(ctx, func(ctx context.Context, tx *sql.Tx) error {
		return CreateTablesIfNotExist(ctx, tx)
	})
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

const (
	sqlCreateTableUser = `CREATE TABLE IF NOT EXISTS "user" (
		id uuid not null primary key,
		login varchar(255) not null unique,
		pass varchar(255) not null
	)`
	//nolint - this is not harcoded credentials
	sqlCreateEnumSecretKind = `
		DO $$
		BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'secret_kind') THEN
				CREATE TYPE secret_kind AS ENUM
				(
					'credentials',
					'credit_card',
					'text',
					'binary'
				);
			END IF;
		END$$; 
		`
	sqlCreateTableSecret = `CREATE TABLE IF NOT EXISTS "secret" (
	    id uuid not null primary key,
	    user_id uuid not null REFERENCES "user" (id) ON DELETE CASCADE,
	    name varchar(255) not null,
	    kind secret_kind not null,
	    unique (user_id, name)
    )`
	sqlCreateTableTag = `CREATE TABLE IF NOT EXISTS "tag" (
    	id uuid not null primary key,
    	secret_id uuid not null REFERENCES secret (id) ON DELETE CASCADE,
    	name varchar(255) not null,
    	value text not null
    )`
	sqlCreateTableCredentials = `CREATE TABLE IF NOT EXISTS "credentials" (
    	id uuid not null primary key REFERENCES secret (id) ON DELETE CASCADE,
    	login varchar(255) not null,
    	password varchar(255) not null
    )`
	sqlCreateTableCreditCard = `CREATE TABLE IF NOT EXISTS "credit_card" (
    	id uuid not null primary key REFERENCES secret (id) ON DELETE CASCADE,
    	date char(5) not null,
    	name varchar(255) not null,
    	code char(3) not null
    )`
	sqlCreateTableBinary = `CREATE TABLE IF NOT EXISTS "binary_data" (
    	id uuid not null primary key REFERENCES secret (id) ON DELETE CASCADE,
    	is_uploaded bool not null default false
    )`
)

func CreateTablesIfNotExist(ctx context.Context, tx *sql.Tx) (err error) {
	log.Info("Initialize database")
	migrations := []string{
		sqlCreateTableUser,
		sqlCreateEnumSecretKind,
		sqlCreateTableSecret,
		sqlCreateTableTag,
		sqlCreateTableCredentials,
		sqlCreateTableCreditCard,
		sqlCreateTableBinary,
	}
	for _, q := range migrations {
		if _, err := tx.ExecContext(ctx, q); err != nil {
			return err
		}
	}
	return nil
}

func (d *Conn) Ping(ctx context.Context) error {
	return d.db.PingContext(ctx)
}
