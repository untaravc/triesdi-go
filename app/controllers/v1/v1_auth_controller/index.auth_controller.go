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

func Register(ctx *gin.Context) {
	authRequest := new(auth_request.AuthRequest)
	if err := ctx.ShouldBindJSON(authRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	if err := validator.ValidateStruct(authRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	db := db_config.GetDB()
	authRepository := auth_repository.NewRepository(db)
	authService := auth_service.NewService(authRepository)

	var authResponse response.AuthResponse
	var status int
	var err error

	// Create user without status code handling (assuming CreateUser doesn't return a status)
	authResponse, status, err = authService.CreateUser(authRequest.Email, authRequest.Password)
	if err != nil {
		ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}
	status = http.StatusCreated // Set status for successful user creation

	// Return Auth Response
	ctx.JSON(status, authResponse)
}

func Login(ctx *gin.Context) {
	authRequest := new(auth_request.AuthRequest)
	if err := ctx.ShouldBindJSON(authRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	if err := validator.ValidateStruct(authRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := db_config.GetDB()
	authRepository := auth_repository.NewRepository(db)
	authService := auth_service.NewService(authRepository)

	var authResponse response.AuthResponse
	var status int
	var err error

	authResponse, status, err = authService.Login(authRequest.Email, authRequest.Password)
	if err != nil {
		ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}

	// Return Auth Response
	ctx.JSON(status, authResponse)
}
