package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/sashaaro/gophkeeper/internal/entity"
	"github.com/sashaaro/gophkeeper/internal/log"
	"github.com/sashaaro/gophkeeper/internal/service"
)

type VaultRepository struct {
	db *Conn
}

var _ service.VaultRepository = &VaultRepository{}

func NewVaultRepository(conn *Conn) *VaultRepository {
	return &VaultRepository{db: conn}
}

func (r VaultRepository) Search(
	ctx context.Context,
	userID uuid.UUID,
	filter service.SecretFilter,
	res []*entity.Secret,
) (n int, err error) {
	limit := cap(res)
	if limit < 1 {
		return 0, ErrEmptyResultSlice
	}
	rows, err := r.db.QueryContext(ctx,
		`SELECT id, user_id, name, kind FROM "secret" WHERE id > $1 AND user_id = $2 ORDER BY name LIMIT $3`,
		filter.After,
		userID,
		limit,
	)
	if err != nil {
		return 0, err
	}
	defer func() {
		if e := rows.Close(); e != nil {
			log.Error("rows.close fails", log.Err(err))
		}
	}()
	for rows.Next() {
		secret := entity.Secret{UserID: userID}
		if err := rows.Scan(&secret.ID, &secret.Name, &secret.Kind); err != nil {
			return 0, err
		}
		res[n] = &secret
		n++
	}
	if err := rows.Err(); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return 0, err
	}
	return n, nil
}

func (r VaultRepository) SaveTags(ctx context.Context, secretID uuid.UUID, tags []*entity.Tag) error {
	return r.db.InTransaction(ctx, func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, `DELETE FROM tag WHERE secret_id = $1`, secretID)
		if err != nil {
			return err
		}
		q := `INSERT INTO tag (id, secret_id, name, value) VALUES ($1, $2, $3, $4)`
		for _, tag := range tags {
			if _, err := tx.ExecContext(ctx, q, tag.ID, secretID, tag.Name, tag.Value); err != nil {
				return err
			}
		}
		return nil
	})
}

func (r VaultRepository) Tags(
	ctx context.Context,
	secretID uuid.UUID,
	filter service.TagFilter,
	res []*entity.Tag,
) (n int, err error) {
	limit := cap(res)
	if limit < 1 {
		return 0, ErrEmptyResultSlice
	}
	rows, err := r.db.QueryContext(ctx,
		`SELECT id, name, value FROM "tag" WHERE id > $1 AND secret_id = $2 ORDER BY id LIMIT $3`,
		filter.After,
		secretID,
		limit,
	)
	if err != nil {
		return 0, err
	}
	defer func() {
		if e := rows.Close(); e != nil {
			log.Error("rows.close fails", log.Err(err))
		}
	}()
	for rows.Next() {
		tag := entity.Tag{SecretID: secretID}
		if err := rows.Scan(&tag.ID, &tag.Name, &tag.Value); err != nil {
			return 0, err
		}
		res[n] = &tag
		n++
	}
	if err := rows.Err(); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return 0, err
	}
	return n, nil
}

func (r VaultRepository) CreateCredentials(ctx context.Context, userID uuid.UUID, m *entity.Credentials) error {
	if m == nil {
		return ErrModelIsNil
	}
	return r.db.InTransaction(ctx, func(ctx context.Context, tx *sql.Tx) error {
		if err := r.insertSecret(ctx, tx, &m.Secret); err != nil {
			return err
		}
		q := `INSERT INTO credentials (id, login, password) VALUES ($1, $2, $3)`
		_, err := r.db.ExecContext(ctx, q, m.ID, m.Login, m.Password)
		return err
	})
}

func (r VaultRepository) insertSecret(ctx context.Context, tx *sql.Tx, m *entity.Secret) error {
	m.ID = uuid.Must(uuid.NewV6())
	q := `INSERT INTO secret (id, user_id, name, kind) VALUES ($1, $2, $3, $4)`
	_, err := tx.ExecContext(ctx, q, m.ID, m.UserID, m.Name, m.Kind)
	return err
}

func (r VaultRepository) GetCredentials(ctx context.Context, userID uuid.UUID, name string, m *entity.Credentials) error {
	if m == nil {
		return ErrModelIsNil
	}
	q := `SELECT s.id, s.user_id, s.name, s.kind, c.login, c.password
		FROM secret s INNER JOIN credentials c ON c.id = s.id
		WHERE s.user_id = $1 AND s.name = $2 AND s.kind = $3`
	row := r.db.QueryRowContext(ctx, q, userID, name, entity.KindCredentials)
	if err := row.Scan(&m.ID, &m.UserID, &m.Name, &m.Kind, &m.Login, &m.Password); err != nil {
		return err
	}
	return nil
}
