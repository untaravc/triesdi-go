package v1_user_controller

import (
	"net/http"
	"triesdi/app/configs/db_config"
	"triesdi/app/repository/user_repository"
	"triesdi/app/requests/user_request"
	"triesdi/app/responses/response/user_response"
	"triesdi/app/service/user_service"
	"triesdi/app/validator"

	"github.com/gin-gonic/gin"
)

func UpdateUser(ctx *gin.Context) {

	id, exist := ctx.Get("id")
	if !exist {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	userRequest := new(user_request.UserRequest)
	if err := ctx.ShouldBindJSON(userRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := validator.ValidateStruct(userRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := db_config.GetDB()
	userRepository := user_repository.NewRepository(db)
	userService := user_service.NewService(userRepository)

	var user user_repository.User
	var userResponse user_response.UpdateUserResponse
	var err error

	user, err = userService.UpdateUser(id.(string), *userRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userResponse = user_response.FormatUpdateUserResponse(user)

	ctx.JSON(http.StatusOK,userResponse)
}

func GetUser(ctx *gin.Context) {

	id, exist := ctx.Get("id")
	if !exist {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	db := db_config.GetDB()
	userRepository := user_repository.NewRepository(db)
	userService := user_service.NewService(userRepository)

	var userResponse user_response.GetUserResponse
	var err error

	// Get user without status code handling (assuming GetUser doesn't return a status)
	user, err := userService.GetUser(id.(string))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	userResponse = user_response.FormatGetUserResponse(user)
		
	ctx.JSON(http.StatusOK, userResponse)
}
