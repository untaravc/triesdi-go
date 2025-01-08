package api_home_controller

import (
	"net/http"
	"triesdi/app/responses/response"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	response.BaseResponse(ctx, http.StatusOK, true, "OK", "Test Hello World")
}
