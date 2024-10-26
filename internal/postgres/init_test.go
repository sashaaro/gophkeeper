package postgres_test

import (
	"context"
	"database/sql"
	"embed"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/sashaaro/gophkeeper/internal/log"
	"github.com/stretchr/testify/suite"

	"github.com/sashaaro/gophkeeper/internal/postgres"
)

type PostgresSuite struct {
	suite.Suite
	conn *postgres.Conn
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
	s.conn = db
	s.Require().NoError(db.InTransaction(context.Background(), loadFixtures))
}

func (s *PostgresSuite) TearDownTest() {
	s.Assert().Nil(s.conn.Close())
}

func TestDBStorage(t *testing.T) {
	suite.Run(t, new(PostgresSuite))
}

func (s *PostgresSuite) TestPing() {
	s.Require().NoError(s.conn.Ping(context.Background()))
}

var TestUserID = uuid.Must(uuid.Parse("01ef6697-3190-6984-9572-74563c32efde"))
var TestUserLogin = "test"
var TestUserHash = "123"

//go:embed fixtures_test.sql
var fixtures embed.FS

func loadFixtures(ctx context.Context, tx *sql.Tx) (err error) {
	q, err := fixtures.ReadFile("fixtures_test.sql")
	if err != nil {
		return err
	}
	_, err = tx.ExecContext(ctx, string(q))
	return err
}
