CREATE TABLE IF NOT EXISTS order_statistics (
    user_id VARCHAR(255) PRIMARY KEY,
    total_orders INT,
    most_active_time VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS user_statistics (
    id SERIAL PRIMARY KEY,
    total_users INT,
    active_users INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);