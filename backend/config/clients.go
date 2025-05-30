package config

import (
	"log"
	"os"

	pb_booking "github.com/reallvaibhav/backend/proto/booking"
	pb_inventory "github.com/reallvaibhav/backend/proto/inventory"
	pb_statistics "github.com/reallvaibhav/backend/proto/statistics"
	pb_user "github.com/reallvaibhav/backend/proto/user"
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
}

func NewClients() (*Clients, error) {
	userServiceAddr := os.Getenv("USER_SERVICE_ADDR")
	inventoryServiceAddr := os.Getenv("INVENTORY_SERVICE_ADDR")
	bookingServiceAddr := os.Getenv("BOOKING_SERVICE_ADDR")
	statisticsServiceAddr := os.Getenv("STATISTICS_SERVICE_ADDR")
	// Get other service addresses from environment variables

	userConn, err := grpc.Dial(userServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	inventoryConn, err := grpc.Dial(inventoryServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		userConn.Close()
		return nil, err
	}

	bookingConn, err := grpc.Dial(bookingServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		userConn.Close()
		inventoryConn.Close()
		return nil, err
	}

	statisticsConn, err := grpc.Dial(statisticsServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		userConn.Close()
		inventoryConn.Close()
		bookingConn.Close()
		return nil, err
	}

	clients := &Clients{
		UserClient:       pb_user.NewUserServiceClient(userConn),
		InventoryClient:  pb_inventory.NewInventoryServiceClient(inventoryConn),
		BookingClient:    pb_booking.NewBookingServiceClient(bookingConn),
		StatisticsClient: pb_statistics.NewStatisticsServiceClient(statisticsConn),
		conns: []*grpc.ClientConn{
			userConn,
			inventoryConn,
			bookingConn,
			statisticsConn,
		},
	}

	return clients, nil
}

func (c *Clients) Close() {
	for _, conn := range c.conns {
		if err := conn.Close(); err != nil {
			log.Printf("Error closing connection: %v", err)
		}
	}
}
