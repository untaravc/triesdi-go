package configs

import (
	"triesdi/app/configs/app_config"
	"triesdi/app/configs/db_config"
	"triesdi/app/configs/log_config"
)

func InitConfig() {
	app_config.InitAppConfig()
	log_config.DefaultLogging()
	db_config.InitDatabaseConfig()
}
