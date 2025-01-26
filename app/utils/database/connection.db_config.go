package database

import (
	"database/sql"
	"fmt"
	"log"
	"triesdi/app/configs/db_config"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDatabase() {
	var err error
	// Replace with your PostgreSQL connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", db_config.DB_HOST, db_config.DB_PORT, db_config.DB_USER, db_config.DB_PASSWORD, db_config.DB_NAME)

	// Open database connection
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Test the connection
	err = DB.Ping()
	if err != nil {
		log.Fatal("Database connection test failed:", err)
	}

	log.Println("Database connected successfully!")
}
