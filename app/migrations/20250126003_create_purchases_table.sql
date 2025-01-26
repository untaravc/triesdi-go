DROP TABLE IF EXISTS purchases;
CREATE TABLE purchases (
    purchase_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    sender_name VARCHAR(255) NOT NULL,
    sender_contact_type VARCHAR(255) NOT NULL,
    sender_contact_detail VARCHAR(255) NOT NULL,
    total_price INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);