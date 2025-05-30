-- Add the 'price' column to the 'order_items' table
ALTER TABLE order_items ADD COLUMN price REAL NOT NULL;
