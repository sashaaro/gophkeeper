package client

import (
	"context"

	"github.com/sashaaro/gophkeeper/internal/config"
	"github.com/sashaaro/gophkeeper/internal/log"
)

// Client - Клиент к серверу.
// TODO Про grpc теоретически не должен знать. Должно описываться интерфейсами.
type Client struct {
	g        *GRPCClient
	jwtToken string
}

func NewClient(
	cfg *config.Client,
) *Client {
	creds, err := cfg.TLS.Certificate()
	if err != nil {
		log.Panic("Fail to load TLS", log.Err(err))
	}
	return &Client{
		g: NewGRPCClient(cfg.ServerAddr, WithTLS(creds)),
	}
}

func (c *Client) Close() {
	if err := c.g.Close(); err != nil {
		log.Error("Fail close", log.Err(err))
	}
}

func (c *Client) Register(ctx context.Context, login, password string) error {
	tokenString, err := c.g.Register(ctx, login, password)
	if err != nil {
		return err
	}
	log.Info("Attach JWT token")
	c.jwtToken = tokenString
	return c.g.ReInitWithAuth(c.jwtToken)
}

func (c *Client) Login(ctx context.Context, login, password string) error {
	tokenString, err := c.g.Login(ctx, login, password)
	if err != nil {
		return err
	}
	log.Info("Attach JWT token")
	c.jwtToken = tokenString
	return c.g.ReInitWithAuth(c.jwtToken)
}

func (c *Client) Ping(ctx context.Context) error {
	return c.g.Ping(ctx)
}
