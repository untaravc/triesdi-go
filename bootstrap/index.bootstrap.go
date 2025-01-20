package bootstrap

import (
	"fmt"
	"log"
	"os"
	"triesdi/app/cache"
	"triesdi/app/commands"
	"triesdi/app/configs"
	"triesdi/app/configs/app_config"
	"triesdi/app/configs/cors_config"

	"triesdi/app/configs/db_config"
	"triesdi/app/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func loadEnv() {
	env := os.Getenv("ENV")
	var err error

	switch env {
	case "production":
		err = godotenv.Load(".env.production")
	case "development":
		fmt.Println("Loading development environment variables")
		err = godotenv.Load(".env.development")
	case "local":
		fmt.Println("Loading local environment variables")
		err = godotenv.Load(".env.local")
	default:
		fmt.Println("Loading default environment variables")
		err = godotenv.Load()
	}

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func BootstrapApp() {

	// Initialize cache
	cache.InitializeCacheActivityTypes()

	loadEnv();

	configs.InitConfig()

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "migrate":
			log.Println("Running migrations...")
			commands.RunMigration()
			return
		}
	}

	if app_config.GIN_MODE == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	app := gin.Default()

	app.Use(cors_config.CORSMiddleware())
	// app.Use(log_config.LoggerMiddleware())

	db_config.ConnectDatabase()
	// db_config.InitRedisClient()

	routes.InitRoute(app)
	routes.InitApiRoute(app)

	if err := app.Run(app_config.PORT); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
