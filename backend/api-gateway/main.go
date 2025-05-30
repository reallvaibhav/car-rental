package main

import (
	"api-gateway/config"
	"api-gateway/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func corsMiddleware(allowedOrigins []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")

		// Check if the request origin is allowed
		allowAll := false
		allowed := false

		for _, allowedOrigin := range allowedOrigins {
			if allowedOrigin == "*" {
				allowAll = true
				break
			}
			if origin == allowedOrigin {
				allowed = true
				break
			}
		}

		// Set CORS headers
		if allowAll {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		} else if allowed {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		}

		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize clients
	clients, err := config.NewClients()
	if err != nil {
		log.Fatalf("Failed to initialize clients: %v", err)
	}
	defer clients.Close()

	// Set up Gin
	if !cfg.Server.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.Use(corsMiddleware(cfg.CORS.AllowedOrigins))
	r.Use(handlers.ErrorHandler())

	metrics := handlers.NewMetrics()
	// Apply Logger and Metrics middleware globally
	r.Use(handlers.Logger(), metrics.Track())

	h := handlers.NewHandler(clients)

	// Public routes (no auth required)
	r.POST("/auth/register", h.Register)
	r.POST("/auth/login", h.Login)

	// Authenticated routes
	authRoutes := r.Group("/").Use(handlers.Auth())
	{
		// Car routes
		authRoutes.POST("/cars", h.CreateCar)
		authRoutes.GET("/cars", h.ListCars)
		authRoutes.GET("/cars/:id", h.GetCar)
		authRoutes.PUT("/cars/:id", h.UpdateCar)
		authRoutes.DELETE("/cars/:id", h.DeleteCar)

		// Booking routes
		authRoutes.POST("/bookings", h.CreateBooking)
		authRoutes.GET("/bookings/user", h.GetUserBookings)
		authRoutes.GET("/bookings/fleet-owner", h.GetFleetOwnerBookings)
		authRoutes.GET("/bookings/:id", h.GetBooking)
		authRoutes.PUT("/bookings/:id/status", h.UpdateBookingStatus)
		authRoutes.DELETE("/bookings/:id", h.CancelBooking)

		// Statistics routes
		authRoutes.GET("/statistics/bookings", h.GetBookingStats)
		authRoutes.GET("/statistics/cars", h.GetCarStats)
		authRoutes.GET("/statistics/revenue", h.GetRevenueStats)
		authRoutes.GET("/statistics/popular-locations", h.GetPopularLocations)
		authRoutes.GET("/statistics/users", h.GetUserStats)

		// Metrics endpoint
		authRoutes.GET("/metrics", h.GetMetrics)
	}

	serverAddr := cfg.Server.Host + ":" + cfg.Server.Port
	log.Printf("API Gateway running on %s", serverAddr)
	if err := r.Run(serverAddr); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
