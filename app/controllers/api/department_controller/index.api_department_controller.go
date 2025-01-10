package api_department_controller

import (
	"net/http"
	"triesdi/app/configs/db_config"
	department_repository "triesdi/app/repository/department"
	department_request "triesdi/app/request/department"
	response "triesdi/app/responses/response"
	department_service "triesdi/app/service/department"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	response.BaseResponse(ctx, http.StatusOK, true, "OK", "Hello World")
}

func CreateDepartment(ctx *gin.Context) {

	var input department_request.CreateDepartmentRequest

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		response.BaseResponse(ctx, http.StatusBadRequest, false, "Gagal", "Gagal membuat")
	}

	if input.Name == "" {
		response.BaseResponse(ctx, http.StatusBadRequest, false, "Gagal", "Gagal membuat")
	}

	// Get DB Config
	db := db_config.GetDB()

	departmentRepository := department_repository.NewRepository(db)
	departmentService := department_service.NewService(departmentRepository)

	// Call the service to create a new department
	newDepartment, err := departmentService.CreateDepartment(input)
	if err != nil {
		response.BaseResponse(ctx, http.StatusInternalServerError, false, "Failed", "Error creating department: "+err.Error())
		return
	}

	// Return the successful response with the created department details
	response.BaseResponse(ctx, http.StatusOK, true, "Success", newDepartment)
}
