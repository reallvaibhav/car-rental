package main

import (
	"booking-service/internal/handlers"
	"booking-service/internal/repository"
	"booking-service/internal/services"
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	pb "proto/booking"
	"syscall"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func main() {
	// MongoDB connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Initialize components
	db := client.Database("car_rental")
	repo := repository.NewRepository(db.Collection("bookings"))
	svc := services.NewBookingService(repo)
	grpcHandler := handlers.NewBookingServer(svc)

	// Initialize gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterBookingServiceServer(grpcServer, grpcHandler)

	// Start gRPC server
	port := os.Getenv("GRPC_PORT")
	if port == "" {
		port = "50053"
	}
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Graceful shutdown
	go func() {
		log.Printf("Booking service running on :%s", port)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Wait for interrupt signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	// Cleanup
	log.Println("Shutting down booking service...")
	grpcServer.GracefulStop()
	if err := client.Disconnect(context.TODO()); err != nil {
		log.Printf("Error disconnecting from MongoDB: %v", err)
	}
	log.Println("Booking service stopped")
}
