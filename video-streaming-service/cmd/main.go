package main

import (
	"context"
	"log"
	"net"

	"streamingvideo/video-streaming-service/internal/handlers"
	"streamingvideo/video-streaming-service/internal/middleware"

	pb "streamingvideo/video-streaming-service/proto"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func main() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	db := client.Database("test")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.AuthInterceptor),
	)

	pb.RegisterVideoServiceServer(s, handlers.NewVideoHandler(db))

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
