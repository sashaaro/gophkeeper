package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sashaaro/gophkeeper/internal/service"
)

type VaultRepository struct {
	db *sqlx.DB
}

func (r VaultRepository) GetAll(ctx context.Context, userID uuid.UUID) (map[string][]byte, error) {
	var res []struct {
		Name  string `db:"name"`
		Value []byte `db:"value"`
	}

	err := r.db.SelectContext(ctx, &res, "SELECT name, value FROM secret where user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	result := make(map[string][]byte)
	for _, i := range res {
		result[i.Name] = i.Value
	}

	return result, nil
}

func (r VaultRepository) Save(ctx context.Context, userID uuid.UUID, key string, bytes []byte) error {
	id := uuid.Must(uuid.NewV6())
	q := `INSERT INTO secret (id, user_id, name, value) VALUES ($1, $2, $3, $4)`
	_, err := r.db.ExecContext(ctx, q, id, userID, key, bytes)
	return err
}

var _ service.VaultRepository = &VaultRepository{}

func NewVaultRepository(db *sqlx.DB) *VaultRepository {
	return &VaultRepository{db: db}
}
