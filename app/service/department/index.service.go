package department

import (
	repository "triesdi/app/repository/department"
	request "triesdi/app/request/department"
)

type Service interface {
	CreateDepartment(input request.CreateDepartmentRequest) (repository.Department, error)
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{repository}
}

func (s *service) CreateDepartment(input request.CreateDepartmentRequest) (repository.Department, error) {
	department := repository.Department{}
	department.Name = input.Name
	department.ManagerId = 1

	newDepartment, err := s.repository.Save(department)
	if err != nil {
		return newDepartment, err
	}

	return newDepartment, nil
}

