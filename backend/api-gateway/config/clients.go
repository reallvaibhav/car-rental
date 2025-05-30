package config

import (
	"log"

	pb_booking "proto/booking"
	pb_inventory "proto/inventory"
	pb_statistics "proto/statistics"
	pb_user "proto/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Clients struct {
	UserClient       pb_user.UserServiceClient
	InventoryClient  pb_inventory.InventoryServiceClient
	BookingClient    pb_booking.BookingServiceClient
	StatisticsClient pb_statistics.StatisticsServiceClient
	// Add other service clients here

	conns []*grpc.ClientConn // Slice to hold connections
	cfg   *Config            // Reference to the config
}

func NewClients() (*Clients, error) {
	cfg, err := LoadConfig()
	if err != nil {
		return nil, err
	}

	clients := &Clients{
		conns: []*grpc.ClientConn{},
		cfg:   cfg,
	}
	// Connect to user service
	userConn, err := grpc.Dial(cfg.Services.UserAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Warning: Failed to connect to user service: %v", err)
	} else {
		clients.UserClient = pb_user.NewUserServiceClient(userConn)
		clients.conns = append(clients.conns, userConn)
		log.Println("Connected to user service at", cfg.Services.UserAddr)
	}

	// Connect to inventory service
	inventoryConn, err := grpc.Dial(cfg.Services.InventoryAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Warning: Failed to connect to inventory service: %v", err)
	} else {
		clients.InventoryClient = pb_inventory.NewInventoryServiceClient(inventoryConn)
		clients.conns = append(clients.conns, inventoryConn)
		log.Println("Connected to inventory service at", cfg.Services.InventoryAddr)
	}

	// Connect to booking service
	bookingConn, err := grpc.Dial(cfg.Services.BookingAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Warning: Failed to connect to booking service: %v", err)
	} else {
		clients.BookingClient = pb_booking.NewBookingServiceClient(bookingConn)
		clients.conns = append(clients.conns, bookingConn)
		log.Println("Connected to booking service at", cfg.Services.BookingAddr)
	}

	// Connect to statistics service
	statisticsConn, err := grpc.Dial(cfg.Services.StatisticsAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Warning: Failed to connect to statistics service: %v", err)
	} else {
		clients.StatisticsClient = pb_statistics.NewStatisticsServiceClient(statisticsConn)
		clients.conns = append(clients.conns, statisticsConn)
		log.Println("Connected to statistics service at", cfg.Services.StatisticsAddr)
	}

	return clients, nil
}

func (c *Clients) Close() {
	log.Println("Closing gRPC connections...")
	for _, conn := range c.conns {
		if conn != nil {
			conn.Close()
		}
	}
	log.Println("gRPC connections closed.")
}
