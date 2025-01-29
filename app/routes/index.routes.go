package routes

import (
	"triesdi/app/controllers/auth_controller"
	"triesdi/app/controllers/product_controller"
	"triesdi/app/controllers/purchase_controller"
	"triesdi/app/controllers/upload_controller"
	"triesdi/app/controllers/user_controller"
	"triesdi/app/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {
	route := app

	// Auth
	route.POST("/v1/login/email", auth_controller.LoginEmail)
	route.POST("/v1/login/phone", auth_controller.LoginPhone)
	route.POST("/v1/register/email", auth_controller.RegisterEmail)
	route.POST("/v1/register/phone", auth_controller.RegisterPhone)

	route.GET("/v1/product", product_controller.GetAll)
	route.POST("/v1/purchase", purchase_controller.Create)
	route.POST("/v1/purchase/:purchaseId", purchase_controller.Update)

	// Assign Middleware
	JWTMiddleware := middleware.JWTMiddleware()
	route.Use(JWTMiddleware)

	// User
	route.GET("/v1/user", user_controller.Auth)
	route.PUT("/v1/user", user_controller.Update)
	route.POST("/v1/user/link/phone", user_controller.LinkPhone)
	route.POST("/v1/user/link/email", user_controller.LinkEmail)

	// Upload
	route.POST("/v1/file", upload_controller.AddFile)

	// Product
	route.POST("/v1/product", product_controller.Store)
	route.PUT("/v1/product/:product_id", product_controller.Update)
	route.DELETE("/v1/product/:product_id", product_controller.Delete)
}
