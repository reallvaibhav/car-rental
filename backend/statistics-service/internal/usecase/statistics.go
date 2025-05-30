package usecase

import (
	"log"
)

type StatisticsUsecase interface {
	GetHourlyItemStatistics() (*ItemStatistics, error)
	GetHourlyOrderStatistics() (*OrderStatistics, error)
	ProcessOrderCreated(data []byte)
	ProcessInventoryUpdated(data []byte)
}

type statisticsUsecase struct {
	repo Repository
}

type Repository interface {
	GetItemStatistics() (*ItemStatistics, error)
	GetOrderStatistics() (*OrderStatistics, error)
}

func NewStatisticsUsecase(repo Repository) StatisticsUsecase {
	return &statisticsUsecase{
		repo: repo,
	}
}

func (u *statisticsUsecase) GetHourlyItemStatistics() (*ItemStatistics, error) {
	stats, err := u.repo.GetItemStatistics()
	if err != nil {
		log.Printf("Error getting item statistics: %v", err)
		return nil, err
	}
	return stats, nil
}

func (u *statisticsUsecase) GetHourlyOrderStatistics() (*OrderStatistics, error) {
	stats, err := u.repo.GetOrderStatistics()
	if err != nil {
		log.Printf("Error getting order statistics: %v", err)
		return nil, err
	}
	return stats, nil
}

func (u *statisticsUsecase) ProcessOrderCreated(data []byte) {
	// Implementation for processing order creation events
	log.Printf("Processing order created event")
}

func (u *statisticsUsecase) ProcessInventoryUpdated(data []byte) {
	// Implementation for processing inventory updates
	log.Printf("Processing inventory updated event")
}

// Statistics types are defined in statistics_models.go
