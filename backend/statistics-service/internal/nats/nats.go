package nats

import (
	"statistics-service/internal/usecase"

	"github.com/nats-io/nats.go"
)

const DefaultURL = nats.DefaultURL

func Connect(url string) (*nats.Conn, error) {
	return nats.Connect(url)
}

func ListenToOrderEvents(nc *nats.Conn, usecase usecase.StatisticsUsecase) error {
	_, err := nc.Subscribe("order.created", func(msg *nats.Msg) {
		usecase.ProcessOrderCreated(msg.Data)
	})
	return err
}

func ListenToInventoryEvents(nc *nats.Conn, usecase usecase.StatisticsUsecase) error {
	_, err := nc.Subscribe("inventory.updated", func(msg *nats.Msg) {
		usecase.ProcessInventoryUpdated(msg.Data)
	})
	return err
}
