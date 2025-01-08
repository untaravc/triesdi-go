package v1_auth_controller

import (
	"net/http"
	"triesdi/app/requests/auth_request"
	"triesdi/app/responses/response"
	"triesdi/app/validator"

	"github.com/gin-gonic/gin"
)

func Auth(ctx *gin.Context) {
	// SHouldBindJSON
	auth_request := new(auth_request.AuthRequest)
	err_req := ctx.ShouldBind(&auth_request)
	if err_req != nil {
		response.BaseResponse(ctx, http.StatusBadRequest, false, "Bad request", nil)
		return
	}

	// Validate the struct
	if err := validator.ValidateStruct(auth_request); err != nil {
		response.BaseResponse(ctx, http.StatusBadRequest, false, err.Error(), nil)
		return
	}

	response.BaseResponse(ctx, http.StatusOK, true, "OK", nil)
}
