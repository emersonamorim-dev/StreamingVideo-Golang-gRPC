package handlers

import (
	"context"

	"streamingvideo/video-upload-service/internal/models"

	pb "streamingvideo/video-upload-service/proto"

	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
)

type VideoHandler struct {
	db     *mongo.Database
	writer *kafka.Writer
}

func NewVideoHandler(db *mongo.Database, writer *kafka.Writer) *VideoHandler {
	return &VideoHandler{db: db, writer: writer}
}

func (h *VideoHandler) UploadVideo(ctx context.Context, req *pb.UploadVideoRequest) (*pb.UploadVideoResponse, error) {
	video := models.Video{
		ID:   req.Id,
		Data: req.Data,
	}

	_, err := h.db.Collection("videos").InsertOne(ctx, video)
	if err != nil {
		return nil, err
	}

	msg := kafka.Message{
		Value: []byte("Video uploaded: " + req.Id),
	}

	err = h.writer.WriteMessages(ctx, msg)
	if err != nil {
		return nil, err
	}

	return &pb.UploadVideoResponse{Success: true}, nil
}
