package client

import (
	"context"

	"github.com/sashaaro/gophkeeper/internal/log"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/credentials/oauth"

	"github.com/sashaaro/gophkeeper/pkg/gophkeeper"
)

// GRPCClient - Реализация клиента к серверу по grpc.
// Знает только о grpc контрактах и ни о чём другом. Всё что связано с логикой приложения, слоем выше.
type GRPCClient struct {
	serverAddr    string
	dialOpts      []grpc.DialOption
	authToken     string
	conn          *grpc.ClientConn
	_keeperClient gophkeeper.KeeperServiceClient
	_authClient   gophkeeper.AuthServiceClient
}

type Opt func(*GRPCClient)

func NewGRPCClient(serverAddr string, opts ...Opt) *GRPCClient {
	c := &GRPCClient{serverAddr: serverAddr}
	for _, o := range opts {
		o(c)
	}
	return c
}

func WithTLS(cred credentials.TransportCredentials) Opt {
	return func(c *GRPCClient) {
		// oauth.TokenSource requires the configuration of transport
		// credentials.
		c.dialOpts = append(c.dialOpts, grpc.WithTransportCredentials(cred))
	}
}

func WithoutTLS() Opt {
	return func(c *GRPCClient) {
		c.dialOpts = append(c.dialOpts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
}

func (c *GRPCClient) ReInitWithAuth(token string) error {
	c.authToken = token
	return c.ReInit()
}

func (c *GRPCClient) ReInit() error {
	log.Info("Re init grpc connection")
	if err := c.conn.Close(); err != nil {
		return err
	}
	log.Info("Old grpc connection is closed")
	c.conn = nil
	c._keeperClient = nil
	c._authClient = nil
	return nil
}

func (c *GRPCClient) Ping(ctx context.Context) error {
	a, err := c.keeperClient()
	if err != nil {
		return err
	}
	_, err = a.Ping(ctx, &gophkeeper.Empty{})
	return err
}

func (c *GRPCClient) Register(ctx context.Context, login, password string) (jwtToken string, err error) {
	a, err := c.authClient()
	if err != nil {
		return "", err
	}
	resp, err := a.Register(ctx, &gophkeeper.Credentials{Login: login, Password: password})
	if err != nil {
		return "", err
	}
	return resp.Jwt, nil
}

func (c *GRPCClient) Login(ctx context.Context, login, password string) (tokenString string, err error) {
	a, err := c.authClient()
	if err != nil {
		return "", err
	}
	resp, err := a.Login(ctx, &gophkeeper.Credentials{Login: login, Password: password})
	if err != nil {
		return "", err
	}
	return resp.Jwt, err
}

func (c *GRPCClient) SendSecretData(ctx context.Context, key string, value SecretData) error {
	k, err := c.keeperClient()
	if err != nil {
		return err
	}
	_, err = k.SendSecretData(ctx, &gophkeeper.SecretData{
		Key:   key,
		Value: value,
	})
	return err
}

func (c *GRPCClient) GetAll(ctx context.Context) (*gophkeeper.SecretDataList, error) {
	k, err := c.keeperClient()
	if err != nil {
		return nil, err
	}
	return k.GetAll(ctx, &gophkeeper.Empty{})
}

func (c *GRPCClient) keeperClient() (gophkeeper.KeeperServiceClient, error) {
	if c._keeperClient == nil {
		conn, err := c.connect()
		if err != nil {
			return nil, err
		}
		c._keeperClient = gophkeeper.NewKeeperServiceClient(conn)
	}
	return c._keeperClient, nil
}

func (c *GRPCClient) authClient() (gophkeeper.AuthServiceClient, error) {
	if c._authClient == nil {
		conn, err := c.connect()
		if err != nil {
			return nil, err
		}
		c._authClient = gophkeeper.NewAuthServiceClient(conn)
	}
	return c._authClient, nil
}

func (c *GRPCClient) connect() (grpc.ClientConnInterface, error) {
	if c.conn != nil {
		return c.conn, nil
	}

	opts := c.dialOpts
	if c.authToken != "" {
		opts = append(opts, grpc.WithPerRPCCredentials(
			oauth.TokenSource{TokenSource: oauth2.StaticTokenSource(
				&oauth2.Token{
					AccessToken: c.authToken,
				},
			)},
		))
	}
	conn, err := grpc.NewClient(c.serverAddr, opts...)
	if err != nil {
		return nil, err
	}
	c.conn = conn
	return conn, nil
}

func (c *GRPCClient) Close() error {
	if c == nil || c.conn == nil {
		return nil
	}
	return c.conn.Close()
}
