package publisher

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
)

type NatsPublisher struct {
	conn *nats.Conn
}

func New(natsURL string) *NatsPublisher {
	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Printf("Warning: Failed to connect to NATS: %v", err)
		return &NatsPublisher{conn: nil}
	}
	return &NatsPublisher{conn: nc}
}

func (p *NatsPublisher) Publish(subject string, data interface{}) error {
	if p.conn == nil {
		log.Printf("Warning: NATS not connected, skipping publish to %s", subject)
		return nil
	}

	msg, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return p.conn.Publish(subject, msg)
}

func (p *NatsPublisher) Close() {
	if p.conn != nil {
		p.conn.Close()
	}
}
