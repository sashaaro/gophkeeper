package config

import "github.com/sashaaro/gophkeeper/internal/ssl"

const DefaultServerAddress = "127.0.0.1:9876"

type Client struct {
	ServerAddr string
	TLS        ssl.ClientConfig
}

func NewClient() *Client {
	return &Client{
		ServerAddr: getEnv("SERVER_ADDR", DefaultServerAddress),
		TLS: ssl.ClientConfig{
			PublicKeyPath: getEnv("TLS_PUBLIC_KEY_PATH", ""),
		},
	}
}
