package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb_user "github.com/Car-Rental/proto/user"

	"github.com/Car-Rental/backend/user-service/cache"
	"github.com/Car-Rental/backend/user-service/internal/auth"
	"github.com/Car-Rental/backend/user-service/internal/config"
	"github.com/Car-Rental/backend/user-service/internal/handler"
	"github.com/Car-Rental/backend/user-service/internal/publisher"
	"github.com/Car-Rental/backend/user-service/internal/repository"
	"github.com/Car-Rental/backend/user-service/internal/service"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()
	var db *mongo.Database
	var client *mongo.Client
	var err error

	// Try to connect to MongoDB
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		log.Printf("Warning: Failed to connect to MongoDB: %v", err)
		log.Println("Running in testing mode without MongoDB...")
		client = nil
	} else {
		// Ping the database
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			log.Printf("Warning: Failed to ping MongoDB: %v", err)
			log.Println("Running in testing mode without MongoDB...")
			client = nil
		} else {
			db = client.Database("car_rental")
		}
	}
	jwtManager := auth.New(cfg.JWTSecret)
	natsPublisher := publisher.New(cfg.NATSURL)

	// Initialize caching
	inMemoryCache := cache.New()
	redisCache := cache.NewRedisCache("localhost:6379", "", 0)
	if redisCache != nil {
		defer redisCache.Close()
		log.Println("Redis cache initialized for user-service")
	} else {
		log.Println("Redis cache not available, using in-memory cache only")
	}

	repo := repository.NewUserRepository(db)
	svc := service.New(repo, jwtManager, natsPublisher, inMemoryCache, redisCache)
	grpcServer := grpc.NewServer()
	pb_user.RegisterUserServiceServer(grpcServer, handler.New(svc))

	lis, err := net.Listen("tcp", ":"+cfg.GRPCPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a channel to listen for interrupt signals
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		fmt.Println("gRPC server running on port", cfg.GRPCPort)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Wait for interrupt signal
	<-c
	fmt.Println("\nShutting down gRPC server...")
	grpcServer.GracefulStop()

	// Clean up MongoDB connection
	if client != nil {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Printf("Error disconnecting from MongoDB: %v", err)
		}
	}

	fmt.Println("Server stopped gracefully")
}
