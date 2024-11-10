package client

import (
	"context"
	"github.com/sashaaro/gophkeeper/internal/config"
	"github.com/sashaaro/gophkeeper/internal/entity"
	"github.com/sashaaro/gophkeeper/internal/log"
)

// Client - Клиент к серверу.
// TODO Про grpc теоретически не должен знать. Должно описываться интерфейсами.
type Client struct {
	g         *GRPCClient
	LoginName string
}

func NewClient(
	cfg *config.Client,
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
		g: NewGRPCClient(cfg.ServerAddr, opts),
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
	c.LoginName = login
	return c.g.ReInitWithAuth(tokenString)
}

func (c *Client) Login(ctx context.Context, login, password string) error {
	tokenString, err := c.g.Login(ctx, login, password)
	if err != nil {
		return err
	}
	log.Info("Attach JWT token")
	c.LoginName = login
	return c.g.ReInitWithAuth(tokenString)
}

func (c *Client) Logout() {
	c.LoginName = ""
	c.g.ReInitWithAuth("")
}

func (c *Client) Ping(ctx context.Context) error {
	return c.g.Ping(ctx)
}

func (c *Client) SendSecretBinary(name string, value []byte) error {
	return c.g.SendSecretData(context.Background(), name, BytesToSecretData(value))
}

func (c *Client) SendSecretText(name string, value string) error {
	return c.g.SendSecretData(context.Background(), name, StringToText(value))
}

func (c *Client) SendSecretCreditCard(name string, value entity.CreditCard) error {
	return c.g.SendSecretData(context.Background(), name, CreditCardToSecretData(value))
}

func (c *Client) GetAll() (map[string][]byte, error) {
	list, err := c.g.GetAll(context.Background())
	if err != nil {
		return nil, err
	}
	res := make(map[string][]byte, len(list.Entity))
	for _, data := range list.Entity {
		res[data.Key] = data.Value
	}

	return res, nil
}
