package server

import (
	"context"

	"github.com/google/uuid"
	"github.com/sashaaro/gophkeeper/internal/log"
	"github.com/sashaaro/gophkeeper/internal/service"
	"github.com/sashaaro/gophkeeper/pkg/gophkeeper"
)

type KeeperServer struct {
	gophkeeper.UnimplementedKeeperServiceServer
	userSvc  *service.UserService
	vaultSvc *service.VaultService
}

func NewKeeperServer(userSvc *service.UserService, vaultSvc *service.VaultService) *KeeperServer {
	return &KeeperServer{
		userSvc:  userSvc,
		vaultSvc: vaultSvc,
	}
}

type CredentialsSaver interface {
	SaveCredentials(userID uuid.UUID)
}

func (s *KeeperServer) Ping(_ context.Context, in *gophkeeper.Empty) (*gophkeeper.Empty, error) {
	log.Info("ping")
	return in, nil
}
