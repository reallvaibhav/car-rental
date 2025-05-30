package delivery

import (
	"context"
	"statistics-service/internal/usecase"
	"statistics-service/statistics-service/proto"
)

type Server struct {
	proto.UnimplementedStatisticsServiceServer
	Usecase usecase.StatisticsUsecase
}

func (s *Server) GetUserOrderStatistics(ctx context.Context, req *proto.UserOrderStatisticsRequest) (*proto.UserOrderStatisticsResponse, error) {
	totalOrders, mostActiveTime, err := s.Usecase.GetUserOrderStatistics(req.UserId)
	if err != nil {
		return nil, err
	}
	return &proto.UserOrderStatisticsResponse{
		TotalOrders:    totalOrders,
		MostActiveTime: mostActiveTime,
	}, nil
}

func (s *Server) GetUserStatistics(ctx context.Context, req *proto.UserStatisticsRequest) (*proto.UserStatisticsResponse, error) {
	totalUsers, activeUsers, err := s.Usecase.GetUserStatistics()
	if err != nil {
		return nil, err
	}
	return &proto.UserStatisticsResponse{
		TotalUsers:  totalUsers,
		ActiveUsers: activeUsers,
	}, nil
}
