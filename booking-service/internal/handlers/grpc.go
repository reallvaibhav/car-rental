package handlers

import (
	"booking-service/internal/services"
	"context"
	pb "booking-service/internal/proto"
)

type Server struct {
	pb.UnimplementedBookingServiceServer
	service services.BookingService
}

func NewBookingServer(svc services.BookingService) *Server {
	return &Server{service: svc}
}

func (s *Server) CreateBooking(ctx context.Context, req *pb.CreateBookingRequest) (*pb.BookingResponse, error) {
	return s.service.CreateBooking(ctx, req)
}

func (s *Server) GetBooking(ctx context.Context, req *pb.GetBookingRequest) (*pb.BookingResponse, error) {
	return s.service.GetBooking(ctx, req)
}

func (s *Server) UpdateBookingStatus(ctx context.Context, req *pb.UpdateBookingStatusRequest) (*pb.BookingResponse, error) {
	return s.service.UpdateBookingStatus(ctx, req)
}

func (s *Server) ListUserBookings(ctx context.Context, req *pb.ListBookingsRequest) (*pb.ListBookingsResponse, error) {
	return s.service.ListUserBookings(ctx, req)
}

func (s *Server) CancelBooking(ctx context.Context, req *pb.CancelBookingRequest) (*pb.BookingResponse, error) {
	return s.service.CancelBooking(ctx, req)
}
