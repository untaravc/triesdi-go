package department_request

type DepartmentRequest struct {
	Name string `json:"name" form:"name" binding:"required,min=3,max=33"`
}

// Filter Offset Limit & Name
type DepartmentFilter struct {
	Offset int    `json:"offset" form:"offset"`
	Limit  int    `json:"limit" form:"limit"`
	Name   string `json:"name" form:"name"`
}

