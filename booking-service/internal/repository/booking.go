package repository

import (
	"context"
	"errors"

	"booking-service/internal/models"
	pb "booking-service/internal/proto"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	collection *mongo.Collection
}

func NewRepository(collection *mongo.Collection) *Repository {
	return &Repository{collection: collection}
}

// CreateBooking creates a new booking from proto request
func (r *Repository) CreateBooking(ctx context.Context, req *pb.CreateBookingRequest) (*pb.BookingResponse, error) {
	// Convert proto to model
	bookingModel := models.BookingModelFromProto(req)

	// Insert into database
	res, err := r.collection.InsertOne(ctx, bookingModel)
	if err != nil {
		return nil, err
	}

	// Update ID
	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		bookingModel.ID = oid.Hex()
	}

	// Calculate total price
	var totalPrice float64
	for _, item := range bookingModel.Bookings {
		totalPrice += item.PricePerDay * float64(item.TotalDays)
	}
	bookingModel.TotalPrice = totalPrice

	// Return proto response
	return models.BookingModelToProto(bookingModel), nil
}

// GetBookingByID retrieves a booking by ID
func (r *Repository) GetBookingByID(ctx context.Context, id string) (*models.BookingModel, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var booking models.BookingModel
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&booking)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("booking not found")
		}
		return nil, err
	}
	return &booking, nil
}

// ListUserBookings retrieves all bookings for a user
func (r *Repository) ListUserBookings(ctx context.Context, req *pb.ListBookingsRequest) (*pb.ListBookingsResponse, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"user_id": req.UserId})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var bookings []*models.BookingModel
	if err = cursor.All(ctx, &bookings); err != nil {
		return nil, err
	}

	response := &pb.ListBookingsResponse{
		Bookings: make([]*pb.BookingResponse, 0, len(bookings)),
	}

	for _, booking := range bookings {
		response.Bookings = append(response.Bookings, models.BookingModelToProto(booking))
	}

	return response, nil
}

// UpdateBookingStatus updates the status of a booking
func (r *Repository) UpdateBookingStatus(ctx context.Context, req *pb.UpdateBookingStatusRequest) (*pb.BookingResponse, error) {
	objectID, err := primitive.ObjectIDFromHex(req.BookingId)
	if err != nil {
		return nil, err
	}

	_, err = r.collection.UpdateOne(
		ctx,
		bson.M{"_id": objectID},
		bson.M{"$set": bson.M{"status": req.Status}},
	)
	if err != nil {
		return nil, err
	}

	// Get updated booking
	booking, err := r.GetBookingByID(ctx, req.BookingId)
	if err != nil {
		return nil, err
	}

	return models.BookingModelToProto(booking), nil
}

// CancelBooking cancels a booking
func (r *Repository) CancelBooking(ctx context.Context, req *pb.CancelBookingRequest) (*pb.BookingResponse, error) {
	return r.UpdateBookingStatus(ctx, &pb.UpdateBookingStatusRequest{
		BookingId: req.BookingId,
		Status:    "Cancelled",
	})
}
