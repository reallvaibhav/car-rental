package handlers

import (
	"net/http"

	pb "proto/booking"

	"github.com/gin-gonic/gin"
)

// Add booking-related handler methods to the Handler struct

// CreateBooking handles POST /bookings
func (h *Handler) CreateBooking(c *gin.Context) {
	var req struct {
		UserID   string `json:"userId"`
		Bookings []struct {
			CarID       string  `json:"carId"`
			StartDate   string  `json:"startDate"`
			EndDate     string  `json:"endDate"`
			PricePerDay float64 `json:"pricePerDay"`
			TotalDays   int32   `json:"totalDays"`
		} `json:"bookings"`
	}

	// Decode request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	// Validate request
	if req.UserID == "" || len(req.Bookings) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID and at least one booking item are required"})
		return
	}

	// Convert to protobuf request
	protoReq := &pb.CreateBookingRequest{
		UserId:   req.UserID,
		Bookings: make([]*pb.CarBookingItem, 0, len(req.Bookings)),
	}

	for _, item := range req.Bookings {
		protoReq.Bookings = append(protoReq.Bookings, &pb.CarBookingItem{
			CarId:       item.CarID,
			StartDate:   item.StartDate,
			EndDate:     item.EndDate,
			PricePerDay: item.PricePerDay,
			TotalDays:   item.TotalDays,
		})
	}

	// Call booking service
	resp, err := h.clients.BookingClient.CreateBooking(c.Request.Context(), protoReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create booking: " + err.Error()})
		return
	}

	// Send response
	c.JSON(http.StatusCreated, resp)
}

// GetBooking handles GET /bookings/:id
func (h *Handler) GetBooking(c *gin.Context) {
	// Get booking ID from URL params
	bookingID := c.Param("id")
	if bookingID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Booking ID is required"})
		return
	}

	// Call booking service
	resp, err := h.clients.BookingClient.GetBooking(c.Request.Context(), &pb.GetBookingRequest{
		BookingId: bookingID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get booking: " + err.Error()})
		return
	}

	// Send response
	c.JSON(http.StatusOK, resp)
}

// GetUserBookings handles GET /bookings/user
func (h *Handler) GetUserBookings(c *gin.Context) {
	// Get user ID from query params
	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id query parameter is required"})
		return
	}

	// Call booking service
	resp, err := h.clients.BookingClient.ListUserBookings(c.Request.Context(), &pb.ListBookingsRequest{
		UserId: userID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list bookings: " + err.Error()})
		return
	}

	// Send response
	c.JSON(http.StatusOK, resp)
}

// GetFleetOwnerBookings handles GET /bookings/fleet-owner
func (h *Handler) GetFleetOwnerBookings(c *gin.Context) {
	// Note: Since the original proto doesn't have a specific method for fleet owner bookings,
	// we'll improvise by adapting the ListUserBookings method

	// Get owner ID from query params
	ownerID := c.Query("owner_id")
	if ownerID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "owner_id query parameter is required"})
		return
	}

	// For now, use the ListUserBookings since we don't have a specific fleet owner API
	// In a production app, you would add a ListCarBookings method to the proto
	resp, err := h.clients.BookingClient.ListUserBookings(c.Request.Context(), &pb.ListBookingsRequest{
		UserId: ownerID, // Using owner ID as user ID for now
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list bookings: " + err.Error()})
		return
	}

	// Send response
	c.JSON(http.StatusOK, resp)
}

// UpdateBookingStatus handles PUT /bookings/:id/status
func (h *Handler) UpdateBookingStatus(c *gin.Context) {
	// Get booking ID from URL params
	bookingID := c.Param("id")
	if bookingID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Booking ID is required"})
		return
	}

	// Decode request body
	var req struct {
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	// Validate request
	if req.Status == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status is required"})
		return
	}

	// Call booking service
	resp, err := h.clients.BookingClient.UpdateBookingStatus(c.Request.Context(), &pb.UpdateBookingStatusRequest{
		BookingId: bookingID,
		Status:    req.Status,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update booking status: " + err.Error()})
		return
	}

	// Send response
	c.JSON(http.StatusOK, resp)
}

// CancelBooking handles DELETE /bookings/:id
func (h *Handler) CancelBooking(c *gin.Context) {
	// Get booking ID from URL params
	bookingID := c.Param("id")
	if bookingID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Booking ID is required"})
		return
	}

	// Call booking service
	resp, err := h.clients.BookingClient.CancelBooking(c.Request.Context(), &pb.CancelBookingRequest{
		BookingId: bookingID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cancel booking: " + err.Error()})
		return
	}

	// Send response
	c.JSON(http.StatusOK, resp)
}
