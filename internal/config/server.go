package config

import "github.com/sashaaro/gophkeeper/internal/ssl"

const DefaultServerListen = ":9876"

type Server struct {
	Listen string
	TLS    ssl.Config
}

func NewServer() *Server {
	return &Server{
		Listen: DefaultServerListen,
		TLS:    ssl.Config{},
	}
}
