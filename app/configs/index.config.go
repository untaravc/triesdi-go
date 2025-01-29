package configs

import (
	"triesdi/app/configs/app_config"
	"triesdi/app/configs/aws_config"
	"triesdi/app/configs/db_config"
)

func InitConfig() {
	app_config.InitAppConfig()
	aws_config.InitAwsConfig()
	// log_config.DefaultLogging()
	db_config.InitDatabaseConfig()
}
