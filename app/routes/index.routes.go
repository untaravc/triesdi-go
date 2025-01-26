package routes

import (
	"triesdi/app/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {
	route := app

	JWTMiddleware := middleware.JWTMiddleware()
	// Middleware
	route.Use(JWTMiddleware)
}
