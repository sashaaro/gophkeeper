package client

import (
	"context"
	"github.com/sashaaro/gophkeeper/internal/config"
	"github.com/sashaaro/gophkeeper/internal/log"
	"github.com/sashaaro/gophkeeper/internal/service"
)

// Client - Клиент к серверу.
// TODO Про grpc теоретически не должен знать. Должно описываться интерфейсами.
type Client struct {
	g          *GRPCClient
	JwtToken   *service.JwtClaims
	jwtService *service.JwtService
}

func NewClient(
	cfg *config.Client,
	jwtService *service.JwtService,
) *Client {
	opts := WithoutTLS()
	if cfg.TLS != nil {
		creds, err := cfg.TLS.Certificate()
		if err != nil {
			log.Panic("Fail to load TLS", log.Err(err))
		}
		opts = WithTLS(creds)
	}

	return &Client{
		g:          NewGRPCClient(cfg.ServerAddr, opts),
		jwtService: jwtService,
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
	c.JwtToken, err = c.jwtService.ParseTokenClaims(tokenString)
	if err != nil {
		return err
	}
	return c.g.ReInitWithAuth(tokenString)
}

func (c *Client) Login(ctx context.Context, login, password string) error {
	tokenString, err := c.g.Login(ctx, login, password)
	if err != nil {
		return err
	}
	log.Info("Attach JWT token")

	c.JwtToken, err = c.jwtService.ParseTokenClaims(tokenString)
	if err != nil {
		return err
	}
	return c.g.ReInitWithAuth(tokenString)
}

func (c *Client) Ping(ctx context.Context) error {
	return c.g.Ping(ctx)
}
