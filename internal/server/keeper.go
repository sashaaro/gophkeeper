package server

import (
	"context"

	"github.com/google/uuid"
	"github.com/sashaaro/gophkeeper/internal/log"
	gophkeeper2 "github.com/sashaaro/gophkeeper/pkg/gophkeeper"
)

type KeeperServer struct {
	gophkeeper2.KeeperServer
}

type UserRegisterer interface {
	Register(login, password string) uuid.UUID
}

type UserLoginer interface {
	Login(login, password string) uuid.NullUUID
}

type CredentialsSaver interface {
	SaveCredentials(userID uuid.UUID)
}

func NewKeeperServer() *KeeperServer {
	return &KeeperServer{}
}

func (s *KeeperServer) Ping(_ context.Context, in *gophkeeper2.Empty) (*gophkeeper2.Empty, error) {
	log.Info("ping")
	return in, nil
}
