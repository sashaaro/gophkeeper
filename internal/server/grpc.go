package server

import (
	"crypto/tls"
	"fmt"
	"net"

	"github.com/sashaaro/gophkeeper/internal/auth"
	"github.com/sashaaro/gophkeeper/internal/log"
	"github.com/sashaaro/gophkeeper/internal/service"
	"github.com/sashaaro/gophkeeper/pkg/gophkeeper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCServer struct {
	opts   []grpc.ServerOption
	server *grpc.Server
	addr   string
	auth   auth.Authenticator
}

type Opt func(*GRPCServer)

func WithTLS(certificate *tls.Certificate) Opt {
	return func(s *GRPCServer) {
		s.opts = append(s.opts, grpc.Creds(credentials.NewServerTLSFromCert(certificate)))
	}
}

func WithoutTLS() Opt {
	return func(s *GRPCServer) {
		s.opts = append(s.opts, grpc.Creds(insecure.NewCredentials()))
	}
}

func WithAuth(authenticator *auth.Authenticator) Opt {
	return func(s *GRPCServer) {
		s.opts = append(s.opts, grpc.UnaryInterceptor(authenticator.AuthorizeInterceptor()))
	}
}

func NewGRPCServer(
	addr string,
	userSvc *service.UserService,
	jwtSvc *service.JwtService,
	opts ...Opt,
) *GRPCServer {
	srv := GRPCServer{
		opts: []grpc.ServerOption{},
		addr: addr,
	}
	for _, o := range opts {
		o(&srv)
	}
	srv.server = grpc.NewServer(srv.opts...)
	gophkeeper.RegisterAuthServiceServer(srv.server, NewAuthServer(userSvc, jwtSvc))
	gophkeeper.RegisterKeeperServiceServer(srv.server, NewKeeperServer(userSvc))
	return &srv
}

func (s *GRPCServer) Serve() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("failed to listen GRPC on %s", s.addr)
	}

	log.Info("Serve GRPC", log.Str("addr", s.addr))
	return s.server.Serve(lis)
}

func (s *GRPCServer) Shutdown() {
	s.server.GracefulStop()
}
