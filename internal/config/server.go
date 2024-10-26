package config

import (
	"crypto/rand"
	"crypto/rsa"

	"github.com/sashaaro/gophkeeper/internal/log"
	"github.com/sashaaro/gophkeeper/internal/service"
	"github.com/sashaaro/gophkeeper/internal/ssl"
)

const DefaultServerListen = ":9876"

type Server struct {
	Listen string
	TLS    ssl.ServerConfig
	JWT    *service.JwtConfig
}

func NewServer() *Server {
	return &Server{
		Listen: getEnv("LISTEN", DefaultServerListen),
		TLS: ssl.ServerConfig{
			PublicKeyPath:  getEnv("TLS_PUBLIC_KEY_PATH", ""),
			PrivateKeyPath: getEnv("TLS_PRIVATE_KEY_PATH", ""),
		},
		JWT: &service.JwtConfig{
			PrivateKeyFile: getEnv("JWT_PRIVATE_KEY_PATH", ""),
			PublicKeyFile:  getEnv("JWT_PUBLIC_KEY_PATH", ""),
		},
	}
}

func (s *Server) GenerateFakeJWT() {
	r, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		log.Panic("Fails to generate RSA", log.Err(err))
	}
	s.JWT.PrivateKey = r
	s.JWT.PublicKey = &r.PublicKey
}
