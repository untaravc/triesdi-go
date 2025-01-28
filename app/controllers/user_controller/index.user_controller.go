package user_controller

import (
	"database/sql"
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

	Email := ""
	Phone := ""
	FileId := ""
	BankAccountName := ""
	BankAccountHolder := ""
	BankAccountNumber := ""
	FileUri := ""
	FileThumbnailUri := ""

	if users[0].Email.Valid {
		Email = users[0].Email.String
	}

	if users[0].Phone.Valid {
		Phone = users[0].Phone.String
	}

	if users[0].FileId.Valid {
		FileId = users[0].FileId.String
	}

	if users[0].BankAccountName.Valid {
		BankAccountName = users[0].BankAccountName.String
	}

	if users[0].BankAccountHolder.Valid {
		BankAccountHolder = users[0].BankAccountHolder.String
	}

	if users[0].BankAccountNumber.Valid {
		BankAccountNumber = users[0].BankAccountNumber.String
	}

	if users[0].FileUri.Valid {
		FileUri = users[0].FileUri.String
	}

	if users[0].FileThumbnailUri.Valid {
		FileThumbnailUri = users[0].FileThumbnailUri.String
	}

	user_response := user_response.UserResponse{
		Email:             Email,
		Phone:             Phone,
		FileId:            FileId,
		BankAccountName:   BankAccountName,
		BankAccountHolder: BankAccountHolder,
		BankAccountNumber: BankAccountNumber,
		FileUri:           FileUri,
		FileThumbnailUri:  FileThumbnailUri,
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
		FileId:            sql.NullString{String: user_request.FileId, Valid: true},
		BankAccountName:   sql.NullString{String: user_request.BankAccountName, Valid: true},
		BankAccountHolder: sql.NullString{String: user_request.BankAccountHolder, Valid: true},
		BankAccountNumber: sql.NullString{String: user_request.BankAccountNumber, Valid: true},
	}
	// get file
	if user_request.FileId != "" {
		file, err := file_repository.GetById(user_request.FileId)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		update.FileUri = sql.NullString{String: file.FileUri, Valid: true}
		update.FileThumbnailUri = sql.NullString{String: file.FileThumbnailUri, Valid: true}
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

func LinkPhone(c *gin.Context) {
	link_phone_request := user_request.LinkPhoneRequest{}

	if err := c.ShouldBindJSON(&link_phone_request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validator.ValidateStruct(link_phone_request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validator.FormatValidationError(err)})
		return
	}

	// get auth user
	auth_user := jwt.GetAuth(c)

	// validate unique
	user, _ := user_repository.UniqueUser(user_repository.User{
		Id:    auth_user.Id,
		Phone: sql.NullString{String: link_phone_request.Phone, Valid: true},
	})

	if user.Id != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "phone already exists"})
		return
	}

	// update user
	ruser, err_user := user_repository.UpdateUser(user_repository.User{Phone: sql.NullString{String: link_phone_request.Phone, Valid: true}}, auth_user.Id)

	if err_user != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err_user.Error()})
		return
	}

	c.JSON(http.StatusOK, user_response.UserToUserResponse(ruser))
}

func LinkEmail(c *gin.Context) {
	link_email_request := user_request.LinkEmailRequest{}

	if err := c.ShouldBindJSON(&link_email_request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validator.ValidateStruct(link_email_request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validator.FormatValidationError(err)})
		return
	}

	// get auth user
	auth_user := jwt.GetAuth(c)

	// validate unique
	user, _ := user_repository.UniqueUser(user_repository.User{
		Id:    auth_user.Id,
		Email: sql.NullString{String: link_email_request.Email, Valid: true},
	})

	if user.Id != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email already exists"})
		return
	}

	// update user
	ruser, err_user := user_repository.UpdateUser(user_repository.User{
		Email: sql.NullString{String: link_email_request.Email, Valid: true}},
		auth_user.Id)

	if err_user != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err_user.Error()})
		return
	}

	c.JSON(http.StatusOK, user_response.UserToUserResponse(ruser))
}
