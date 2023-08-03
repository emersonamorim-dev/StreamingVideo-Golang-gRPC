package main

import (
	"context"
	"log"
	"net"

	"streamingvideo/video-upload-service/internal/handlers"
	"streamingvideo/video-upload-service/internal/middleware"

	pb "streamingvideo/video-upload-service/upload.proto"

	"github.com/segmentio/kafka-go"
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

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "video-uploads",
		Balancer: &kafka.LeastBytes{},
	})

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.AuthInterceptor),
	)

	pb.RegisterVideoServiceServer(s, handlers.NewVideoHandler(db, writer))

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
