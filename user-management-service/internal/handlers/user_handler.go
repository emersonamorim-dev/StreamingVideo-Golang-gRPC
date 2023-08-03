package handlers

import (
	"context"

	"streamingvideo/user-management-service/internal/models"
	pb "streamingvideo/user-management-service/proto"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserHandler struct {
	db *mongo.Database
}

func NewUserHandler(db *mongo.Database) *UserHandler {
	return &UserHandler{db: db}
}

func (h *UserHandler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user := models.User{
		Username: req.Username,
		Password: req.Password,
	}

	_, err := h.db.Collection("users").InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{Success: true}, nil
}

func (h *UserHandler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	var user models.User
	err := h.db.Collection("users").FindOne(ctx, bson.M{"username": req.Username}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &pb.GetUserResponse{
		Username: user.Username,
		Password: user.Password,
	}, nil
}
