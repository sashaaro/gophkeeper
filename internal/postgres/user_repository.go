package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"

	"github.com/google/uuid"
	"github.com/sashaaro/gophkeeper/internal/entity"
	"github.com/sashaaro/gophkeeper/internal/log"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(ctx context.Context, m *entity.User) error {
	if m == nil {
		return ErrModelIsNil
	}
	if m.ID == uuid.Nil {
		m.ID = uuid.Must(uuid.NewV6())
	}
	log.Info(fmt.Sprintf("%s, %s, %s", m.ID, m.Password, m.Login))
	_, err := r.db.ExecContext(ctx, `INSERT INTO "user" (id, login, pass) VALUES ($1, $2, $3)`, m.ID, m.Login, m.Password)
	return err
}

func (r *UserRepository) Get(ctx context.Context, id uuid.UUID, m *entity.User) error {
	if m == nil {
		return ErrModelIsNil
	}
	row := r.db.QueryRowContext(ctx, `SELECT id, login, pass FROM "user" WHERE id = $1`, id)
	if err := row.Err(); err != nil {
		if errors.Is(row.Err(), sql.ErrNoRows) {
			return nil
		}
		return err
	}
	return row.Scan(&m.ID, &m.Login, &m.Password)
}

func (r *UserRepository) GetByLogin(ctx context.Context, login string, m *entity.User) error {
	if m == nil {
		return ErrModelIsNil
	}
	row := r.db.QueryRowContext(ctx, `SELECT id, login, pass FROM "user" WHERE login = $1`, login)
	err := row.Scan(&m.ID, &m.Login, &m.Password)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		return err
	}
	return nil
}
