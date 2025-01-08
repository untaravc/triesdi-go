package routes

import (
	"triesdi/app/controllers/v1/v1_auth_controller"
	"triesdi/app/controllers/v1/v1_department_controller"
	"triesdi/app/controllers/v1/v1_employee_controller"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {
	route := app

	route.POST("/v1/auth", v1_auth_controller.Auth)

	// Department
	route.GET("/v1/department", v1_department_controller.GetDepartments)
	route.POST("/v1/department", v1_department_controller.CreateDepartment)
	route.PATCH("/v1/department/:id", v1_department_controller.UpdateDepartment)
	route.DELETE("/v1/department/:id", v1_department_controller.DeleteDepartment)

	// Employee
	route.GET("/v1/employee", v1_employee_controller.GetEmployees)
	route.POST("/v1/employee", v1_employee_controller.CreateEmployee)
	route.PATCH("/v1/employee/:id", v1_employee_controller.UpdateEmployee)
	route.DELETE("/v1/employee/:id", v1_employee_controller.DeleteEmployee)
}
