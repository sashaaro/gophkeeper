package client

import (
	"github.com/sashaaro/gophkeeper/internal/config"
)

type Client struct {
	g *GRPCClient
}

func NewClient(
	cfg *config.Client,
) *Client {
	return &Client{
		g: NewGRPCClient(cfg.ServerAddr),
	}
}
