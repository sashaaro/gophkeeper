package config

import "github.com/sashaaro/gophkeeper/internal/ssl"

const DefaultServerListen = ":9876"

type Server struct {
	Listen string
	TLS    ssl.ServerConfig
}

func NewServer() *Server {
	return &Server{
		Listen: getEnv("LISTEN", DefaultServerListen),
		TLS: ssl.ServerConfig{
			PublicKeyPath:  getEnv("TLS_PUBLIC_KEY_PATH", ""),
			PrivateKeyPath: getEnv("TLS_PRIVATE_KEY_PATH", ""),
		},
	}
}
