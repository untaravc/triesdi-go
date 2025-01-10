package department_service

import (
	"time"
	"triesdi/app/repository/department_repository"
	"triesdi/app/requests/department_request"
)

type Service interface {
	CreateDepartment(input department_request.DepartmentRequest) (department_repository.Department, error)
	UpdateDepartment(id int, input department_request.DepartmentRequest) (department_repository.Department, error)
	DeleteDepartment(id int) (department_repository.Department, error)
	GetDepartments(input department_request.DepartmentFilter) ([]department_repository.Department, error)
}

type service struct {
	repository department_repository.Repository
}

func NewService(repository department_repository.Repository) *service {
	return &service{repository}
}

// Get All Department base On Filter Request
func (s *service) GetDepartments(input department_request.DepartmentFilter) ([]department_repository.Department, error) {

	departmentFilter := department_request.DepartmentFilter{}
	departmentFilter.Name = input.Name
	// default limit=5&offset=0
	if input.Limit == 0 {
		departmentFilter.Limit = 5
	} else {
		departmentFilter.Limit = input.Limit
	}

	if input.Offset == 0 {
		departmentFilter.Offset = 0
	} else {
		departmentFilter.Offset = input.Offset
	}

	departments, err := s.repository.GetAll(departmentFilter)
	if err != nil {
		return departments, err
	}

	return departments, nil
}

func (s *service) CreateDepartment(input department_request.DepartmentRequest) (department_repository.Department, error) {
	department := department_repository.Department{}
	department.Name = input.Name
	department.ManagerId = 1

	newDepartment, err := s.repository.Save(department)
	if err != nil {
		return newDepartment, err
	}

	return newDepartment, nil
}

func (s *service) UpdateDepartment(id int, input department_request.DepartmentRequest) (department_repository.Department, error) {
	department := department_repository.Department{}
	department.ID = id
	department.Name = input.Name
	department.ManagerId = 1
	department.UpdatedAt = time.Now()

	newDepartment, err := s.repository.Update(id, department)
	if err != nil {
		return newDepartment, err
	}

	return newDepartment, nil
}

func (s *service) DeleteDepartment(id int) (error) {
	err := s.repository.SoftDelete(id)
	if err != nil {
		return err
	}

	return nil
}

