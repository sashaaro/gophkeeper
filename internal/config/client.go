package config

import (
	"github.com/sashaaro/gophkeeper/internal/service"
	"github.com/sashaaro/gophkeeper/internal/ssl"
)

const DefaultServerAddress = "127.0.0.1:9876"

type Client struct {
	ServerAddr string
	TLS        *ssl.ClientConfig
	JWT        *service.JwtConfig
}

func NewClient() *Client {
	tls := getEnv("TLS_PUBLIC_KEY_PATH", "")
	client := &Client{
		ServerAddr: getEnv("SERVER_ADDR", DefaultServerAddress),
		JWT: &service.JwtConfig{
			PrivateKeyFile: getEnv("JWT_PRIVATE_KEY_PATH", ""),
			PublicKeyFile:  getEnv("JWT_PUBLIC_KEY_PATH", ""),
		},
	}
	if tls != "" {
		client.TLS = &ssl.ClientConfig{PublicKeyPath: tls}
	}

	return client
}
