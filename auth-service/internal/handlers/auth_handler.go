package handlers

import (
	"context"

	"streamingvideo/auth-service/internal/models"
	pb "streamingvideo/auth-service/proto"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthHandler struct {
	db *mongo.Database
}

func NewAuthHandler(db *mongo.Database) *AuthHandler {
	return &AuthHandler{db: db}
}

func (h *AuthHandler) Authenticate(ctx context.Context, req *pb.AuthRequest) (*pb.AuthResponse, error) {
	var user models.User
	err := h.db.Collection("users").FindOne(ctx, bson.M{"username": req.Username}).Decode(&user)
	if err != nil {
		return nil, err
	}

	if user.Password != req.Password {
		return &pb.AuthResponse{Success: false}, nil
	}

	return &pb.AuthResponse{Success: true}, nil
}
