package db_config

import "os"

var DB_DRIVER = ""
var DB_HOST = ""
var DB_PORT = ""
var DB_NAME = ""
var DB_USER = ""
var DB_PASSWORD = ""

func InitDatabaseConfig() {
	portDbDriver := os.Getenv("DB_DRIVER")
	portDbHost := os.Getenv("DB_HOST")
	portDbPort := os.Getenv("DB_PORT")
	portDbName := os.Getenv("DB_NAME")
	portDbUser := os.Getenv("DB_USER")
	portDbPassword := os.Getenv("DB_PASSWORD")

	if portDbDriver != "" {
		DB_DRIVER = portDbDriver
	}
	if portDbHost != "" {
		DB_HOST = portDbHost
	}
	if portDbPort != "" {
		DB_PORT = portDbPort
	}
	if portDbName != "" {
		DB_NAME = portDbName
	}
	if portDbUser != "" {
		DB_USER = portDbUser
	}
	if portDbPassword != "" {
		DB_PASSWORD = portDbPassword
	}
}
