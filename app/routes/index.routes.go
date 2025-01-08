package routes

import (
	"triesdi/app/controllers/v1/v1_auth_controller"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {
	route := app

	route.POST("/v1/auth", v1_auth_controller.Auth)
}
