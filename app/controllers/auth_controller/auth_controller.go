package auth_controller

import (
	"database/sql"
	"net/http"
	"strconv"
	"triesdi/app/repository/user_repository"
	"triesdi/app/requests/auth_request"
	"triesdi/app/responses/user_response"
	"triesdi/app/utils/common"
	"triesdi/app/utils/jwt"
	"triesdi/app/utils/validator"

	"github.com/gin-gonic/gin"
)

func LoginEmail(c *gin.Context) {
	login_email_request := auth_request.LoginEmail{}

	if err := c.ShouldBindJSON(&login_email_request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validator.ValidateStruct(login_email_request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validator.FormatValidationError(err)})
		return
	}

	users, _ := user_repository.GetUsers(user_repository.UserFilter{Email: login_email_request.Email})

	if len(users) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// check password bcript
	if !common.CheckPasswordHash(login_email_request.Password, users[0].Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	token, err := jwt.GenerateToken(users[0].Id, users[0].Email.String, users[0].Phone.String)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	auth_response := user_response.AuthResponse{
		Email: users[0].Email.String,
		Phone: users[0].Phone.String,
		Token: token,
	}

	c.JSON(200, auth_response)
}

func LoginPhone(c *gin.Context) {
	login_phone_request := auth_request.LoginPhone{}

	if err := c.ShouldBindJSON(&login_phone_request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validator.ValidateStruct(login_phone_request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validator.FormatValidationError(err)})
		return
	}

	users, _ := user_repository.GetUsers(user_repository.UserFilter{Phone: login_phone_request.Phone})

	if len(users) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// check password bcript
	if !common.CheckPasswordHash(login_phone_request.Password, users[0].Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	token, err := jwt.GenerateToken(users[0].Id, users[0].Email.String, users[0].Phone.String)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	auth_response := user_response.AuthResponse{
		Email: users[0].Email.String,
		Phone: users[0].Phone.String,
		Token: token,
	}

	c.JSON(200, auth_response)
}

func RegisterEmail(c *gin.Context) {
	register_email_request := auth_request.RegisterEmailRequest{}

	if err := c.ShouldBindJSON(&register_email_request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validator.ValidateStruct(register_email_request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validator.FormatValidationError(err)})
		return
	}

	users, _ := user_repository.GetUsers(user_repository.UserFilter{Email: register_email_request.Email})

	if len(users) > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		return
	}

	user := user_repository.User{
		Email:    sql.NullString{String: register_email_request.Email},
		Password: register_email_request.Password,
	}

	id, err := user_repository.CreateUser(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := jwt.GenerateToken(strconv.Itoa(id), user.Email.String, user.Phone.String)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	auth_response := user_response.AuthResponse{Email: user.Email.String, Phone: "", Token: token}
	c.JSON(200, auth_response)
}

func RegisterPhone(c *gin.Context) {
	register_phone_request := auth_request.RegisterPhoneRequest{}

	if err := c.ShouldBindJSON(&register_phone_request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validator.ValidateStruct(register_phone_request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validator.FormatValidationError(err)})
		return
	}

	users, _ := user_repository.GetUsers(user_repository.UserFilter{Phone: register_phone_request.Phone})

	if len(users) > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Phone already exists"})
		return
	}

	user := user_repository.User{
		Phone:    sql.NullString{String: register_phone_request.Phone},
		Password: register_phone_request.Password,
	}

	id, err := user_repository.CreateUser(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := jwt.GenerateToken(strconv.Itoa(id), user.Email.String, user.Phone.String)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	auth_response := user_response.AuthResponse{Email: "", Phone: user.Phone.String, Token: token}
	c.JSON(200, auth_response)
}
