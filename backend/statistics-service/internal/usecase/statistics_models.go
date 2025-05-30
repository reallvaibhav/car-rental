package usecase

type ItemStatistics struct {
	TotalItems      int     `json:"total_items"`
	TotalValue      float64 `json:"total_value"`
	AvgPrice        float64 `json:"avg_price"`
	LowStockItems   int     `json:"low_stock_items"`
	OutOfStockItems int     `json:"out_of_stock_items"`
}

type OrderStatistics struct {
	TotalOrders     int     `json:"total_orders"`
	TotalRevenue    float64 `json:"total_revenue"`
	AvgOrderValue   float64 `json:"avg_order_value"`
	PendingOrders   int     `json:"pending_orders"`
	CompletedOrders int     `json:"completed_orders"`
}
