package bootstrap

import (
	// "context"
	"log"
	"os"
	"triesdi/app/cache"
	"triesdi/app/commands"
	"triesdi/app/configs"
	"triesdi/app/configs/app_config"

	// "triesdi/app/configs/aws_config"
	"triesdi/app/configs/cors_config"
	// "triesdi/app/controllers/upload_controller"
	// "triesdi/app/services/upload_service"
	"triesdi/app/utils/database"

	"triesdi/app/configs/log_config"
	"triesdi/app/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	// "github.com/aws/aws-sdk-go-v2/config"
	// "github.com/aws/aws-sdk-go-v2/credentials"
	// "github.com/aws/aws-sdk-go-v2/service/s3"
)

func BootstrapApp() {

	// Initialize cache
	cache.InitializeCacheActivityTypes()

	err := godotenv.Load()

	if err != nil {
		log.Println("Error loading .env file")
	}

	configs.InitConfig()

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "migrate":
			log.Println("Running migrations...")
			commands.RunMigrationPostgres()
			return
		}
	}

	if app_config.GIN_MODE == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	app := gin.Default()

	app.Use(cors_config.CORSMiddleware())
	app.Use(log_config.LoggerMiddleware())

	database.InitDatabase()
	defer database.DB.Close()

	// Load AWS configuration
	// cfg, err := config.LoadDefaultConfig(context.TODO(),
	// 	config.WithRegion(aws_config.AWS_REGION),
	// 	config.WithCredentialsProvider(
	// 		credentials.NewStaticCredentialsProvider(aws_config.AWS_ACCESS_KEY_ID, aws_config.AWS_SECRET_ACCESS_KEY, ""),
	// 	),
	// )
	// if err != nil {
	// 	log.Fatalf("unable to load AWS SDK config, %v", err)
	// }

	// Create an S3 client
	// s3Client := s3.NewFromConfig(cfg)
	// s3Uploader := upload_service.NewS3Uploader(s3Client)

	routes.InitRoute(app)

	// app.POST("/v1/file", upload_controller.UploadHandler(s3Uploader))

	if err := app.Run(app_config.PORT); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
