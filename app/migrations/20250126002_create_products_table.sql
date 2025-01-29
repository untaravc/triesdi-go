DROP TABLE IF EXISTS products CASCADE;
CREATE TABLE products (
    product_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
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
-- CREATE INDEX idx_products_category ON products (category);
-- CREATE INDEX idx_products_sku ON products (sku);
CREATE INDEX idx_products_user_id ON products (user_id);
-- foreign key constraint
ALTER TABLE products
ADD CONSTRAINT fk_products_user_id FOREIGN KEY(user_id) REFERENCES users(id);