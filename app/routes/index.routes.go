package routes

import (
	"triesdi/app/controllers/v1/v1_auth_controller"
	"triesdi/app/controllers/v1/v1_upload_controller"
	"triesdi/app/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {
	route := app

	JWTMiddleware := middleware.JWTMiddleware()

	route.POST("/v1/auth", v1_auth_controller.AuthNew)

	// Middleware
	route.Use(JWTMiddleware)

	// Image
	route.POST("/v1/file", v1_upload_controller.UploadImage)
}
