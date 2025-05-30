package main

import (
	"context"
	"fmt"
	"github.com/Car-Rental/backend/inventory-service/internal/config"
	"github.com/Car-Rental/backend/inventory-service/internal/handler"
	"github.com/Car-Rental/backend/inventory-service/internal/publisher"
	"github.com/Car-Rental/backend/inventory-service/internal/repository"
	"github.com/Car-Rental/backend/inventory-service/internal/service"	"log"
	"net"
	"os"
	"os/signal"
	pb "proto/inventory"
	"syscall"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()

	// Initialize MongoDB
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		log.Printf("Warning: Failed to connect to MongoDB: %v", err)
		log.Println("Running in testing mode without MongoDB...")
	} else {
		// Ping the database
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			log.Printf("Warning: Failed to ping MongoDB: %v", err)
			log.Println("Running in testing mode without MongoDB...")
			client = nil
		}
	}

	// Initialize dependencies
	db := client.Database("car_rental")
	natsPublisher := publisher.New(cfg.NATSURL)
	repo := repository.NewCarRepository(db)
	svc := service.NewCarService(repo, natsPublisher)

	// Initialize gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterInventoryServiceServer(grpcServer, handler.NewInventoryServer(svc))

	// Start gRPC server
	lis, err := net.Listen("tcp", ":"+cfg.GRPCPort)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", cfg.GRPCPort, err)
	}

	// Create a channel to listen for interrupt signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Start server in a goroutine
	go func() {
		fmt.Printf("Inventory service gRPC server running on port %s\n", cfg.GRPCPort)
		if err := grpcServer.Serve(lis); err != nil {
			log.Printf("Error serving gRPC: %v", err)
		}
	}()

	// Wait for interrupt signal
	<-sigChan
	log.Println("Received shutdown signal")

	// Gracefully stop the gRPC server
	grpcServer.GracefulStop()
	log.Println("Server stopped gracefully")

	// Clean up MongoDB connection
	if client != nil {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Printf("Error disconnecting from MongoDB: %v", err)
		}
	}
}
