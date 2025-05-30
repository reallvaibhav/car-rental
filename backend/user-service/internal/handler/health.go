package handler

import (
	"context"
	"time"

	pb_health "google.golang.org/grpc/health/grpc_health_v1"
)

type HealthServer struct {
	pb_health.UnimplementedHealthServer
}

func NewHealthServer() *HealthServer {
	return &HealthServer{}
}

func (s *HealthServer) Check(ctx context.Context, req *pb_health.HealthCheckRequest) (*pb_health.HealthCheckResponse, error) {
	return &pb_health.HealthCheckResponse{
		Status: pb_health.HealthCheckResponse_SERVING,
	}, nil
}

func (s *HealthServer) Watch(req *pb_health.HealthCheckRequest, stream pb_health.Health_WatchServer) error {
	for {
		select {
		case <-stream.Context().Done():
			return nil
		default:
			if err := stream.Send(&pb_health.HealthCheckResponse{
				Status: pb_health.HealthCheckResponse_SERVING,
			}); err != nil {
				return err
			}
			time.Sleep(time.Second)
		}
	}
}
