package models

import (
  "go.mongodb.org/mongo-driver/bson/primitive"
  "time"
)

type User struct {
  ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
  Email     string             `bson:"email"`
  Password  string             `bson:"password"`
  Name      string             `bson:"name"`
  CreatedAt time.Time          `bson:"created_at"`
}
