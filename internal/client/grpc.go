package client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/sashaaro/gophkeeper/pkg/gophkeeper"
)

type GRPCClient struct {
	serverAddr    string
	dialOpts      []grpc.DialOption
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

func (c *GRPCClient) Ping(ctx context.Context) error {
	a, err := c.keeperClient()
	if err != nil {
		return err
	}
	_, err = a.Ping(ctx, &gophkeeper.Empty{})
	return err
}

func (c *GRPCClient) Register(ctx context.Context, login, password string) error {
	a, err := c.authClient()
	if err != nil {
		return err
	}
	_, err = a.Register(ctx, &gophkeeper.Credentials{Login: login, Password: password})
	return err
}

func (c *GRPCClient) Login(ctx context.Context, login, password string) error {
	a, err := c.authClient()
	if err != nil {
		return err
	}
	_, err = a.Login(ctx, &gophkeeper.Credentials{Login: login, Password: password})
	return err
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

	conn, err := grpc.NewClient(c.serverAddr, c.dialOpts...)
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
