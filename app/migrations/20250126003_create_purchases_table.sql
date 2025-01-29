DROP TABLE IF EXISTS purchases CASCADE;
CREATE TABLE purchases (
    purchase_id SERIAL PRIMARY KEY,
    sender_name VARCHAR(255) NOT NULL,
    sender_contact_type VARCHAR(255) NOT NULL,
    sender_contact_detail VARCHAR(255) NOT NULL,
    total_price INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_purchases_purchase_id ON purchases (purchase_id);