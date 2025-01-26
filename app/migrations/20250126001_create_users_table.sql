DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE,
    phone VARCHAR(20) UNIQUE,
    password VARCHAR(255) NOT NULL,
    token VARCHAR(255),
    file_id VARCHAR(255),
    bank_account_name VARCHAR(255),
    bank_account_holder VARCHAR(255),
    bank_account_number VARCHAR(255),
    file_uri VARCHAR(255),
    file_thumbnail_uri VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_users_email ON users (email);
CREATE INDEX idx_users_phone ON users (phone);