package main

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"

	"email-service/internal/service"

	"github.com/nats-io/nats.go"
)

func main() {
	// Email configuration
	emailConfig := service.EmailConfig{
		Host:     os.Getenv("SMTP_HOST"), // e.g., "smtp.gmail.com"
		Port:     587,                    // Standard SMTP port
		Username: os.Getenv("SMTP_USERNAME"),
		Password: os.Getenv("SMTP_PASSWORD"),
		From:     os.Getenv("SMTP_FROM"),
	}

	// Create email service
	emailService := service.NewEmailService(emailConfig)

	// Connect to NATS
	nc, err := nats.Connect(os.Getenv("NATS_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}
	defer nc.Close()

	// Subscribe to booking events
	nc.Subscribe("booking.created", func(msg *nats.Msg) {
		var data map[string]interface{}
		if err := json.Unmarshal(msg.Data, &data); err != nil {
			log.Printf("Error unmarshaling booking data: %v", err)
			return
		}

		userEmail := data["user_email"].(string)
		if err := emailService.SendBookingConfirmation(userEmail, data); err != nil {
			log.Printf("Error sending booking confirmation: %v", err)
		}
	})

	nc.Subscribe("booking.status_updated", func(msg *nats.Msg) {
		var data map[string]interface{}
		if err := json.Unmarshal(msg.Data, &data); err != nil {
			log.Printf("Error unmarshaling status update data: %v", err)
			return
		}

		userEmail := data["user_email"].(string)
		bookingID := data["booking_id"].(string)
		status := data["status"].(string)

		if err := emailService.SendBookingStatusUpdate(userEmail, bookingID, status); err != nil {
			log.Printf("Error sending status update email: %v", err)
		}
	})

	nc.Subscribe("user.registered", func(msg *nats.Msg) {
		var data map[string]interface{}
		if err := json.Unmarshal(msg.Data, &data); err != nil {
			log.Printf("Error unmarshaling user data: %v", err)
			return
		}

		userEmail := data["email"].(string)
		userName := data["name"].(string)

		if err := emailService.SendWelcomeEmail(userEmail, userName); err != nil {
			log.Printf("Error sending welcome email: %v", err)
		}
	})

	// Wait for interrupt signal
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	log.Println("Shutting down email service...")
}
