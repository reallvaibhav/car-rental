package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to NATS
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("Error connecting to NATS: %v", err)
	}
	defer nc.Close()

	log.Println("Connected to NATS, waiting for statistics messages...")

	// Subscribe to statistics updates
	sub, err := nc.Subscribe("ap2.statistics.event.updated", func(msg *nats.Msg) {
		log.Printf("Received statistics message:\n%s", string(msg.Data))
	})
	if err != nil {
		log.Fatalf("Error subscribing: %v", err)
	}
	defer sub.Unsubscribe()

	// Wait for interrupt signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan

	log.Println("Shutting down subscriber...")
}
