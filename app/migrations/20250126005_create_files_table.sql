DROP TABLE IF EXISTS files;
CREATE TABLE files (
    file_id SERIAL PRIMARY KEY,
    file_uri VARCHAR(255),
    file_thumbnail_uri VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);