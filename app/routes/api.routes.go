package routes

import (
	api_department_controller "triesdi/app/controllers/api/department_controller"
	api_home_controller "triesdi/app/controllers/api/home_controller"
	"triesdi/app/controllers/v1/v1_upload_controller"

	"github.com/gin-gonic/gin"
)

func InitApiRoute(app *gin.Engine) {
	route := app

	api := route.Group("/api")

	api.POST("/department", api_department_controller.CreateDepartment)
	route.POST("/v1/upload", v1_upload_controller.UploadImage)
	api.GET("/", api_home_controller.Index)
}
