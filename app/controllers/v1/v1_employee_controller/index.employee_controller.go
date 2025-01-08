package v1_employee_controller

import (
	"net/http"
	"triesdi/app/configs/db_config"
	"triesdi/app/repository/employee_repository"
	"triesdi/app/requests/employee_request"
	"triesdi/app/responses/response"
	"triesdi/app/service/employee_service"
	"triesdi/app/utils/converter"
	"triesdi/app/validator"

	"github.com/gin-gonic/gin"
)

// GetEmployees function
func GetEmployees(ctx *gin.Context) {
	response.BaseResponse(ctx, http.StatusOK, true, "OK", nil)
}

// CreateEmployee function
func CreateEmployee(ctx *gin.Context) {

	var input employee_request.EmployeeRequest

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		// Validation error
		errors := validator.FormatValidationError(err)
		response.BaseResponse(ctx, http.StatusBadRequest, false, "Bad request", errors)
		return
	}

	// Get DB Config
	db := db_config.GetDB()

	// Initialize the employee request
	employeeRepository := employee_repository.NewRepository(db)
	employeeService := employee_service.NewService(employeeRepository)

	// Call the service to create the employee
	employee, err := employeeService.CreateEmployee(input)
	if err != nil {
		// Return error response if something goes wrong
		response.BaseResponse(ctx, http.StatusInternalServerError, false, "Failed", "Error creating employee: "+err.Error())
		return
	}

	employeeFormatted := converter.FormatEmployee(employee)

	response.BaseResponse(ctx, http.StatusOK, true, "Success", employeeFormatted)
}

// UpdateEmployee function
func UpdateEmployee(ctx *gin.Context) {
	response.BaseResponse(ctx, http.StatusOK, true, "OK", nil)
}

// DeleteEmployee function
func DeleteEmployee(ctx *gin.Context) {
	response.BaseResponse(ctx, http.StatusOK, true, "OK", nil)
}

// Index function
func Index(ctx *gin.Context) {
	response.BaseResponse(ctx, http.StatusOK, true, "OK", "Hello World")
}
