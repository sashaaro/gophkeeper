package service_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/sashaaro/gophkeeper/internal/entity"
	"github.com/sashaaro/gophkeeper/internal/service"
	"github.com/stretchr/testify/require"
)

func TestJwtService_CreateToken(t *testing.T) {
	rsaPrivateKey, err := GenerateRSAKeys(4096)
	require.NoError(t, err)
	cfg := service.JwtConfig{PrivateKey: rsaPrivateKey, PublicKey: &rsaPrivateKey.PublicKey}
	jwtService, err := service.NewJwtService(&cfg)
	require.NoError(t, err)
	uid := uuid.Must(uuid.NewV6())
	t.Run("Create JWT", func(t *testing.T) {
		tokenString, err := jwtService.CreateToken(entity.User{ID: uid}, time.Second)
		require.NoError(t, err)
		got, err := jwtService.ParseToken(tokenString)
		require.NoError(t, err)
		require.Equal(t, uid, got)
	})
}
