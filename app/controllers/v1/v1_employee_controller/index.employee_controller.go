package v1_employee_controller

import (
	"net/http"
	"strconv"
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
	// Get DB Config
	db := db_config.GetDB()

	// Initialize the employee filter
	var filter employee_request.EmployeeFilter

	// Parse limit, offset, and name query parameters
	limitStr := ctx.DefaultQuery("limit", "5")  // Default value is 5
	offsetStr := ctx.DefaultQuery("offset", "0") // Default value is 0
	name := ctx.DefaultQuery("name", "")         // Default value is empty string
	gender := ctx.DefaultQuery("gender","")
	departmentId := ctx.DefaultQuery("departmentId","")

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
	filter.Gender = gender
	filter.DepartmentId = converter.StringToInt(departmentId)

	// Initialize the repository and service
	employeeRepository := employee_repository.NewRepository(db)
	employeeService := employee_service.NewService(employeeRepository)

	// Call the service to get the list of employees
	employees, err := employeeService.GetEmployees(filter)
	if err != nil {
		// Return error response if something goes wrong
		response.BaseResponse(ctx, http.StatusInternalServerError, false, "Failed", "Error getting employees: "+err.Error())
		return
	}

	formattedEmployees := []converter.EmployeeFormatter{}

	for _, employee := range employees {
		formattedEmployees = append(formattedEmployees, converter.FormatEmployee(employee))
	}
	
	response.BaseResponse(ctx, http.StatusOK, true, "Success", formattedEmployees)
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
	
	identityNumber := ctx.Param("identityNumber")

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

	employeeRepository := employee_repository.NewRepository(db)
	employeeService := employee_service.NewService(employeeRepository)

	// Call the service to update the employee
	employee, err := employeeService.UpdateEmployee(identityNumber, input)
	if err != nil {
		// Return error response if something goes wrong
		response.BaseResponse(ctx, http.StatusInternalServerError, false, "Failed", "Error updating employee: "+err.Error())
		return
	}

	employeeFormatted := converter.FormatEmployee(employee)

	response.BaseResponse(ctx, http.StatusOK, true, "Success", employeeFormatted)
}

// DeleteEmployee function
func DeleteEmployee(ctx *gin.Context) {
	identityNumber := ctx.Param("identityNumber")

	// Get DB Config
	db := db_config.GetDB()

	employeeRepository := employee_repository.NewRepository(db)
	employeeService := employee_service.NewService(employeeRepository)

	// Call the service to delete the employee
	_, err := employeeService.DeleteEmployee(identityNumber)
	if err != nil {
		// Return error response if something goes wrong
		response.BaseResponse(ctx, http.StatusInternalServerError, false, "Failed", "Error deleting employee: "+err.Error())
		return
	}

	// employeeFormatted := converter.FormatEmployee(employee)

	response.BaseResponse(ctx, http.StatusOK, true, "Success", nil)
}

// Index function
func Index(ctx *gin.Context) {
	response.BaseResponse(ctx, http.StatusOK, true, "OK", nil)
}
