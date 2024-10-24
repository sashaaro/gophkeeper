package server

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/sashaaro/gophkeeper/internal/contract"
	"github.com/sashaaro/gophkeeper/internal/log"
)

var (
	errMissingMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
	errInvalidToken    = status.Errorf(codes.Unauthenticated, "invalid token")
)

type GRPCServer struct {
	opts   []grpc.ServerOption
	server *grpc.Server
	addr   string
}

type Opt func(GRPCServer)

func WithTLS(certificate *tls.Certificate) func(*GRPCServer) {
	return func(s *GRPCServer) {
		s.opts = append(s.opts, grpc.Creds(credentials.NewServerTLSFromCert(certificate)))
	}
}

func NewGRPCServer(addr string, opts ...Opt) (*GRPCServer, error) {
	srv := GRPCServer{
		opts: []grpc.ServerOption{
			grpc.UnaryInterceptor(ensureValidToken),
		},
		addr: addr,
	}
	for _, o := range opts {
		o(srv)
	}
	return &srv, nil
}

func (s *GRPCServer) Run() error {
	s.server = grpc.NewServer(s.opts...)
	contract.RegisterAuthServer(s.server, &AuthGRPCServer{})
	contract.RegisterVaultServer(s.server, &VaultGRPCServer{})

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

// valid validates the authorization.
func valid(authorization []string) bool {
	if len(authorization) < 1 {
		return false
	}
	token := strings.TrimPrefix(authorization[0], "Bearer ")
	// Perform the token validation here. For the sake of this example, the code
	// here forgoes any of the usual OAuth2 token validation and instead checks
	// for a token matching an arbitrary string.
	return token == "some-secret-token"
}

func ensureValidToken(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errMissingMetadata
	}
	// The keys within metadata.MD are normalized to lowercase.
	// See: https://godoc.org/google.golang.org/grpc/metadata#New
	if !valid(md["authorization"]) {
		return nil, errInvalidToken
	}
	// Continue execution of handler after ensuring a valid token.
	return handler(ctx, req)
}
