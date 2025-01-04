package api_home_controller

import (
	"net/http"
	response "triesdi/app/responses"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	response.BaseResponse(ctx, http.StatusOK, true, "OK", "Hello World")
}
