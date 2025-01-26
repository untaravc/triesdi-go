package commands

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"triesdi/app/configs/db_config"
)

func RunMigrationMysql() {
	// Connect to the database
	dsnMysql := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_config.DB_USER, db_config.DB_PASSWORD, db_config.DB_HOST, db_config.DB_PORT, db_config.DB_NAME)
	db, err := sql.Open("mysql", dsnMysql)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	// Path to migrations folder
	migrationsPath := "./app/migrations"

	// // Fetch migration files
	files, err := os.ReadDir(migrationsPath)
	if err != nil {
		log.Fatalf("Error reading migrations directory: %v", err)
	}

	// Sort files by name
	var migrationFiles []string
	for _, file := range files {
		migrationFiles = append(migrationFiles, file.Name())
	}
	sort.Strings(migrationFiles)

	// Execute migrations
	for _, file := range migrationFiles {
		filePath := filepath.Join(migrationsPath, file)
		content, err := os.ReadFile(filePath)
		if err != nil {
			log.Fatalf("Error reading migration file %s: %v", file, err)
		}

		fmt.Printf("Running migration: %s\n", file)
		if _, err := db.Exec(string(content)); err != nil {
			log.Fatalf("Error executing migration %s: %v", file, err)
		}
		fmt.Printf("Migration %s applied successfully.\n", file)
	}
}

func RunMigrationPostgres() {
	// Connect to the database
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", db_config.DB_HOST, db_config.DB_PORT, db_config.DB_USER, db_config.DB_PASSWORD, db_config.DB_NAME)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	// Path to migrations folder
	migrationsPath := "./app/migrations"

	// // Fetch migration files
	files, err := os.ReadDir(migrationsPath)
	if err != nil {
		log.Fatalf("Error reading migrations directory: %v", err)
	}

	// Sort files by name
	var migrationFiles []string
	for _, file := range files {
		migrationFiles = append(migrationFiles, file.Name())
	}
	sort.Strings(migrationFiles)

	// Execute migrations
	for _, file := range migrationFiles {
		filePath := filepath.Join(migrationsPath, file)
		content, err := os.ReadFile(filePath)
		if err != nil {
			log.Fatalf("Error reading migration file %s: %v", file, err)
		}

		fmt.Printf("Running migration: %s\n", file)
		if _, err := db.Exec(string(content)); err != nil {
			log.Fatalf("Error executing migration %s: %v", file, err)
		}
		fmt.Printf("Migration %s applied successfully.\n", file)
	}
}
