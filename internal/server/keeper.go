package server

import (
	"context"

	"github.com/google/uuid"
	"github.com/sashaaro/gophkeeper/internal/contract"
	"github.com/sashaaro/gophkeeper/internal/log"
)

type KeeperServer struct {
	contract.KeeperServer
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

func (s *KeeperServer) Ping(_ context.Context, in *contract.Empty) (*contract.Empty, error) {
	log.Info("ping")
	return in, nil
}
