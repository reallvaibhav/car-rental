package handlers

import (
	"net/http"
	"strconv"
	"time"

	pb "proto/statistics"

	"github.com/gin-gonic/gin"
)

// GetBookingStats handles GET /statistics/bookings
func (h *Handler) GetBookingStats(c *gin.Context) {
	timeRange := c.Query("time_range") // e.g., "daily", "weekly", "monthly"
	if timeRange == "" {
		timeRange = "monthly" // Default to monthly if not specified
	}

	// Convert timeRange string to appropriate start and end dates
	var startDate, endDate string
	now := time.Now()

	switch timeRange {
	case "daily":
		startDate = now.Add(-24 * time.Hour).Format("2006-01-02")
		endDate = now.Format("2006-01-02")
	case "weekly":
		startDate = now.Add(-7 * 24 * time.Hour).Format("2006-01-02")
		endDate = now.Format("2006-01-02")
	case "monthly":
		startDate = now.Add(-30 * 24 * time.Hour).Format("2006-01-02")
		endDate = now.Format("2006-01-02")
	case "yearly":
		startDate = now.Add(-365 * 24 * time.Hour).Format("2006-01-02")
		endDate = now.Format("2006-01-02")
	default:
		startDate = now.Add(-30 * 24 * time.Hour).Format("2006-01-02")
		endDate = now.Format("2006-01-02")
	}

	// Get location from query params (optional)
	location := c.Query("location")

	res, err := h.clients.StatisticsClient.GetBookingStats(c.Request.Context(), &pb.BookingStatsRequest{
		TimeRange: &pb.TimeRange{
			StartDate: startDate,
			EndDate:   endDate,
		},
		Location: location,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetCarStats handles GET /statistics/cars
func (h *Handler) GetCarStats(c *gin.Context) {
	timeRange := c.Query("time_range")
	if timeRange == "" {
		timeRange = "monthly"
	}

	// Convert timeRange string to appropriate start and end dates
	var startDate, endDate string
	now := time.Now()

	switch timeRange {
	case "daily":
		startDate = now.Add(-24 * time.Hour).Format("2006-01-02")
		endDate = now.Format("2006-01-02")
	case "weekly":
		startDate = now.Add(-7 * 24 * time.Hour).Format("2006-01-02")
		endDate = now.Format("2006-01-02")
	case "monthly":
		startDate = now.Add(-30 * 24 * time.Hour).Format("2006-01-02")
		endDate = now.Format("2006-01-02")
	case "yearly":
		startDate = now.Add(-365 * 24 * time.Hour).Format("2006-01-02")
		endDate = now.Format("2006-01-02")
	default:
		startDate = now.Add(-30 * 24 * time.Hour).Format("2006-01-02")
		endDate = now.Format("2006-01-02")
	}

	// Get category from query params (optional)
	category := c.Query("category")

	res, err := h.clients.StatisticsClient.GetCarStats(c.Request.Context(), &pb.CarStatsRequest{
		TimeRange: &pb.TimeRange{
			StartDate: startDate,
			EndDate:   endDate,
		},
		Category: category,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetRevenueStats handles GET /statistics/revenue
func (h *Handler) GetRevenueStats(c *gin.Context) {
	timeRange := c.Query("time_range")
	if timeRange == "" {
		timeRange = "monthly"
	}

	// Convert timeRange string to appropriate start and end dates
	var startDate, endDate string
	now := time.Now()

	switch timeRange {
	case "daily":
		startDate = now.Add(-24 * time.Hour).Format("2006-01-02")
		endDate = now.Format("2006-01-02")
	case "weekly":
		startDate = now.Add(-7 * 24 * time.Hour).Format("2006-01-02")
		endDate = now.Format("2006-01-02")
	case "monthly":
		startDate = now.Add(-30 * 24 * time.Hour).Format("2006-01-02")
		endDate = now.Format("2006-01-02")
	case "yearly":
		startDate = now.Add(-365 * 24 * time.Hour).Format("2006-01-02")
		endDate = now.Format("2006-01-02")
	default:
		startDate = now.Add(-30 * 24 * time.Hour).Format("2006-01-02")
		endDate = now.Format("2006-01-02")
	}

	// Get location from query params (optional)
	location := c.Query("location")

	res, err := h.clients.StatisticsClient.GetRevenueStats(c.Request.Context(), &pb.RevenueStatsRequest{
		TimeRange: &pb.TimeRange{
			StartDate: startDate,
			EndDate:   endDate,
		},
		Location: location,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetPopularLocations handles GET /statistics/popular-locations
func (h *Handler) GetPopularLocations(c *gin.Context) {
	timeRange := c.Query("time_range")
	if timeRange == "" {
		timeRange = "monthly"
	}

	// Convert timeRange string to appropriate start and end dates
	var startDate, endDate string
	now := time.Now()

	switch timeRange {
	case "daily":
		startDate = now.Add(-24 * time.Hour).Format("2006-01-02")
		endDate = now.Format("2006-01-02")
	case "weekly":
		startDate = now.Add(-7 * 24 * time.Hour).Format("2006-01-02")
		endDate = now.Format("2006-01-02")
	case "monthly":
		startDate = now.Add(-30 * 24 * time.Hour).Format("2006-01-02")
		endDate = now.Format("2006-01-02")
	case "yearly":
		startDate = now.Add(-365 * 24 * time.Hour).Format("2006-01-02")
		endDate = now.Format("2006-01-02")
	default:
		startDate = now.Add(-30 * 24 * time.Hour).Format("2006-01-02")
		endDate = now.Format("2006-01-02")
	}

	// Get limit from query params (optional)
	limitStr := c.Query("limit")
	limit := int32(10) // Default limit
	if limitStr != "" {
		if limitInt, err := strconv.Atoi(limitStr); err == nil && limitInt > 0 {
			limit = int32(limitInt)
		}
	}

	res, err := h.clients.StatisticsClient.GetPopularLocations(c.Request.Context(), &pb.PopularLocationsRequest{
		TimeRange: &pb.TimeRange{
			StartDate: startDate,
			EndDate:   endDate,
		},
		Limit: limit,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetUserStats handles GET /statistics/users
func (h *Handler) GetUserStats(c *gin.Context) {
	timeRange := c.Query("time_range")
	if timeRange == "" {
		timeRange = "monthly"
	}

	// Convert timeRange string to appropriate start and end dates
	var startDate, endDate string
	now := time.Now()

	switch timeRange {
	case "daily":
		startDate = now.Add(-24 * time.Hour).Format("2006-01-02")
		endDate = now.Format("2006-01-02")
	case "weekly":
		startDate = now.Add(-7 * 24 * time.Hour).Format("2006-01-02")
		endDate = now.Format("2006-01-02")
	case "monthly":
		startDate = now.Add(-30 * 24 * time.Hour).Format("2006-01-02")
		endDate = now.Format("2006-01-02")
	case "yearly":
		startDate = now.Add(-365 * 24 * time.Hour).Format("2006-01-02")
		endDate = now.Format("2006-01-02")
	default:
		startDate = now.Add(-30 * 24 * time.Hour).Format("2006-01-02")
		endDate = now.Format("2006-01-02")
	}

	res, err := h.clients.StatisticsClient.GetUserStats(c.Request.Context(), &pb.UserStatsRequest{
		TimeRange: &pb.TimeRange{
			StartDate: startDate,
			EndDate:   endDate,
		},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetMetrics handles GET /metrics
func (h *Handler) GetMetrics(c *gin.Context) {
	// This is a simple implementation that returns basic service status
	metrics := map[string]interface{}{
		"status": "ok",
		"services": map[string]string{
			"user_service":       "connected",
			"inventory_service":  "connected",
			"booking_service":    "connected",
			"statistics_service": "connected",
		},
		"timestamp": time.Now().Format(time.RFC3339),
	}

	c.JSON(http.StatusOK, metrics)
}
