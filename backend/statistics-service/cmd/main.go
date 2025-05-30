package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"statistics-service/internal/delivery"
	"statistics-service/internal/nats"
	"statistics-service/internal/repository"
	"statistics-service/internal/usecase"
	"statistics-service/proto"
	"syscall"

	natsgo "github.com/nats-io/nats.go"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Starting Statistics Service...")

	// Initialize database connections
	repo, err := repository.NewRepository(
		"postgres://postgres:1234@localhost:5432/statistics?sslmode=disable",
		"mongodb://root:example@localhost:27017/ecommerce?authSource=admin",
	)
	if err != nil {
		log.Fatalf("Failed to initialize repositories: %v", err)
	}

	// Initialize NATS connection
	natsConn, err := natsgo.Connect(natsgo.DefaultURL)
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}
	defer natsConn.Close()

	// Initialize use case layer
	statsUsecase := usecase.NewStatisticsUsecase(repo)

	// Initialize NATS publisher
	publisher := nats.NewStatisticsPublisher(natsConn, statsUsecase)
	publisher.StartHourlyPublisher()
	defer publisher.Stop()

	// Start NATS message listeners
	startNATSListeners(natsConn, statsUsecase)

	// Start gRPC server in a goroutine
	go startGRPCServer(statsUsecase)

	// Wait for interrupt signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	log.Println("Shutting down Statistics Service...")
}

func startNATSListeners(nc *nats.Conn, usecase usecase.StatisticsUsecase) {
	if err := listenToOrderEvents(nc, usecase); err != nil {
		log.Fatalf("Failed to start order events listener: %v", err)
	}
	if err := listenToInventoryEvents(nc, usecase); err != nil {
		log.Fatalf("Failed to start inventory events listener: %v", err)
	}
}

func listenToOrderEvents(nc *nats.Conn, usecase usecase.StatisticsUsecase) error {
	_, err := nc.Subscribe("order.created", func(msg *nats.Msg) {
		usecase.ProcessOrderCreated(msg.Data)
	})
	return err
}

func listenToInventoryEvents(nc *nats.Conn, usecase usecase.StatisticsUsecase) error {
	_, err := nc.Subscribe("inventory.updated", func(msg *nats.Msg) {
		usecase.ProcessInventoryUpdated(msg.Data)
	})
	return err
}

func startGRPCServer(usecase usecase.StatisticsUsecase) {
	lis, err := net.Listen("tcp", ":8084")
	if err != nil {
		log.Fatalf("Failed to listen on port 8084: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterStatisticsServiceServer(grpcServer, &delivery.Server{Usecase: usecase})

	log.Printf("gRPC server listening on port 8084")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
