package services

import (
	"encoding/json"

	"github.com/nats-io/nats.go"
)

type NatsService struct {
	conn *nats.Conn
}

func NewNatsService(url string) (*NatsService, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}
	return &NatsService{conn: nc}, nil
}

func (s *NatsService) Publish(subject string, data interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return s.conn.Publish(subject, jsonData)
}

func (s *NatsService) Subscribe(subject string, handler func([]byte)) error {
	_, err := s.conn.Subscribe(subject, func(msg *nats.Msg) {
		handler(msg.Data)
	})
	return err
}

func (s *NatsService) Close() {
	if s.conn != nil {
		s.conn.Close()
	}
}

// Example usage of NATS events
const (
	BookingCreatedEvent = "booking.created"
	BookingUpdatedEvent = "booking.updated"
	BookingDeletedEvent = "booking.deleted"
)
