package handlers

import (
	"context"

	pb "streamingvideo/kafka-service/proto"

	"github.com/segmentio/kafka-go"
)

type KafkaHandler struct {
	writer *kafka.Writer
}

func NewKafkaHandler(writer *kafka.Writer) *KafkaHandler {
	return &KafkaHandler{writer: writer}
}

func (h *KafkaHandler) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (*pb.SendMessageResponse, error) {
	msg := kafka.Message{
		Value: []byte(req.Message),
	}

	err := h.writer.WriteMessages(ctx, msg)
	if err != nil {
		return nil, err
	}

	return &pb.SendMessageResponse{Success: true}, nil
}
