-- Initialize car rental database schema
CREATE TABLE IF NOT EXISTS cars (
    id TEXT PRIMARY KEY,
    brand TEXT NOT NULL,
    model TEXT NOT NULL,
    year TEXT NOT NULL,
    license_plate TEXT UNIQUE NOT NULL,
    daily_rate REAL NOT NULL,
    is_available BOOLEAN DEFAULT true,
    category TEXT NOT NULL,
    transmission TEXT NOT NULL,
    seats INTEGER NOT NULL,
    fuel_type TEXT NOT NULL,
    mileage REAL NOT NULL,
    features TEXT  -- Stored as JSON array
);

-- Create index for faster searches
CREATE INDEX IF NOT EXISTS idx_cars_category ON cars(category);
CREATE INDEX IF NOT EXISTS idx_cars_availability ON cars(is_available);