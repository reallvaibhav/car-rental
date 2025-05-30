package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
    // Get NATS URL from environment variable or use default
    natsURL := os.Getenv("NATS_URL")
    if natsURL == "" {
        natsURL = nats.DefaultURL // localhost:4222
    }
    
    log.Printf("Connecting to NATS server at %s...", natsURL)
    
    // Connect to NATS with retry
    var nc *nats.Conn
    var err error
    
    for i := 0; i < 5; i++ {
        nc, err = nats.Connect(natsURL)
        if err == nil {
            break
        }
        log.Printf("Error connecting to NATS: %v. Retrying in 3 seconds...", err)
        time.Sleep(3 * time.Second)
    }
    
    if err != nil {
        log.Fatalf("Failed to connect to NATS after multiple attempts: %v", err)
    }
    defer nc.Close()

    log.Println("Connected to NATS server")

    // Create JetStream Context
    js, err := nc.JetStream()
    if err != nil {
        log.Fatalf("Error getting JetStream context: %v", err)
    }

    // Create Streams
    _, err = js.AddStream(&nats.StreamConfig{
        Name:     "ORDERS",
        Subjects: []string{"order.*"},
    })
    if err != nil {
        log.Printf("Error creating ORDERS stream: %v", err)
    } else {
        log.Println("Successfully created/updated ORDERS stream")
    }

    _, err = js.AddStream(&nats.StreamConfig{
        Name:     "INVENTORY",
        Subjects: []string{"inventory.*"},
    })
    if err != nil {
        log.Printf("Error creating INVENTORY stream: %v", err)
    } else {
        log.Println("Successfully created/updated INVENTORY stream")
    }
    
    // Create Stream for user events
    _, err = js.AddStream(&nats.StreamConfig{
        Name:     "USERS",
        Subjects: []string{"user.*"},
    })
    if err != nil {
        log.Printf("Error creating USERS stream: %v", err)
    } else {
        log.Println("Successfully created/updated USERS stream")
    }
    
    // Create Stream for booking events
    _, err = js.AddStream(&nats.StreamConfig{
        Name:     "BOOKINGS",
        Subjects: []string{"booking.*"},
    })
    if err != nil {
        log.Printf("Error creating BOOKINGS stream: %v", err)
    } else {
        log.Println("Successfully created/updated BOOKINGS stream")
    }

    // Wait for interrupt signal
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
    <-sigChan

    log.Println("Shutting down NATS service...")
}