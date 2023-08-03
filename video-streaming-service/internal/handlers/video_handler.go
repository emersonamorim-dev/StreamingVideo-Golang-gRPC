package handlers

import (
	"context"

	"streamingvideo/video-streaming-service/internal/models"

	pb "streamingvideo/video-streaming-service/proto"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type VideoHandler struct {
	db *mongo.Database
}

func NewVideoHandler(db *mongo.Database) *VideoHandler {
	return &VideoHandler{db: db}
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

	return &pb.UploadVideoResponse{Success: true}, nil
}

func (h *VideoHandler) GetVideo(ctx context.Context, req *pb.GetVideoRequest) (*pb.GetVideoResponse, error) {
	var video models.Video
	err := h.db.Collection("videos").FindOne(ctx, bson.M{"id": req.Id}).Decode(&video)
	if err != nil {
		return nil, err
	}

	return &pb.GetVideoResponse{
		Data: video.Data,
	}, nil
}
