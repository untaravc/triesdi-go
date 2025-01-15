package converter

import (
	"strconv"
)

type DepartmentFormatter struct {
	DepartmentId int `json:"departmentId"`
	Name string `json:"name"`
}

// type EmployeeFormatter struct {
// 	IdentityNumber string `json:"identityNumber"`
// 	Name string `json:"name"`
// 	EmployeeImageUri string `json:"employeeImageUri"`
// 	Gender string `json:"gender"`
// 	DepartmentId int `json:"departmentId"`
// }

// func FormatEmployee(employee employee_repository.Employee) EmployeeFormatter {
// 	formatter := EmployeeFormatter{
// 		IdentityNumber: employee.IdentityNumber,
// 		Name: employee.Name,
// 		EmployeeImageUri: employee.EmployeeImageUri,
// 		Gender: employee.Gender,
// 		DepartmentId: employee.DepartmentId,
// 	}

// 	return formatter
// }

func StringToInt(s string) int {
	result, _ := strconv.Atoi(s)
	return result
}