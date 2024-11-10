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

func (s *KeeperServer) Ping(_ context.Context, in *gophkeeper.Empty) (*gophkeeper.Empty, error) {
	log.Info("ping")
	return in, nil
}

func (s *KeeperServer) GetAll(ctx context.Context, v *gophkeeper.Empty) (*gophkeeper.SecretDataList, error) {
	list, err := s.vaultSvc.GetAll(ctx, uuid.UUID{})

	res := &gophkeeper.SecretDataList{
		Entity: make([]*gophkeeper.SecretData, 0, len(list)),
	}
	for name, value := range list {
		res.Entity = append(res.Entity, &gophkeeper.SecretData{
			Key:   name,
			Value: value,
		})
	}

	return res, err
}
func (s *KeeperServer) SendSecretData(ctx context.Context, v *gophkeeper.SecretData) (*gophkeeper.Empty, error) {
	err := s.vaultSvc.Save(ctx, uuid.UUID{}, v.Key, v.Value)
	return &gophkeeper.Empty{}, err
}
