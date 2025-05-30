package repository

import (
	"context"
	"time"

	"github.com/Car-Rental/backend/inventory-service/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CarRepository struct {
	collection *mongo.Collection
}

func NewCarRepository(db *mongo.Database) *CarRepository {
	if db == nil {
		return &CarRepository{collection: nil}
	}
	return &CarRepository{collection: db.Collection("cars")}
}

func (r *CarRepository) Create(car *model.Car) error {
	if r.collection == nil {
		return nil // Mock mode
	}
	car.CreatedAt = time.Now()
	car.Available = true
	result, err := r.collection.InsertOne(context.TODO(), car)
	if err != nil {
		return err
	}
	car.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *CarRepository) GetByID(id string) (*model.Car, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var car model.Car
	err = r.collection.FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&car)
	return &car, err
}

func (r *CarRepository) Update(car *model.Car) error {
	_, err := r.collection.ReplaceOne(
		context.TODO(),
		bson.M{"_id": car.ID},
		car,
	)
	return err
}

func (r *CarRepository) Delete(id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.collection.DeleteOne(context.TODO(), bson.M{"_id": oid})
	return err
}

func (r *CarRepository) Search(filter bson.M) ([]*model.Car, error) {
	cursor, err := r.collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var cars []*model.Car
	if err := cursor.All(context.TODO(), &cars); err != nil {
		return nil, err
	}
	return cars, nil
}

func (r *CarRepository) UpdateAvailability(id string, available bool) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": oid},
		bson.M{"$set": bson.M{"available": available}},
	)
	return err
}
