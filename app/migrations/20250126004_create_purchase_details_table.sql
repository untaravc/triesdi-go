DROP TABLE IF EXISTS purchase_details;
CREATE TABLE purchase_details (
    purchase_detail_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    purchase_id INT NOT NULL,
    product_id INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    category VARCHAR(255) NOT NULL,
    qty INT NOT NULL,
    sku VARCHAR(255) NOT NULL,
    file_id VARCHAR(255),
    file_uri VARCHAR(255),
    file_thumbnail_uri VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);