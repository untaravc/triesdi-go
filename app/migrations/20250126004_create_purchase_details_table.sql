DROP TABLE IF EXISTS purchase_details CASCADE;
CREATE TABLE purchase_details (
    purchase_detail_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    purchase_id INT NOT NULL,
    product_id INT NOT NULL,
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
CREATE INDEX idx_purchase_details_user_id ON purchase_details (purchase_id);
-- foreign key constraint
ALTER TABLE purchase_details
ADD CONSTRAINT fk_purchase_details_user_id FOREIGN KEY(user_id) REFERENCES users(id);
ALTER TABLE purchase_details
ADD CONSTRAINT fk_purchase_details_purchase_id FOREIGN KEY(purchase_id) REFERENCES purchases(purchase_id);
ALTER TABLE purchase_details
ADD CONSTRAINT fk_purchase_details_product_id FOREIGN KEY(product_id) REFERENCES products(product_id);