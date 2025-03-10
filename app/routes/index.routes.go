package routes

import (
	"triesdi/app/controllers/v1/v1_auth_controller"
	"triesdi/app/controllers/v1/v1_upload_controller"
	"triesdi/app/controllers/v1/v1_user_controller"
	"triesdi/app/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {
	route := app

	JWTMiddleware := middleware.JWTMiddleware()

	route.POST("/v1/login", v1_auth_controller.Login)
	route.POST("/v1/register", v1_auth_controller.Register)

	// Middleware
	route.Use(JWTMiddleware)

	// User
	route.GET("/v1/user", v1_user_controller.GetUser)
	route.PATCH("/v1/user", v1_user_controller.UpdateUser)

	// Image
	route.POST("/v1/file", v1_upload_controller.UploadImage)
}
