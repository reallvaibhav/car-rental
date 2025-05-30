package service

import (
	"github.com/Car-Rental/backend/inventory-service/internal/publisher"
	"github.com/Car-Rental/backend/inventory-service/internal/repository"
	"github.com/Car-Rental/backend/inventory-service/model"

	"go.mongodb.org/mongo-driver/bson"
)

type CarService struct {
	repo      *repository.CarRepository
	publisher *publisher.NatsPublisher
}

func NewCarService(repo *repository.CarRepository, pub *publisher.NatsPublisher) *CarService {
	return &CarService{
		repo:      repo,
		publisher: pub,
	}
}

func (s *CarService) AddCar(make, modelName string, year int32, category, location string, pricePerDay float64, features []string) (*model.Car, error) {
	car := &model.Car{
		Make:        make,
		Model:       modelName,
		Year:        year,
		Category:    category,
		Location:    location,
		PricePerDay: pricePerDay,
		Features:    features,
		Available:   true,
	}

	if err := s.repo.Create(car); err != nil {
		return nil, err
	}

	// Publish car.created event
	s.publisher.Publish("car.created", map[string]interface{}{
		"car_id":   car.ID.Hex(),
		"make":     car.Make,
		"model":    car.Model,
		"location": car.Location,
	})

	return car, nil
}

func (s *CarService) UpdateCar(id, make, modelName string, year int32, category, location string, pricePerDay float64, features []string) (*model.Car, error) {
	car, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	car.Make = make
	car.Model = modelName
	car.Year = year
	car.Category = category
	car.Location = location
	car.PricePerDay = pricePerDay
	car.Features = features

	if err := s.repo.Update(car); err != nil {
		return nil, err
	}

	// Publish car.updated event
	s.publisher.Publish("car.updated", map[string]interface{}{
		"car_id":   car.ID.Hex(),
		"make":     car.Make,
		"model":    car.Model,
		"location": car.Location,
	})

	return car, nil
}

func (s *CarService) DeleteCar(id string) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}

	// Publish car.deleted event
	s.publisher.Publish("car.deleted", map[string]interface{}{
		"car_id": id,
	})

	return nil
}

func (s *CarService) SearchCars(location, category string, maxPrice float64, availableOnly bool) ([]*model.Car, error) {
	filter := bson.M{}

	if location != "" {
		filter["location"] = location
	}
	if category != "" {
		filter["category"] = category
	}
	if maxPrice > 0 {
		filter["price_per_day"] = bson.M{"$lte": maxPrice}
	}
	if availableOnly {
		filter["available"] = true
	}

	return s.repo.Search(filter)
}

func (s *CarService) UpdateAvailability(carID string, available bool) error {
	if err := s.repo.UpdateAvailability(carID, available); err != nil {
		return err
	}

	// Publish car.availability.updated event
	s.publisher.Publish("car.availability.updated", map[string]interface{}{
		"car_id":    carID,
		"available": available,
	})

	return nil
}

func (s *CarService) GetCarByID(id string) (*model.Car, error) {
	return s.repo.GetByID(id)
}
