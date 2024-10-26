package postgres_test

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/sashaaro/gophkeeper/internal/entity"
	"github.com/sashaaro/gophkeeper/internal/postgres"
)

func (s *PostgresSuite) TestUserRepository_Create() {
	r := postgres.NewUserRepository(s.conn)
	s.Run("Create and get a user", func() {
		m := entity.User{
			Login:    "test" + time.Now().Format("2006-01-02T15-04-05.999999999"),
			Password: "hash",
		}
		s.Require().NoError(r.Create(context.Background(), &m))
		s.Require().NotEqual(uuid.Nil, m.ID)
		var got entity.User
		s.Require().NoError(r.Get(context.Background(), m.ID, &got))
		s.Require().Equal(m, got)
	})
	s.Run("Fail create a user when model is nil", func() {
		var m *entity.User
		s.Require().Error(r.Create(context.Background(), m))
	})
}
