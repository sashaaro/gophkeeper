package server

import (
	"context"

	"github.com/google/uuid"
	"github.com/sashaaro/gophkeeper/internal/log"
	"github.com/sashaaro/gophkeeper/internal/service"
	"github.com/sashaaro/gophkeeper/pkg/gophkeeper"
)

type KeeperServer struct {
	gophkeeper.KeeperServiceServer
	userSvc *service.UserService
}

func NewKeeperServer(userSvc *service.UserService) *KeeperServer {
	return &KeeperServer{
		userSvc: userSvc,
	}
}

type CredentialsSaver interface {
	SaveCredentials(userID uuid.UUID)
}

func (s *KeeperServer) Ping(_ context.Context, in *gophkeeper.Empty) (*gophkeeper.Empty, error) {
	log.Info("ping")
	return in, nil
}
