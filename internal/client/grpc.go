package client

import (
	"context"

	"github.com/sashaaro/gophkeeper/internal/contract"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClient struct {
	serverAddr string
	dialOpts   []grpc.DialOption
	conn       *grpc.ClientConn
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
	conn, err := c.connect()
	if err != nil {
		return err
	}
	a := contract.NewKeeperClient(conn)
	_, err = a.Ping(ctx, &contract.Empty{})
	if err != nil {
		return err
	}
	return nil
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

func (c *GRPCClient) Stop() error {
	return c.conn.Close()
}
