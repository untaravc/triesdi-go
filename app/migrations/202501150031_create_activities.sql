CREATE TABLE activities (
    id CHAR(36) PRIMARY KEY, 
    user_id CHAR(36) NOT NULL, 
    activity_type_id CHAR(36) NOT NULL,
    activity_type VARCHAR(255) NOT NULL,
    done_at DATETIME NOT NULL,
    duration_in_minutes INT NOT NULL,
    calories_burned INT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_activity_type_id FOREIGN KEY (activity_type_id) REFERENCES activity_types(id) ON DELETE RESTRICT
);
