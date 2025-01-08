package department_request

type CreateDepartmentRequest struct {
	Name string `json:"name" form:"name" binding:"required,min=3,max=33"`
}