package employee_request

type EmployeeRequest struct {
	IdentityNumber   string `json:"identityNumber" form:"identityNumber" binding:"required,min=5,max=33"`
	Name             string `json:"name" form:"name" binding:"required,min=4,max=33"`
	EmployeeImageUri string `json:"employeeImageUri" form:"employeeImageUri" binding:"required,uri"`
	Gender           string `json:"gender" form:"gender" binding:"required,oneof=male female"`
	DepartmentId     int `json:"departmentId" form:"departmentId" binding:"required"`
}

type EmployeeFilter struct {
	Offset int    `json:"offset" form:"offset"`
	Limit  int    `json:"limit" form:"limit"`
	// IdentityNumber  string `json:"identity_number" form:"identity_number"`
	Name   string `json:"name" form:"name"`
	Gender string `json:"gender" form:"gender" binding:"oneof=male female"`
	DepartmentId int `json:"department_id" form:"department_id"`
}
