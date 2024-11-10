package service

import (
	"context"

	"github.com/google/uuid"
)

type Paginator struct {
	After uuid.UUID
}

type VaultRepository interface {
	Save(ctx context.Context, userID uuid.UUID, key string, bytes []byte) error
	GetAll(ctx context.Context, userID uuid.UUID) (map[string][]byte, error)
}

type VaultService struct {
	repo VaultRepository
}

func NewVaultService(repo VaultRepository) *VaultService {
	return &VaultService{repo: repo}
}

func (svc *VaultService) Save(ctx context.Context, userID uuid.UUID, key string, bytes []byte) error {
	return svc.repo.Save(ctx, userID, key, bytes)
}

func (svc *VaultService) GetAll(ctx context.Context, userID uuid.UUID) (map[string][]byte, error) {
	return svc.repo.GetAll(ctx, userID)
}
