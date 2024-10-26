package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/sashaaro/gophkeeper/internal/entity"
)

type CredentialsCreator interface {
	CreateCredentials(ctx context.Context, userID uuid.UUID, m *entity.Credentials) error
}

type CredentialsGetter interface {
	GetCredentials(ctx context.Context, userID uuid.UUID, name string, m *entity.Credentials) error
}

type Paginator struct {
	After uuid.UUID
}

type SecretFilter struct {
	Paginator
}

type TagFilter struct {
	Paginator
}

type SecretFinder interface {
	// Search - вытаскивает секреты заданного юзера отсортированные по имени в количестве не превышающем вместимость слайса res
	Search(ctx context.Context, userID uuid.UUID, filter SecretFilter, res []*entity.Secret) (n int, err error)
}

type TagsGetter interface {
	Tags(ctx context.Context, secretID uuid.UUID, filter TagFilter, tags []*entity.Tag) (n int, err error)
}

type TagsCreator interface {
	SaveTags(ctx context.Context, secretID uuid.UUID, tags []*entity.Tag) error
}

type VaultRepository interface {
	SecretFinder
	TagsCreator
	TagsGetter
	CredentialsCreator
	CredentialsGetter
}

type VaultService struct {
	repo VaultRepository
}

func NewVaultService(repo VaultRepository) *VaultService {
	return &VaultService{repo: repo}
}

func (svc *VaultService) CreateCredentialsCreateCredentials(ctx context.Context, userID uuid.UUID, m *entity.Credentials) error {
	return svc.repo.CreateCredentials(ctx, userID, m)
}
