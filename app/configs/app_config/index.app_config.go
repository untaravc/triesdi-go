package app_config

import (
	"os"
	"strconv"
)

var PORT = ":8000"
var STATIC_ROUTE = "/public"
var STATIC_DIR = "./public"
var GIN_MODE = "debug"
var TIMEZONE = 0
var REDIS_EXPIRE = 600

func InitAppConfig() {
	portEnv := os.Getenv("APP_PORT")
	staticRouteEnv := os.Getenv("STATIC_ROUTE")
	staticDirEnv := os.Getenv("STATIC_DIR")
	ginModeEnv := os.Getenv("GIN_MODE")
	timezone := os.Getenv("TIMEZONE")
	redisExpire := os.Getenv("REDIS_EXPIRE")

	if portEnv != "" {
		PORT = portEnv
	}

	if staticRouteEnv != "" {
		STATIC_ROUTE = staticRouteEnv
	}

	if staticDirEnv != "" {
		STATIC_DIR = staticDirEnv
	}

	if ginModeEnv != "" {
		GIN_MODE = ginModeEnv
	}

	if timezone != "" {
		TIMEZONE, _ = strconv.Atoi(timezone)
	}

	if redisExpire != "" {
		REDIS_EXPIRE, _ = strconv.Atoi(redisExpire)
	}
}
