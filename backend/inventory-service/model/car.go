package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Car struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Make        string             `bson:"make"`
	Model       string             `bson:"model"`
	Year        int32              `bson:"year"`
	Category    string             `bson:"category"`
	Location    string             `bson:"location"`
	Available   bool               `bson:"available"`
	PricePerDay float64            `bson:"price_per_day"`
	Features    []string           `bson:"features"`
	CreatedAt   time.Time          `bson:"created_at"`
}
