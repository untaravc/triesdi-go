package converter

import (
	"strconv"
	"triesdi/app/repository/department_repository"
	"triesdi/app/repository/employee_repository"
)

type DepartmentFormatter struct {
	DepartmentId int `json:"departmentId"`
	Name string `json:"name"`
}

type EmployeeFormatter struct {
	IdentityNumber string `json:"identityNumber"`
	Name string `json:"name"`
	EmployeeImageUri string `json:"employeeImageUri"`
	Gender string `json:"gender"`
	DepartmentId int `json:"departmentId"`
}

func FormatDepartment(department department_repository.Department) DepartmentFormatter {
	formatter := DepartmentFormatter{
		DepartmentId: department.ID,
		Name: department.Name,
	}

	return formatter
}

func FormatEmployee(employee employee_repository.Employee) EmployeeFormatter {
	formatter := EmployeeFormatter{
		IdentityNumber: employee.IdentityNumber,
		Name: employee.Name,
		EmployeeImageUri: employee.EmployeeImageUri,
		Gender: employee.Gender,
		DepartmentId: employee.DepartmentId,
	}

	return formatter
}

func StringToInt(s string) int {
	result, _ := strconv.Atoi(s)
	return result
}