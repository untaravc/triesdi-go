CREATE TABLE users (
    id CHAR(36) PRIMARY KEY,
    preference ENUM('CARDIO', 'WEIGHT') DEFAULT NULL,
    weight_unit ENUM('KG', 'LBS') DEFAULT NULL,
    height_unit ENUM('CM', 'INCH') DEFAULT NULL,
    weight INT DEFAULT NULL,
    height INT DEFAULT NULL,
    name VARCHAR(60) DEFAULT NULL,
    image_uri TEXT DEFAULT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
);
