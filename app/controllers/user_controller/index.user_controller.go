package user_controller

import (
	"net/http"
	"triesdi/app/repository/file_repository"
	"triesdi/app/repository/user_repository"
	"triesdi/app/requests/user_request"
	"triesdi/app/responses/user_response"
	"triesdi/app/utils/jwt"
	"triesdi/app/utils/validator"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	auth_user := jwt.GetAuth(c)

	users, err := user_repository.GetUsers(user_repository.UserFilter{Id: auth_user.Id})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(users) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	user_response := user_response.UserResponse{
		Email:             users[0].Email,
		Phone:             users[0].Phone,
		FileId:            users[0].FileId,
		BankAccountName:   users[0].BankAccountName,
		BankAccountHolder: users[0].BankAccountHolder,
		BankAccountNumber: users[0].BankAccountNumber,
		FileUri:           users[0].FileUri,
		FileThumbnailUri:  users[0].FileThumbnailUri,
	}
	c.JSON(http.StatusOK, user_response)
}

func Update(c *gin.Context) {
	user_request := user_request.UserRequest{}

	if err := c.ShouldBindJSON(&user_request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validator.ValidateStruct(user_request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validator.FormatValidationError(err)})
		return
	}

	// file := file_repository.File{}

	update := user_repository.User{
		FileId:            user_request.FileId,
		BankAccountName:   user_request.BankAccountName,
		BankAccountHolder: user_request.BankAccountHolder,
		BankAccountNumber: user_request.BankAccountNumber,
	}
	// get file
	if user_request.FileId != "" {
		file, err := file_repository.GetById(user_request.FileId)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		update.FileUri = file.FileUri
		update.FileThumbnailUri = file.FileThumbnailUri
	}

	// get auth user
	auth_user := jwt.GetAuth(c)

	// update user
	ruser, err_user := user_repository.UpdateUser(update, auth_user.Id)

	if err_user != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err_user.Error()})
		return
	}

	c.JSON(http.StatusOK, user_response.UserToUserResponse(ruser))
}

func LinkPhone(c *gin.Context) {}

func LinkEmail(c *gin.Context) {}
