package repository

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	SaveOrderStatistics(userID string, totalOrders int, mostActiveTime string) error
	GetUserOrderStatistics(userID string) (int32, string, error)
	GetUserStatistics() (int32, int32, error)
}

type repository struct {
	db          *sql.DB
	mongoClient *mongo.Client
}

func NewRepository(postgresURI, mongoURI string) (*Repository, error) {
	// Initialize PostgreSQL
	postgresDB, err := sql.Open("postgres", postgresURI)
	if err != nil {
		return nil, err
	}

	// Initialize MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}

	return &repository{
		db:          postgresDB,
		mongoClient: mongoClient,
	}, nil
}

func (r *repository) SaveOrderStatistics(userID string, totalOrders int, mostActiveTime string) error {
	query := `
        INSERT INTO order_statistics (user_id, total_orders, most_active_time)
        VALUES ($1, $2, $3)
        ON CONFLICT (user_id) DO UPDATE
        SET total_orders = $2, most_active_time = $3
    `
	_, err := r.db.Exec(query, userID, totalOrders, mostActiveTime)
	return err
}

func (r *repository) GetUserOrderStatistics(userID string) (int32, string, error) {
	query := `
        SELECT total_orders, most_active_time 
        FROM order_statistics 
        WHERE user_id = $1
    `
	var totalOrders int32
	var mostActiveTime string
	err := r.db.QueryRow(query, userID).Scan(&totalOrders, &mostActiveTime)
	return totalOrders, mostActiveTime, err
}

func (r *repository) GetUserStatistics() (int32, int32, error) {
	query := `
        SELECT total_users, active_users 
        FROM user_statistics 
        ORDER BY created_at DESC 
        LIMIT 1
    `
	var totalUsers, activeUsers int32
	err := r.db.QueryRow(query).Scan(&totalUsers, &activeUsers)
	return totalUsers, activeUsers, err
}
