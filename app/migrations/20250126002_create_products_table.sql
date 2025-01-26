DROP TABLE IF EXISTS products;
CREATE TABLE products (
    product_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    category VARCHAR(255) NOT NULL,
    qty INT NOT NULL,
    price INT NOT NULL,
    sku VARCHAR(255) NOT NULL,
    file_id VARCHAR(255),
    file_uri VARCHAR(255),
    file_thumbnail_uri VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_products_category ON products (category);
CREATE INDEX idx_products_sku ON products (sku);