package v1_management_controller

import (
	"net/http"
	"triesdi/app/configs/db_config"
	model "triesdi/app/models"
	"triesdi/app/requests/manager_request"
	"triesdi/app/responses/manager_response"
	"triesdi/app/responses/response"
	"triesdi/app/validator"

	"github.com/gin-gonic/gin"
)

func GetAuth(ctx *gin.Context) {

	email, exists := ctx.Get("email")

	if !exists {
		ctx.JSON(401, gin.H{"message": "Unauthorized"})
	}

	management := new(model.Manager)

	err := db_config.DB.Table("managers").Where("email = ?", email).Find(management).Error

	if err != nil {
		ctx.JSON(401, gin.H{"message": err.Error()})
		return
	}

	manager_res := manager_response.NewManagerResponse(*management)
	ctx.JSON(200, gin.H{
		"email":           manager_res.Email,
		"name":            manager_res.Name,
		"userImageUri":    manager_res.UserImageUri,
		"companyName":     manager_res.CompanyName,
		"companyImageUri": manager_res.CompanyImageUri,
	})
}

func Update(ctx *gin.Context) {
	email, exists := ctx.Get("email")

	if !exists {
		ctx.JSON(401, gin.H{"message": "Unauthorized"})
	}

	manager_request := new(manager_request.ManagerRequest)

	err := ctx.ShouldBindJSON(&manager_request)
	if err != nil {
		// Validator
		errors := validator.FormatValidationError(err)
		errorsMessage := gin.H{"errors": errors}
		response.BaseResponse(ctx, http.StatusBadRequest, false, "Failed", errorsMessage)
		return
	}

	if err := db_config.DB.Table("managers").Where("email = ?", email).Updates(manager_request).Error; err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "success"})
}
