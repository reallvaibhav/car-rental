package nats

import (
	"encoding/json"
	"log"
	"time"

	"statistics-service/internal/usecase"

	"github.com/nats-io/nats.go"
)

type StatisticsPublisher struct {
	nc       *nats.Conn
	usecase  usecase.StatisticsUsecase
	stopChan chan struct{}
}

type StatisticsMessage struct {
	Timestamp time.Time   `json:"timestamp"`
	Type      string      `json:"type"` // "items" or "orders"
	Data      interface{} `json:"data"`
}

func NewStatisticsPublisher(nc *nats.Conn, usecase usecase.StatisticsUsecase) *StatisticsPublisher {
	return &StatisticsPublisher{
		nc:       nc,
		usecase:  usecase,
		stopChan: make(chan struct{}),
	}
}

func (sp *StatisticsPublisher) StartHourlyPublisher() {
	log.Println("Starting hourly statistics publisher...")
	ticker := time.NewTicker(1 * time.Hour)
	go func() {
		// Publish immediately on start
		sp.publishStatistics()

		for {
			select {
			case <-ticker.C:
				sp.publishStatistics()
			case <-sp.stopChan:
				ticker.Stop()
				return
			}
		}
	}()
}

func (sp *StatisticsPublisher) Stop() {
	close(sp.stopChan)
}

func (sp *StatisticsPublisher) publishStatistics() {
	// Alternate between sending items and orders statistics
	now := time.Now()
	if now.Hour()%2 == 0 {
		sp.publishItemsStatistics()
	} else {
		sp.publishOrdersStatistics()
	}
}

func (sp *StatisticsPublisher) publishItemsStatistics() {
	stats, err := sp.usecase.GetHourlyItemStatistics()
	if err != nil {
		log.Printf("Error getting items statistics: %v", err)
		return
	}

	message := StatisticsMessage{
		Timestamp: time.Now(),
		Type:      "items",
		Data:      stats,
	}

	if err := sp.publishMessage(message); err != nil {
		log.Printf("Error publishing items statistics: %v", err)
	} else {
		log.Printf("Published items statistics to ap2.statistics.event.updated")
	}
}

func (sp *StatisticsPublisher) publishOrdersStatistics() {
	stats, err := sp.usecase.GetHourlyOrderStatistics()
	if err != nil {
		log.Printf("Error getting orders statistics: %v", err)
		return
	}

	message := StatisticsMessage{
		Timestamp: time.Now(),
		Type:      "orders",
		Data:      stats,
	}

	if err := sp.publishMessage(message); err != nil {
		log.Printf("Error publishing orders statistics: %v", err)
	} else {
		log.Printf("Published orders statistics to ap2.statistics.event.updated")
	}
}

func (sp *StatisticsPublisher) publishMessage(message StatisticsMessage) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return sp.nc.Publish("ap2.statistics.event.updated", data)
}
