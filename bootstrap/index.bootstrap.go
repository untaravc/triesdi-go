package bootstrap

import (
	"log"
	"os"
	"triesdi/app/commands"
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
	app.Use(log_config.LoggerMiddleware())

	db_config.InitRedisClient()

	routes.InitRoute(app)
	routes.InitApiRoute(app)

	app.Run(app_config.PORT)
}
