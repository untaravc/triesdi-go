package employee_repository

import "time"

type Employee struct {
	ID        int
	IdentityNumber string
	DepartmentId int
	Name      string
	Gender   string
	EmployeeImageUri string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}