package main

import (
	"context"
	"log"
	"net"

	"streamingvideo/auth-service/internal/handlers"
	"streamingvideo/auth-service/internal/middleware"
	pb "streamingvideo/auth-service/proto"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func main() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Falha ao conectar-se ao MongoDB: %v", err)
	}

	db := client.Database("test")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Falha ao ouvir: %v", err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(middleware.AuthInterceptor))
	pb.RegisterAuthServiceServer(s, handlers.NewAuthHandler(db))

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Falha no servidor: %v", err)
	}

}

func ChainUnaryServer(interceptors ...grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		for i := len(interceptors) - 1; i >= 0; i-- {
			handler = func(currentHandler grpc.UnaryHandler, currentInterceptor grpc.UnaryServerInterceptor) grpc.UnaryHandler {
				return func(currentCtx context.Context, currentReq interface{}) (interface{}, error) {
					return currentInterceptor(currentCtx, currentReq, info, currentHandler)
				}
			}(handler, interceptors[i])
		}
		return handler(ctx, req)
	}
}
