package middleware

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// Suponha que o token de autenticação seja "valid-token".
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadados não são fornecidos")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "o token de autorização não foi fornecido")
	}

	if values[0] != "valid-token" {
		return nil, status.Errorf(codes.Unauthenticated, "token de autorização não é válido")
	}

	return handler(ctx, req)
}
