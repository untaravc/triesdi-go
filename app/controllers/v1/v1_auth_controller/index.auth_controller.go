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
	Email    string `json:"email" binding:"required,email"`                 // Must be a valid email format
	Password string `json:"password" binding:"required,min=8,max=32"`      // Must be between 8 and 32 characters
	Action   string `json:"action" binding:"required,oneof=login create"`  // Must be either 'login' or 'create'
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
	var status int
	var err error

	if authRequest.Action == "login" {
		// Login with status code
		authResponse, status, err = authService.Login(authRequest.Email, authRequest.Password)
		if err != nil {
			ctx.JSON(status, gin.H{"error": err.Error()})
			return
		}
	} else if authRequest.Action == "create" {
		// Create user without status code handling (assuming CreateUser doesn't return a status)
		authResponse, status, err = authService.CreateUser(authRequest.Email, authRequest.Password)
		if err != nil {
			ctx.JSON(status, gin.H{"error": err.Error()})
			return
		}
		status = http.StatusCreated // Set status for successful user creation
	} else {
		// Invalid action
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid action"})
		return
	}

	// Return Auth Response
	ctx.JSON(status, authResponse)
}
