package auth

import (
	"errors"
	"fmt"

	"github.com/sashaaro/gophkeeper/internal/entity"
	"github.com/sashaaro/gophkeeper/internal/log"
)

//go:generate mockgen -source ./authenticator.go -destination ./authenticator_mock_test.go -package auth_test
type UserGetter interface {
	// GetUser - получить пользователя по login
	GetUser(login string, user *entity.User) error
}

type PasswordHasher interface {
	Hash(pwd string) (string, error)
}

type Authenticator struct {
	hasher PasswordHasher
	repo   UserGetter
}

var ErrWrongCredentials = errors.New("wrong credentials")

func NewAuthenticator(hasher PasswordHasher, repo UserGetter) *Authenticator {
	return &Authenticator{hasher: hasher, repo: repo}
}

// Authenticate - Найти юзера по логину и паролю. Вернёт ошибку ErrWrongCredentials если такой комбинации нет
func (a Authenticator) Authenticate(login, password string) (*entity.User, error) {
	pwdHash, err := a.hasher.Hash(password)
	if err != nil {
		log.Error("Fail to create hash", log.Str("val", password), log.Err(err))
		return nil, fmt.Errorf("authenticate: %w", err)
	}
	var u *entity.User
	if err := a.repo.GetUser(login, u); err != nil {
		log.Error("Fail to get user from repo", log.Err(err), log.Str("login", login))
		return nil, fmt.Errorf("fail to get user: %w", err)
	}
	if u == nil {
		return nil, fmt.Errorf("login not exists: %w", ErrWrongCredentials)
	}
	if pwdHash != u.Password {
		return nil, fmt.Errorf("password mismatch: %w", ErrWrongCredentials)
	}
	return u, nil
}
