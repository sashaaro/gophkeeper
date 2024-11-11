package postgres_test

import (
	"context"
	"github.com/jmoiron/sqlx"
	"os"
	"testing"

	"github.com/sashaaro/gophkeeper/internal/log"
	"github.com/stretchr/testify/suite"

	"github.com/sashaaro/gophkeeper/internal/postgres"
)

type PostgresSuite struct {
	suite.Suite
	db *sqlx.DB
}

func (s *PostgresSuite) SetupTest() {
	dsn := os.Getenv("TEST_DSN")
	if dsn == "" {
		s.T().Skip("Unable to connect with DB. Env TEST_DSN is not specified")
		return
	}
	db, err := postgres.NewConn(dsn)
	s.Require().NoError(err)
	log.Info(`Truncate table "user"`)
	_, err = db.ExecContext(context.Background(), `TRUNCATE TABLE "user" CASCADE`)
	s.Require().NoError(err)
	s.db = db
}

func (s *PostgresSuite) TearDownTest() {
	s.Assert().Nil(s.db.Close())
}

func TestDBStorage(t *testing.T) {
	suite.Run(t, new(PostgresSuite))
}

func (s *PostgresSuite) TestPing() {
	s.Require().NoError(s.db.Ping())
}
