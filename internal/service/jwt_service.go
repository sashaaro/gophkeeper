package service

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/sashaaro/gophkeeper/internal/entity"
	"github.com/sashaaro/gophkeeper/internal/log"
)

type JwtClaims struct {
	jwt.RegisteredClaims
}

type JwtConfig struct {
	PublicKey      *rsa.PublicKey
	PublicKeyFile  string
	PrivateKey     *rsa.PrivateKey
	PrivateKeyFile string
	Leeway         time.Duration
}

func (cfg *JwtConfig) LoadKeys() error {
	if cfg.PublicKeyFile != "" {
		keyData, err := os.ReadFile(cfg.PublicKeyFile)
		if err != nil {
			return err
		}
		cfg.PublicKey, err = jwt.ParseRSAPublicKeyFromPEM(keyData)
		if err != nil {
			return err
		}
	}
	if cfg.PrivateKeyFile != "" {
		keyData, err := os.ReadFile(cfg.PrivateKeyFile)
		if err != nil {
			return err
		}
		cfg.PrivateKey, err = jwt.ParseRSAPrivateKeyFromPEM(keyData)
		if err != nil {
			return err
		}
	}
	return nil
}

func (cfg *JwtConfig) Export() (privateKeyPEM []byte, publicKeyPEM []byte, err error) {
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(cfg.PrivateKey)
	privateKeyPEM = pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	})
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(cfg.PublicKey)
	if err != nil {
		return nil, nil, err
	}
	publicKeyPEM = pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	})
	return
}

type JwtService struct {
	cfg *JwtConfig
}

func NewJwtService(cfg *JwtConfig) (*JwtService, error) {
	if err := cfg.LoadKeys(); err != nil {
		return nil, err
	}
	return &JwtService{cfg: cfg}, nil
}

func (s *JwtService) CreateToken(u entity.User, duration time.Duration) (string, error) {
	claims := JwtClaims{}
	claims.ID = u.ID.String()
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(duration))
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)
	tokenString, err := token.SignedString(s.cfg.PrivateKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s *JwtService) ParseToken(tokenString string) (uuid.UUID, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return s.cfg.PublicKey, nil
	}, jwt.WithLeeway(s.cfg.Leeway))
	if err != nil {
		return uuid.Nil, err
	}
	claims, ok := token.Claims.(*JwtClaims)
	if !ok || !token.Valid || claims.ID == "" {
		log.Info("wrong jwt", log.Str("ID", claims.ID))
		return uuid.Nil, errors.New("wrong structure jwt")
	}
	return uuid.Parse(claims.ID)
}
