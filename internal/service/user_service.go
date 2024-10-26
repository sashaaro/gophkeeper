package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/sashaaro/gophkeeper/internal/entity"
	"github.com/sashaaro/gophkeeper/internal/log"
)

//go:generate mockgen -source ./user_service.go -destination ./mocks/user_service.go -package mocks
type PasswordHasher interface {
	Hash(pwd string) string
}

type UserCreator interface {
	Create(ctx context.Context, m *entity.User) error
}

type UserGetter interface {
	Get(ctx context.Context, id uuid.UUID, m *entity.User) error
	GetByLogin(ctx context.Context, login string, m *entity.User) error
}

//	type RepositorySearcher interface {
//		Search(ctx context.Context, filter ???, res []*entity.User) error
//	}

type UserRepository interface {
	UserCreator
	UserGetter
}

type UserService struct {
	hasher PasswordHasher
	repo   UserRepository
}

var ErrWrongCredentials = errors.New("wrong credentials")

func NewUserService(hasher PasswordHasher, repo UserRepository) *UserService {
	return &UserService{repo: repo, hasher: hasher}
}

func (s *UserService) Login(ctx context.Context, login, password string) (*entity.User, error) {
	var u *entity.User
	if err := s.repo.GetByLogin(ctx, login, u); err != nil {
		log.Error("Fail to get user from repo", log.Err(err), log.Str("login", login))
		return nil, fmt.Errorf("fail to get user: %w", err)
	}
	if u == nil {
		return nil, fmt.Errorf("login not exists: %w", ErrWrongCredentials)
	}
	if s.hasher.Hash(password) != u.Password {
		return nil, fmt.Errorf("password mismatch: %w", ErrWrongCredentials)
	}
	return u, nil
}

// Get - получить юзера по id из репозитория
func (s *UserService) Get(ctx context.Context, id uuid.UUID, m *entity.User) error {
	return s.repo.Get(ctx, id, m)
}

// Create - создать юзера. У модели entity.User генерится новый ID и она сохраняется в репозитории
func (s *UserService) Create(ctx context.Context, login, password string) (_ *entity.User, err error) {
	m := entity.User{
		Login:    login,
		Password: s.hasher.Hash(password),
	}
	if m.ID, err = uuid.NewV6(); err != nil {
		return nil, err
	}
	if err := s.repo.Create(ctx, &m); err != nil {
		return nil, fmt.Errorf("register user fails: %w", err)
	}
	return &m, nil
}
