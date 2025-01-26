package routes

import (
	"triesdi/app/controllers/auth_controller"
	"triesdi/app/controllers/upload_controller"
	"triesdi/app/controllers/user_controller"
	"triesdi/app/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {
	route := app

	route.POST("/v1/login/email", auth_controller.LoginEmail)
	route.POST("/v1/login/phone", auth_controller.LoginPhone)
	route.POST("/v1/register/email", auth_controller.RegisterEmail)
	route.POST("/v1/register/phone", auth_controller.RegisterPhone)

	JWTMiddleware := middleware.JWTMiddleware()
	// Middleware
	route.Use(JWTMiddleware)
	route.GET("/v1/user", user_controller.Auth)
	route.PUT("/v1/user", user_controller.Update)
	route.POST("/v1/user/link/phone", user_controller.LinkPhone)
	route.POST("/v1/user/link/email", user_controller.LinkEmail)

	route.POST("/v1/file", upload_controller.AddFile)
}
