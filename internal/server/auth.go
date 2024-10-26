package server

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/sashaaro/gophkeeper/internal/service"
	"github.com/sashaaro/gophkeeper/pkg/gophkeeper"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthServer struct {
	gophkeeper.AuthServiceServer
	userSvc *service.UserService
	jwtSvc  *service.JwtService
}

func NewAuthServer(userSvc *service.UserService, jwtSvc *service.JwtService) *AuthServer {
	return &AuthServer{
		userSvc: userSvc,
		jwtSvc:  jwtSvc,
	}
}

type UserRegisterer interface {
	Register(login, password string) uuid.UUID
}

type UserLoginer interface {
	Login(login, password string) uuid.NullUUID
}

func (s *AuthServer) Login(ctx context.Context, in *gophkeeper.Credentials) (*gophkeeper.AuthToken, error) {
	u, err := s.userSvc.Login(ctx, in.Login, in.Password)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	t, err := s.jwtSvc.CreateToken(*u, service.WithExpiration(time.Hour))
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &gophkeeper.AuthToken{Jwt: t, UserID: u.ID.String()}, nil
}

func (s *AuthServer) Register(ctx context.Context, in *gophkeeper.Credentials) (*gophkeeper.AuthToken, error) {
	u, err := s.userSvc.Create(ctx, in.Login, in.Password)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	t, err := s.jwtSvc.CreateToken(*u, service.WithExpiration(time.Hour))
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &gophkeeper.AuthToken{Jwt: t, UserID: u.ID.String()}, nil
}
