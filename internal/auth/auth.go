package auth

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/sashaaro/gophkeeper/internal/log"
)

const AuthHeader = "authorization"

type JwtParser interface {
	ParseToken(tokenString string) (uuid.UUID, error)
}

type Authenticator struct {
	jwtParser JwtParser
}

func NewAuthenticator(jwtParser JwtParser) *Authenticator {
	return &Authenticator{jwtParser: jwtParser}
}

func (a *Authenticator) AuthorizeInterceptor() grpc.UnaryServerInterceptor {
	return a.ensureValidToken
}

var (
	ErrMissingMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
	ErrInvalidToken    = status.Errorf(codes.Unauthenticated, "invalid token")
)

func (a *Authenticator) ensureValidToken(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	mn, ok := grpc.Method(ctx)
	if !ok {
		return nil, ErrMissingMetadata
	}
	log.Info("Call grpc method", log.Str("mn", mn))
	if strings.Index(mn, "/gophkeeper.v1.AuthService/") == 0 {
		// Пропускаем аутентификацию для сервиса авторизации. В будущем надо заменить на grpc.authz с RBAC
		// Некогда разбираться, сроки горят
		return handler(ctx, req)
	}
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, ErrMissingMetadata
	}
	authorization := md[AuthHeader]
	if len(authorization) < 1 {
		log.Info("No token provided")
		return nil, ErrInvalidToken
	}
	token := strings.TrimPrefix(authorization[0], "Bearer ")
	uid, err := a.jwtParser.ParseToken(token)
	if err != nil {
		log.Info("wrong token provided", log.Err(err))
		return nil, ErrInvalidToken
	}
	return handler(context.WithValue(ctx, userIDKey, uid), req)
}

var userIDKey struct{}

// ExtractUserIDFromContext - вытащить userID из контекста. Вернёт nil, если нет
func ExtractUserIDFromContext(ctx context.Context) *uuid.UUID {
	if v, ok := ctx.Value(userIDKey).(uuid.UUID); ok {
		return &v
	}
	return nil
}
