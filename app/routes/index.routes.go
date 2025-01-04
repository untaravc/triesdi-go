package routes

import (
	api_home_controller "triesdi/app/controllers/api/home_controller"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {
	route := app
	route.GET("/", api_home_controller.Index)
}
