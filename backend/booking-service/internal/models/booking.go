package models

import pb "proto/booking"

// CarBookingItemModel is the MongoDB representation of a booking item
type CarBookingItemModel struct {
	CarID       string  `bson:"car_id"`
	StartDate   string  `bson:"start_date"`
	EndDate     string  `bson:"end_date"`
	PricePerDay float64 `bson:"price_per_day"`
	TotalDays   int32   `bson:"total_days"`
}

type BookingModel struct {
	ID         string                `bson:"_id,omitempty"`
	UserID     string                `bson:"user_id"`
	Bookings   []CarBookingItemModel `bson:"bookings"`
	Status     string                `bson:"status"`
	TotalPrice float64               `bson:"total_price"`
}

// CarBookingItemModelFromProto converts a proto item to a model item
func CarBookingItemModelFromProto(p *pb.CarBookingItem) CarBookingItemModel {
	return CarBookingItemModel{
		CarID:       p.GetCarId(),
		StartDate:   p.GetStartDate(),
		EndDate:     p.GetEndDate(),
		PricePerDay: p.GetPricePerDay(),
		TotalDays:   p.GetTotalDays(),
	}
}

// CarBookingItemModelToProto converts a model item to a proto item
func CarBookingItemModelToProto(m CarBookingItemModel) *pb.CarBookingItem {
	return &pb.CarBookingItem{
		CarId:       m.CarID,
		StartDate:   m.StartDate,
		EndDate:     m.EndDate,
		PricePerDay: m.PricePerDay,
		TotalDays:   m.TotalDays,
	}
}

// BookingModelFromProto creates a booking model from a proto request
func BookingModelFromProto(p *pb.CreateBookingRequest) *BookingModel {
	items := make([]CarBookingItemModel, len(p.GetBookings()))
	for i, item := range p.GetBookings() {
		items[i] = CarBookingItemModelFromProto(item)
	}
	return &BookingModel{
		UserID:   p.GetUserId(),
		Bookings: items,
		Status:   "Pending",
	}
}

// BookingModelToProto converts a booking model to a proto response
func BookingModelToProto(m *BookingModel) *pb.BookingResponse {
	items := make([]*pb.CarBookingItem, len(m.Bookings))
	for i, item := range m.Bookings {
		items[i] = CarBookingItemModelToProto(item)
	}
	return &pb.BookingResponse{
		BookingId:  m.ID,
		UserId:     m.UserID,
		Bookings:   items,
		Status:     m.Status,
		TotalPrice: m.TotalPrice,
	}
}
