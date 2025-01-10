package routes

import (
	api_home_controller "triesdi/app/controllers/api/home_controller"

	"github.com/gin-gonic/gin"
)

func InitApiRoute(app *gin.Engine) {
	route := app

	api := route.Group("/api")
	api.GET("/", api_home_controller.Index)
}
