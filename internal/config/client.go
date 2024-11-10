package config

import (
	"github.com/sashaaro/gophkeeper/internal/ssl"
)

const DefaultServerAddress = "127.0.0.1:9876"

type Client struct {
	ServerAddr string
	TLS        *ssl.ClientConfig
}

func NewClient() *Client {
	tls := getEnv("TLS_PUBLIC_KEY_PATH", "")
	client := &Client{
		ServerAddr: getEnv("SERVER_ADDR", DefaultServerAddress),
	}
	if tls != "" {
		client.TLS = &ssl.ClientConfig{PublicKeyPath: tls}
	}

	return client
}
