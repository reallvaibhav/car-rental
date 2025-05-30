package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Car-Rental/backend/user-service/internal/models"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	if db == nil {
		log.Println("Warning: Running UserRepository in mock mode")
		return &UserRepository{collection: nil}
	}
	return &UserRepository{collection: db.Collection("users")}
}

func (r *UserRepository) Create(user *models.User) error {
	if r.collection == nil {
		log.Printf("Mock: Creating user %s", user.Email)
		// In mock mode, check for duplicate emails
		existingUser, _ := r.GetByEmail(user.Email)
		if existingUser != nil && existingUser.Email == user.Email {
			log.Printf("Mock: Email already exists: %s", user.Email)
			return fmt.Errorf("email already exists: %s", user.Email)
		}

		// In mock mode, check for duplicate usernames
		existingUserByName, _ := r.GetByName(user.Name)
		if existingUserByName != nil {
			log.Printf("Mock: Username already exists: %s", user.Name)
			return fmt.Errorf("username already exists: %s", user.Name)
		}

		user.ID = primitive.NewObjectID()
		user.CreatedAt = time.Now()
		return nil
	}
	// Real MongoDB implementation
	// First check if user already exists with the same email
	var existingUserByEmail models.User
	err := r.collection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&existingUserByEmail)
	if err == nil {
		// User with this email already exists
		log.Printf("MongoDB: Email already exists: %s", user.Email)
		return fmt.Errorf("email already exists: %s", user.Email)
	} else if err != mongo.ErrNoDocuments {
		// Some other error occurred
		log.Printf("MongoDB error checking for existing user: %v", err)
		return err
	}

	// Check if user already exists with the same name
	var existingUserByName models.User
	err = r.collection.FindOne(context.TODO(), bson.M{"name": user.Name}).Decode(&existingUserByName)
	if err == nil {
		// User with this name already exists
		log.Printf("MongoDB: Username already exists: %s", user.Name)
		return fmt.Errorf("username already exists: %s", user.Name)
	} else if err != mongo.ErrNoDocuments {
		// Some other error occurred
		log.Printf("MongoDB error checking for existing username: %v", err)
		return err
	}

	// No duplicate found, proceed with creation
	user.CreatedAt = time.Now()
	_, err = r.collection.InsertOne(context.TODO(), user)
	return err
}

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	if r.collection == nil {
		log.Printf("Mock: Getting user by email %s", email)

		// Check if the email is empty
		if email == "" {
			return nil, mongo.ErrNoDocuments
		}

		// For specific test users in mock mode
		// This ensures we can test login with consistent data
		mockUsers := map[string]*models.User{
			"john@example.com": {
				ID:        primitive.NewObjectID(),
				Email:     "john@example.com",
				Password:  "$2a$10$zomlPsSFVvSXnA9zVcxv1eHeSBbGW/5xzHvp.2feEsW4ApAglwitG", // hashed "password123"
				Name:      "John Doe",
				CreatedAt: time.Now(),
			},
			"test123@example.com": {
				ID:        primitive.NewObjectID(),
				Email:     "test123@example.com",
				Password:  "$2a$10$zomlPsSFVvSXnA9zVcxv1eHeSBbGW/5xzHvp.2feEsW4ApAglwitG", // hashed "password123"
				Name:      "Test User",
				CreatedAt: time.Now(),
			},
		}

		// Look up the user by email
		if user, exists := mockUsers[email]; exists {
			return user, nil
		}

		// If email not in our mock database, return a MongoDB-like error
		return nil, mongo.ErrNoDocuments
	}

	// Real MongoDB implementation
	var user models.User
	err := r.collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	return &user, err
}

func (r *UserRepository) GetByID(id string) (*models.User, error) {
	if r.collection == nil {
		log.Printf("Mock: Getting user by ID %s", id)
		return &models.User{
			ID:        primitive.NewObjectID(),
			Email:     "mock@example.com",
			Name:      "Mock User",
			CreatedAt: time.Now(),
		}, nil
	}
	oid, _ := primitive.ObjectIDFromHex(id)
	var user models.User
	err := r.collection.FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&user)
	return &user, err
}

// GetByName retrieves a user by their name
func (r *UserRepository) GetByName(name string) (*models.User, error) {
	if r.collection == nil {
		log.Printf("Mock: Getting user by name %s", name)

		// Check if the name is empty
		if name == "" {
			return nil, mongo.ErrNoDocuments
		}

		// For specific test users in mock mode
		mockUsers := map[string]*models.User{
			"John Doe": {
				ID:        primitive.NewObjectID(),
				Email:     "john@example.com",
				Password:  "$2a$10$zomlPsSFVvSXnA9zVcxv1eHeSBbGW/5xzHvp.2feEsW4ApAglwitG", // hashed "password123"
				Name:      "John Doe",
				CreatedAt: time.Now(),
			},
			"Test User": {
				ID:        primitive.NewObjectID(),
				Email:     "test123@example.com",
				Password:  "$2a$10$zomlPsSFVvSXnA9zVcxv1eHeSBbGW/5xzHvp.2feEsW4ApAglwitG", // hashed "password123"
				Name:      "Test User",
				CreatedAt: time.Now(),
			},
		}

		// Look up the user by name
		if user, exists := mockUsers[name]; exists {
			return user, nil
		}

		// If name not in our mock database, return a MongoDB-like error
		return nil, mongo.ErrNoDocuments
	}

	// Real MongoDB implementation
	var user models.User
	err := r.collection.FindOne(context.TODO(), bson.M{"name": name}).Decode(&user)
	return &user, err
}
