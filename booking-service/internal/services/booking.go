package services

import (
	pb "booking-service/internal/proto"
	"context"
	"fmt"
	"log"
	"time"

	"booking-service/cache"
	"booking-service/internal/models"
	"booking-service/internal/repository"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BookingService interface {
	CreateBooking(ctx context.Context, req *pb.CreateBookingRequest) (*pb.BookingResponse, error)
	GetBooking(ctx context.Context, req *pb.GetBookingRequest) (*pb.BookingResponse, error)
	UpdateBookingStatus(ctx context.Context, req *pb.UpdateBookingStatusRequest) (*pb.BookingResponse, error)
	ListUserBookings(ctx context.Context, req *pb.ListBookingsRequest) (*pb.ListBookingsResponse, error)
	CancelBooking(ctx context.Context, req *pb.CancelBookingRequest) (*pb.BookingResponse, error)
}

type bookingService struct {
	repo       *repository.Repository
	cache      *cache.Cache
	redisCache *cache.RedisCache
}

func NewBookingService(repo *repository.Repository, cache *cache.Cache, redisCache *cache.RedisCache) BookingService {
	return &bookingService{
		repo:       repo,
		cache:      cache,
		redisCache: redisCache,
	}
}

func (s *bookingService) CreateBooking(ctx context.Context, req *pb.CreateBookingRequest) (*pb.BookingResponse, error) {
	if req.UserId == "" || len(req.Bookings) == 0 {
		return nil, status.Error(codes.InvalidArgument, "user ID and bookings must be provided")
	}

	result, err := s.repo.CreateBooking(ctx, req)
	if err != nil {
		return nil, err
	}

	// Cache the new booking if creation was successful
	if result != nil && result.BookingId != "" {
		cacheKey := fmt.Sprintf("booking:%s", result.BookingId)
		// Convert proto to model for caching
		bookingModel := &models.BookingModel{
			ID:         result.BookingId,
			UserID:     result.UserId,
			Status:     result.Status,
			TotalPrice: result.TotalPrice,
		}
		// Convert booking items
		for _, item := range result.Bookings {
			bookingModel.Bookings = append(bookingModel.Bookings, models.CarBookingItemModelFromProto(item))
		}

		s.cache.Set(cacheKey, bookingModel, 2*time.Hour)
		if s.redisCache != nil {
			s.redisCache.Set(ctx, cacheKey, bookingModel, 2*time.Hour)
		}
		log.Printf("Cached new booking: %s", result.BookingId)
	}

	return result, nil
}

func (s *bookingService) GetBooking(ctx context.Context, req *pb.GetBookingRequest) (*pb.BookingResponse, error) {
	if req.BookingId == "" {
		return nil, status.Error(codes.InvalidArgument, "booking ID must be provided")
	}

	// Try to get from cache first
	cacheKey := fmt.Sprintf("booking:%s", req.BookingId)

	// Try Redis cache first if available
	if s.redisCache != nil {
		var cachedBooking models.BookingModel
		if err := s.redisCache.Get(ctx, cacheKey, &cachedBooking); err == nil {
			log.Printf("Cache hit (Redis) for booking: %s", req.BookingId)
			return models.BookingModelToProto(&cachedBooking), nil
		}
	}

	// Try in-memory cache
	if cached, found := s.cache.Get(cacheKey); found {
		if booking, ok := cached.(*models.BookingModel); ok {
			log.Printf("Cache hit (memory) for booking: %s", req.BookingId)
			return models.BookingModelToProto(booking), nil
		}
	}

	// If not in cache, get from database
	booking, err := s.repo.GetBookingByID(ctx, req.BookingId)
	if err != nil {
		log.Printf("Error getting booking: %v", err)
		return nil, status.Errorf(codes.NotFound, "booking not found: %v", err)
	}

	// Cache the result
	s.cache.Set(cacheKey, booking, 2*time.Hour)
	if s.redisCache != nil {
		s.redisCache.Set(ctx, cacheKey, booking, 2*time.Hour)
	}

	return models.BookingModelToProto(booking), nil
}

func (s *bookingService) UpdateBookingStatus(ctx context.Context, req *pb.UpdateBookingStatusRequest) (*pb.BookingResponse, error) {
	if req.BookingId == "" || req.Status == "" {
		return nil, status.Error(codes.InvalidArgument, "booking ID and status must be provided")
	}

	result, err := s.repo.UpdateBookingStatus(ctx, req)
	if err != nil {
		return nil, err
	}

	// Invalidate cache for this booking
	cacheKey := fmt.Sprintf("booking:%s", req.BookingId)
	s.cache.Delete(cacheKey)
	if s.redisCache != nil {
		s.redisCache.Delete(ctx, cacheKey)
	}
	log.Printf("Invalidated cache for updated booking: %s", req.BookingId)

	// Update cache with new data if operation was successful
	if result != nil {
		bookingModel := &models.BookingModel{
			ID:         result.BookingId,
			UserID:     result.UserId,
			Status:     result.Status,
			TotalPrice: result.TotalPrice,
		}
		// Convert booking items
		for _, item := range result.Bookings {
			bookingModel.Bookings = append(bookingModel.Bookings, models.CarBookingItemModelFromProto(item))
		}

		s.cache.Set(cacheKey, bookingModel, 2*time.Hour)
		if s.redisCache != nil {
			s.redisCache.Set(ctx, cacheKey, bookingModel, 2*time.Hour)
		}
		log.Printf("Updated cache for booking: %s", result.BookingId)
	}

	return result, nil
}

func (s *bookingService) ListUserBookings(ctx context.Context, req *pb.ListBookingsRequest) (*pb.ListBookingsResponse, error) {
	if req.UserId == "" {
		return nil, status.Error(codes.InvalidArgument, "user ID must be provided")
	}
	return s.repo.ListUserBookings(ctx, req)
}

func (s *bookingService) CancelBooking(ctx context.Context, req *pb.CancelBookingRequest) (*pb.BookingResponse, error) {
	if req.BookingId == "" {
		return nil, status.Error(codes.InvalidArgument, "booking ID must be provided")
	}

	result, err := s.repo.CancelBooking(ctx, req)
	if err != nil {
		return nil, err
	}

	// Invalidate cache for this booking
	cacheKey := fmt.Sprintf("booking:%s", req.BookingId)
	s.cache.Delete(cacheKey)
	if s.redisCache != nil {
		s.redisCache.Delete(ctx, cacheKey)
	}
	log.Printf("Invalidated cache for cancelled booking: %s", req.BookingId)

	// Update cache with new data if operation was successful
	if result != nil {
		bookingModel := &models.BookingModel{
			ID:         result.BookingId,
			UserID:     result.UserId,
			Status:     result.Status,
			TotalPrice: result.TotalPrice,
		}
		// Convert booking items
		for _, item := range result.Bookings {
			bookingModel.Bookings = append(bookingModel.Bookings, models.CarBookingItemModelFromProto(item))
		}

		s.cache.Set(cacheKey, bookingModel, 2*time.Hour)
		if s.redisCache != nil {
			s.redisCache.Set(ctx, cacheKey, bookingModel, 2*time.Hour)
		}
		log.Printf("Updated cache for cancelled booking: %s", result.BookingId)
	}

	return result, nil
}
