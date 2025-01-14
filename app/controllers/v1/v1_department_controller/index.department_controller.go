package v1_department_controller

import (
	"net/http"
	"strconv"
	"triesdi/app/configs/db_config"
	"triesdi/app/repository/department_repository"
	"triesdi/app/requests/department_request"
	"triesdi/app/responses/response"
	"triesdi/app/service/department_service"
	"triesdi/app/utils/converter"
	"triesdi/app/validator"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	response.BaseResponse(ctx, http.StatusOK, true, "OK", "Hello World")
}

func GetDepartments(ctx *gin.Context) {
	// Get DB Config
	db := db_config.GetDB()

	// Initialize the department filter
	var filter department_request.DepartmentFilter

	// Parse limit, offset, and name query parameters
	limitStr := ctx.DefaultQuery("limit", "5")  // Default value is 5
	offsetStr := ctx.DefaultQuery("offset", "0") // Default value is 0
	name := ctx.DefaultQuery("name", "")         // Default value is empty string

	// Convert limit and offset to integers
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 5 // Default to 5 if invalid limit
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = 0 // Default to 0 if invalid offset
	}

	// Populate the filter struct with query parameters
	filter.Limit = limit
	filter.Offset = offset
	filter.Name = name

	// Initialize the repository and service
	departmentRepository := department_repository.NewRepository(db)
	departmentService := department_service.NewService(departmentRepository)

	// Call the service to get the list of departments
	departments, err := departmentService.GetDepartments(filter)
	if err != nil {
		// Return error response if something goes wrong
		response.BaseResponse(ctx, http.StatusInternalServerError, false, "Failed", "Error getting departments: "+err.Error())
		return
	}

	formattedDepartments := []converter.DepartmentFormatter{}
	for _, department := range departments {
		formattedDepartments = append(formattedDepartments, converter.FormatDepartment(department))
	}

	// Return the success response with the department list
	response.BaseResponse(ctx, http.StatusOK, true, "Success", formattedDepartments)
}

func CreateDepartment (ctx *gin.Context) {
	
	var input department_request.DepartmentRequest

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		// Validator
		errors := validator.FormatValidationError(err)
		errorsMessage := gin.H{"errors": errors}
		response.BaseResponse(ctx, http.StatusBadRequest, false, "Failed",errorsMessage)
		return
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

	convertedResponse := converter.FormatDepartment(newDepartment)

	// Return the successful response with the created department details
	response.BaseResponse(ctx, http.StatusCreated, true, "Success", convertedResponse)
}

func UpdateDepartment (ctx *gin.Context) {
	// Get the department ID from the URL
	id := ctx.Param("id")

	// Convert the ID to an integer
	idConverted := converter.StringToInt(id)

	// Get the department details from the request body
	var input department_request.DepartmentRequest

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		// Validator
		errors := validator.FormatValidationError(err)
		errorsMessage := gin.H{"errors": errors}
		response.BaseResponse(ctx, http.StatusBadRequest, false, "Failed",errorsMessage)
		return
	}

	// Get DB Config
	db := db_config.GetDB()

	departmentRepository := department_repository.NewRepository(db)
	departmentService := department_service.NewService(departmentRepository)

	// Check If Exist
	_, err = departmentRepository.FindById(idConverted)
	if err != nil {
		response.BaseResponse(ctx, http.StatusNotFound, false, "Failed", "Department not found")
		return
	}

	// Call the service to update the department
	newDepartment, err := departmentService.UpdateDepartment(idConverted, input)
	if err != nil {
		// Return the error response if the service returns an error
		response.BaseResponse(ctx, http.StatusInternalServerError, false, "Failed", "Error updating department: "+err.Error())
		return
	}

	convertedResponse := converter.FormatDepartment(newDepartment)

	response.BaseResponse(ctx, http.StatusOK, true, "Success", convertedResponse)
}

func GetDepartment (ctx *gin.Context) {
	// Get the department ID from the URL
	id := ctx.Param("id")

	// Convert the ID to an integer
	idConverted := converter.StringToInt(id)

	// Get DB Config
	db := db_config.GetDB()

	departmentRepository := department_repository.NewRepository(db)

	// Call the repository to get the department details
	department, err := departmentRepository.FindById(idConverted)
	if err != nil {
		response.BaseResponse(ctx, http.StatusNotFound, false, "Failed", "Department not found")
		return
	}

	convertedResponse := converter.FormatDepartment(department)

	response.BaseResponse(ctx, http.StatusOK, true, "Success", convertedResponse)
}

func DeleteDepartment (ctx *gin.Context) {
	// Get the department ID from the URL
	id := ctx.Param("id")

	// Convert the ID to an integer
	idConverted := converter.StringToInt(id)

	// Get DB Config
	db := db_config.GetDB()

	departmentRepository := department_repository.NewRepository(db)

	// Check If Exist
	_, err := departmentRepository.FindById(idConverted)
	if err != nil {
		response.BaseResponse(ctx, http.StatusNotFound, false, "Failed", "Department not found")
		return
	}

	// Call the repository to delete the department
	err = departmentRepository.SoftDelete(idConverted)
	if err != nil {
		response.BaseResponse(ctx, http.StatusInternalServerError, false, "Failed", "Error deleting department: "+err.Error())
		return
	}

	response.BaseResponse(ctx, http.StatusOK, true, "Success", "Department deleted successfully")
}