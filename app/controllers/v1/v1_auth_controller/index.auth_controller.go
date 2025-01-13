package v1_auth_controller

import (
	"net/http"
	"triesdi/app/configs/db_config"
	"triesdi/app/repository/auth_repository"
	"triesdi/app/requests/auth_request"
	"triesdi/app/responses/response"
	"triesdi/app/service/auth_service"
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


type AuthRequest struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Action string `json:"action" binding:"required,oneof='login' 'create'"`
}

func AuthNew(ctx *gin.Context) {
	authRequest := new(AuthRequest)
	if err := ctx.ShouldBindJSON(authRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	if err := validator.ValidateStruct(authRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check Action Login or Register
	db := db_config.GetDB()
	authRepository := auth_repository.NewRepository(db)
	authService := auth_service.NewService(authRepository)

	var authResponse response.AuthResponse
	var err error
	if authRequest.Action == "login" {
		authResponse, err = authService.Login(authRequest.Email, authRequest.Password)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
	} else if authRequest.Action == "create" {
		authResponse, err = authService.CreateUser(authRequest.Email, authRequest.Password)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	// Return Auth Response
	ctx.JSON(http.StatusOK, authResponse)
}
