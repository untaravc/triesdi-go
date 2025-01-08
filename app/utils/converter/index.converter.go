package converter

import (
	"strconv"
	"triesdi/app/repository/department_repository"
)

type DepartmentFormatter struct {
	DepartmentId int `json:"departmentId"`
	Name string `json:"name"`
}

func FormatDepartment(department department_repository.Department) DepartmentFormatter {
	formatter := DepartmentFormatter{
		DepartmentId: department.ID,
		Name: department.Name,
	}

	return formatter
}

func StringToInt(s string) int {
	result, _ := strconv.Atoi(s)
	return result
}