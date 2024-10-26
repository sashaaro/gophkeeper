package postgres_test

import (
	"context"

	"github.com/google/uuid"
	"github.com/sashaaro/gophkeeper/internal/entity"
	"github.com/sashaaro/gophkeeper/internal/postgres"
)

func (s *PostgresSuite) TestVaultRepository_CreateCredentials() {
	r := postgres.NewVaultRepository(s.conn)
	m := entity.NewCredentials("lilu creds", "lilu", "tratata")
	s.Require().NoError(r.CreateCredentials(context.Background(), TestUserID, m))
	s.Require().NotEqual(uuid.Nil, m.ID)
}
