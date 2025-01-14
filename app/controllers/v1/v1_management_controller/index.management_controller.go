package v1_management_controller

import (
	"triesdi/app/configs/db_config"
	model "triesdi/app/models"
	"triesdi/app/responses/manager_response"

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

}
