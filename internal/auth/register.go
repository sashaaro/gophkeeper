package auth

import (
	"errors"
	"fmt"

	"github.com/sashaaro/gophkeeper/internal/entity"
	"github.com/sashaaro/gophkeeper/internal/log"
)

//go:generate mockgen -source ./register.go -destination ./register_mock_test.go  -package auth_test
type UserCreator interface {
	// CreateUser - создаст пользователя, заполнит поле ID у сущности
	CreateUser(user *entity.User) error
}

type Registerer struct {
	hasher PasswordHasher
	repo   UserCreator
}

var ErrAlreadyExists = errors.New("already exists")

func NewRegisterer(hasher PasswordHasher, repo UserCreator) *Registerer {
	return &Registerer{hasher: hasher, repo: repo}
}

func (r *Registerer) Register(login, password string) (*entity.User, error) {
	pwdHash, err := r.hasher.Hash(password)
	if err != nil {
		log.Error("Fail to create hash", log.Str("val", password), log.Err(err))
		return nil, fmt.Errorf("register user: %w", err)
	}
	u := entity.User{
		Login:    login,
		Password: pwdHash,
	}
	if err := r.repo.CreateUser(&u); err != nil {
		return nil, fmt.Errorf("register user fails: %w", err)
	}
	return &u, nil
}
