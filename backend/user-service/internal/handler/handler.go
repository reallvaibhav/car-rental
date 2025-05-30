package handler

import (
	"context"

	pb_user "github.com/Car-Rental/proto/user"

	"github.com/Car-Rental/backend/user-service/internal/service"
)

type UserServer struct {
	pb_user.UnimplementedUserServiceServer
	service *service.UserService
}

func New(service *service.UserService) *UserServer {
	return &UserServer{service: service}
}

func (s *UserServer) Register(ctx context.Context, req *pb_user.RegisterRequest) (*pb_user.AuthResponse, error) {
	token, err := s.service.Register(req.Email, req.Password, req.Name)
	if err != nil {
		return nil, err
	}
	return &pb_user.AuthResponse{Token: token}, nil
}

func (s *UserServer) Login(ctx context.Context, req *pb_user.LoginRequest) (*pb_user.AuthResponse, error) {
	token, err := s.service.Login(req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	return &pb_user.AuthResponse{Token: token}, nil
}

func (s *UserServer) GetUserByID(ctx context.Context, req *pb_user.UserIDRequest) (*pb_user.UserResponse, error) {
	user, err := s.service.GetUserByID(req.UserId)
	if err != nil {
		return nil, err
	}
	return &pb_user.UserResponse{
		UserId:    user.ID.Hex(),
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt.String(),
	}, nil
}

func (s *UserServer) ValidateToken(ctx context.Context, req *pb_user.TokenRequest) (*pb_user.ValidateResponse, error) {
	userID, err := s.service.ValidateToken(req.Token)
	if err != nil {
		return &pb_user.ValidateResponse{Valid: false}, nil
	}
	return &pb_user.ValidateResponse{
		Valid:  true,
		UserId: userID,
	}, nil
}
