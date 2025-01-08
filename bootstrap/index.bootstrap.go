package bootstrap

import (
	// "gin-gorm/configs"
	// "gin-gorm/configs/app_config"
	// "gin-gorm/configs/cors_config"

	// "gin-gorm/app/middleware"
	// "gin-gorm/configs/log_config"
	// "gin-gorm/database"
	// "gin-gorm/routes"
	"fmt"
	"log"
	"triesdi/app/configs"
	"triesdi/app/configs/app_config"
	"triesdi/app/configs/cors_config"

	"triesdi/app/configs/db_config"
	"triesdi/app/configs/log_config"
	"triesdi/app/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func BootstrapApp() {
	err := godotenv.Load()

	if err != nil {
		log.Println("Error loading .env file")
	}

	configs.InitConfig()

	if app_config.GIN_MODE == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	app := gin.Default()

	app.Use(cors_config.CORSMiddleware())
	app.Use(log_config.LoggerMiddleware())

	// db_config.InitRedisClient()
	db_config.ConnectDatabase()

	routes.InitRoute(app)
	routes.InitApiRoute(app)
	
	fmt.Printf("Server is starting on port %s...\n", app_config.PORT)
	if err := app.Run(app_config.PORT); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
