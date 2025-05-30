package repository

import (
	"context"
	"database/sql"
	"log"
	"time"

	"statistics-service/internal/usecase"
)

type Repository struct {
	db *sql.DB
}

type Statistics struct {
	OrderCount   int       `json:"order_count"`
	TotalRevenue float64   `json:"total_revenue"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (r *Repository) GetItemStatistics() (*usecase.ItemStatistics, error) {
	stats := &usecase.ItemStatistics{}

	// Get total items and value
	err := r.db.QueryRow(`
		SELECT 
			COUNT(*) as total_items,
			COALESCE(SUM(price * stock), 0) as total_value,
			COALESCE(AVG(price), 0) as avg_price,
			COUNT(CASE WHEN stock < 10 THEN 1 END) as low_stock,
			COUNT(CASE WHEN stock = 0 THEN 1 END) as out_of_stock
		FROM products
	`).Scan(
		&stats.TotalItems,
		&stats.TotalValue,
		&stats.AvgPrice,
		&stats.LowStockItems,
		&stats.OutOfStockItems,
	)

	if err != nil {
		log.Printf("Error querying item statistics: %v", err)
		return nil, err
	}

	return stats, nil
}

func (r *Repository) GetOrderStatistics() (*usecase.OrderStatistics, error) {
	stats := &usecase.OrderStatistics{}

	// Get order statistics
	err := r.db.QueryRow(`
		SELECT 
			COUNT(*) as total_orders,
			COALESCE(SUM(total), 0) as total_revenue,
			COALESCE(AVG(total), 0) as avg_order_value,
			COUNT(CASE WHEN status = 'Pending' THEN 1 END) as pending_orders,
			COUNT(CASE WHEN status = 'Completed' THEN 1 END) as completed_orders
		FROM orders
	`).Scan(
		&stats.TotalOrders,
		&stats.TotalRevenue,
		&stats.AvgOrderValue,
		&stats.PendingOrders,
		&stats.CompletedOrders,
	)

	if err != nil {
		log.Printf("Error querying order statistics: %v", err)
		return nil, err
	}

	return stats, nil
}

func (r *Repository) GetStatistics(ctx context.Context) (*Statistics, error) {
	stats := &Statistics{}
	err := r.db.QueryRowContext(ctx, `
		SELECT order_count, total_revenue, updated_at 
		FROM statistics 
		ORDER BY updated_at DESC 
		LIMIT 1
	`).Scan(&stats.OrderCount, &stats.TotalRevenue, &stats.UpdatedAt)
	return stats, err
}

func (r *Repository) UpdateStatistics(ctx context.Context, stats *Statistics) error {
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO statistics (order_count, total_revenue, updated_at)
		VALUES ($1, $2, $3)
	`, stats.OrderCount, stats.TotalRevenue, time.Now())
	return err
}
