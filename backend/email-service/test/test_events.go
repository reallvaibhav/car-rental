package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to NATS
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}
	defer nc.Close()

	// Test user registration event
	userEvent := map[string]interface{}{
		"email": "test@example.com",
		"name":  "Test User",
	}
	userJSON, _ := json.Marshal(userEvent)
	if err := nc.Publish("user.registered", userJSON); err != nil {
		log.Fatalf("Failed to publish user event: %v", err)
	}
	fmt.Println("Published user.registered event")

	// Test booking created event
	bookingEvent := map[string]interface{}{
		"booking_id":  "123456",
		"user_email":  "test@example.com",
		"car_make":    "Toyota",
		"car_model":   "Camry",
		"start_date":  "2025-06-01",
		"end_date":    "2025-06-05",
		"total_price": 250.00,
	}
	bookingJSON, _ := json.Marshal(bookingEvent)
	if err := nc.Publish("booking.created", bookingJSON); err != nil {
		log.Fatalf("Failed to publish booking event: %v", err)
	}
	fmt.Println("Published booking.created event")

	// Ensure messages are delivered before exiting
	nc.Flush()
}
