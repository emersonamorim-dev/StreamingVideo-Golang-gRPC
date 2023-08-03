package main

import (
	"log"
	"net"

	"streamingvideo/kafka-service/internal/handlers"

	pb "streamingvideo/kafka-service/proto"

	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc"
)

func main() {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "auth",
		Balancer: &kafka.LeastBytes{},
	})

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterKafkaServiceServer(s, handlers.NewKafkaHandler(writer))

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
