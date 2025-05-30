package handler

import (
	"context"	
	pb "proto/inventory"

	"github.com/Car-Rental/backend/inventory-service/internal/service"
	"github.com/Car-Rental/backend/inventory-service/model"
)

type InventoryServer struct {	
	pb.UnimplementedInventoryServiceServer
	service *service.CarService
}

func NewInventoryServer(svc *service.CarService) *InventoryServer {
	return &InventoryServer{service: svc}
}

func (s *InventoryServer) CreateCar(ctx context.Context, req *pb.CreateCarRequest) (*pb.CreateCarResponse, error) {
	car, err := s.service.AddCar(
		req.Make,
		req.Model,
		req.Year,
		"", // Category - not in proto
		"", // Location - not in proto
		0,  // PricePerDay - not in proto
		[]string{}, // Features - not in proto
	)
	if err != nil {
		return nil, err
	}

	return &pb.CreateCarResponse{
		CarId:   car.ID.Hex(),
		Message: "Car created successfully",
	}, nil
}

func (s *InventoryServer) ListCars(ctx context.Context, req *pb.ListCarsRequest) (*pb.ListCarsResponse, error) {
	// Use empty filter to get all cars
	cars, err := s.service.SearchCars("", "", 0, false)
	if err != nil {
		return nil, err
	}

	var pbCars []*pb.Car
	for _, car := range cars {
		pbCars = append(pbCars, &pb.Car{
			CarId:   car.ID.Hex(),
			OwnerId: "", // Not in our model
			Make:    car.Make,
			Model:   car.Model,
			Year:    car.Year,
		})
	}

	return &pb.ListCarsResponse{Cars: pbCars}, nil
}

// Helper function to convert model.Car to pb.Car
func modelCarToPbCar(car *model.Car) *pb.Car {
	return &pb.Car{
		CarId:   car.ID.Hex(),
		OwnerId: "", // Not in our model
		Make:    car.Make,
		Model:   car.Model,
		Year:    car.Year,
	}
}
