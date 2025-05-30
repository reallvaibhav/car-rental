package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Car-Rental/backend/user-service/cache"
	"github.com/Car-Rental/backend/user-service/internal/auth"
	"github.com/Car-Rental/backend/user-service/internal/models"
	"github.com/Car-Rental/backend/user-service/internal/publisher"
	"github.com/Car-Rental/backend/user-service/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo       *repository.UserRepository
	jwt        *auth.JWTManager
	publisher  *publisher.NatsPublisher
	cache      *cache.Cache
	redisCache *cache.RedisCache
}

func New(repo *repository.UserRepository, jwt *auth.JWTManager, pub *publisher.NatsPublisher, cache *cache.Cache, redisCache *cache.RedisCache) *UserService {
	return &UserService{
		repo:       repo,
		jwt:        jwt,
		publisher:  pub,
		cache:      cache,
		redisCache: redisCache,
	}
}

func (s *UserService) Register(email, password, name string) (string, error) {
	log.Printf("Attempting to register user with email: %s", email)
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := &models.User{Email: email, Password: string(hash), Name: name}
	if err := s.repo.Create(user); err != nil {
		log.Printf("Registration failed: %v", err)
		return "", err
	}

	// Cache the new user
	cacheKey := fmt.Sprintf("user:%s", user.ID.Hex())
	ctx := context.Background()
	s.cache.Set(cacheKey, user, 24*time.Hour)
	if s.redisCache != nil {
		s.redisCache.Set(ctx, cacheKey, user, 24*time.Hour)
	}
	log.Printf("Cached new registered user: %s", user.ID.Hex())

	// 🎯 Publish user.created event to NATS
	s.publisher.Publish("user.created", map[string]interface{}{
		"user_id": user.ID.Hex(),
		"email":   user.Email,
		"name":    user.Name,
	})

	return s.jwt.Generate(user.ID.Hex())
}

func (s *UserService) Login(email, password string) (string, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return "", err
	}

	log.Printf("Stored password hash: %s", user.Password)
	log.Printf("Input password: %s", password)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		log.Printf("Password validation failed: %v", err)
		return "", err
	}

	// Cache the user after successful login
	cacheKey := fmt.Sprintf("user:%s", user.ID.Hex())
	ctx := context.Background()
	s.cache.Set(cacheKey, user, 24*time.Hour)
	if s.redisCache != nil {
		s.redisCache.Set(ctx, cacheKey, user, 24*time.Hour)
	}
	log.Printf("Cached user after login: %s", user.ID.Hex())

	return s.jwt.Generate(user.ID.Hex())
}

func (s *UserService) GetUserByID(id string) (*models.User, error) {
	// Try to get from cache first
	cacheKey := fmt.Sprintf("user:%s", id)
	ctx := context.Background()

	// Try Redis cache first if available
	if s.redisCache != nil {
		var cachedUser models.User
		if err := s.redisCache.Get(ctx, cacheKey, &cachedUser); err == nil {
			log.Printf("Cache hit (Redis) for user: %s", id)
			return &cachedUser, nil
		}
	}

	// Try in-memory cache
	if cached, found := s.cache.Get(cacheKey); found {
		if user, ok := cached.(*models.User); ok {
			log.Printf("Cache hit (memory) for user: %s", id)
			return user, nil
		}
	}

	// If not in cache, get from database
	user, err := s.repo.GetByID(id)
	if err != nil {
		log.Printf("Error getting user: %v", err)
		return nil, err
	}

	// Cache the result
	s.cache.Set(cacheKey, user, 24*time.Hour) // Cache user data for 24 hours
	if s.redisCache != nil {
		s.redisCache.Set(ctx, cacheKey, user, 24*time.Hour)
	}
	log.Printf("Cached user: %s", id)

	return user, nil
}

func (s *UserService) ValidateToken(token string) (string, error) {
	// Try to get from cache first
	cacheKey := fmt.Sprintf("token:%s", token)
	ctx := context.Background()

	// Try Redis cache first if available
	if s.redisCache != nil {
		var cachedUserID string
		if err := s.redisCache.Get(ctx, cacheKey, &cachedUserID); err == nil {
			log.Printf("Cache hit (Redis) for token validation")
			return cachedUserID, nil
		}
	}

	// Try in-memory cache
	if cached, found := s.cache.Get(cacheKey); found {
		if userID, ok := cached.(string); ok {
			log.Printf("Cache hit (memory) for token validation")
			return userID, nil
		}
	}

	// If not in cache, validate with JWT
	userID, err := s.jwt.Validate(token)
	if err != nil {
		return "", err
	}

	// Cache the result for shorter duration (tokens should be validated more frequently)
	s.cache.Set(cacheKey, userID, 30*time.Minute)
	if s.redisCache != nil {
		s.redisCache.Set(ctx, cacheKey, userID, 30*time.Minute)
	}
	log.Printf("Cached token validation result")

	return userID, nil
}
