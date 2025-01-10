package employee_service

import (
	"triesdi/app/configs/db_config"
	"triesdi/app/repository/department_repository"
	"triesdi/app/repository/employee_repository"
	"triesdi/app/requests/employee_request"
)

type Service interface {
	CreateEmployee(input employee_request.EmployeeRequest) (employee_repository.Employee, error)
	UpdateEmployee(id int, input employee_request.EmployeeRequest) (employee_repository.Employee, error)
	DeleteEmployee(id int) (employee_repository.Employee, error)
	GetEmployees(input employee_repository.Employee) ([]employee_repository.Employee, error)
	FindById(id int) (employee_repository.Employee, error)
}

type service struct {
	repository employee_repository.Repository
}

func NewService(repository employee_repository.Repository) *service {
	return &service{repository}
}

// Get All Employee base On Filter Request
func (s *service) GetEmployees(input employee_request.EmployeeFilter) ([]employee_repository.Employee, error) {
	
	employeeFilter := employee_request.EmployeeFilter{}
	employeeFilter.Name = input.Name
	employeeFilter.DepartmentId = input.DepartmentId
	employeeFilter.Gender = input.Gender

	// default limit=5&offset=0
	if input.Limit == 0 {
		employeeFilter.Limit = 5
	} else {
		employeeFilter.Limit = input.Limit
	}

	if input.Offset == 0 {
		employeeFilter.Offset = 0
	} else {
		employeeFilter.Offset = input.Offset
	}

	employees, err := s.repository.GetAll(employeeFilter)
	if err != nil {
		return employees, err
	}

	return employees, nil
}

func (s *service) CreateEmployee(input employee_request.EmployeeRequest) (employee_repository.Employee, error) {

	// Check if the department exists
	_, err := department_repository.NewRepository(db_config.GetDB()).FindById(input.DepartmentId)
	if err != nil {
		// Return error response if department does not exist
		return employee_repository.Employee{}, err
	}

	employee := employee_repository.Employee{}
	employee.Name = input.Name
	employee.DepartmentId = input.DepartmentId
	employee.IdentityNumber = input.IdentityNumber
	employee.EmployeeImageUri = input.EmployeeImageUri
	employee.Gender = input.Gender

	newEmployee, err := s.repository.Save(employee)
	if err != nil {
		return newEmployee, err
	}

	return newEmployee, nil
}

func (s *service) UpdateEmployee(identityNumber string, input employee_request.EmployeeRequest) (employee_repository.Employee, error) {
	
	// Check if the department exists
	_, err := department_repository.NewRepository(db_config.GetDB()).FindById(input.DepartmentId)
	if err != nil {
		// Return error response if department does not exist
		return employee_repository.Employee{}, err
	}

	// Check if employee exists
	_, err = s.repository.FindByIdentityNumber(identityNumber)
	if err != nil {
		// Return error response if employee does not exist
		return employee_repository.Employee{}, err
	}
	
	employee := employee_repository.Employee{}
	employee.Name = input.Name
	employee.DepartmentId = input.DepartmentId
	employee.IdentityNumber = input.IdentityNumber
	employee.EmployeeImageUri = input.EmployeeImageUri
	employee.Gender = input.Gender

	newEmployee, err := s.repository.Update(identityNumber, employee)
	if err != nil {
		return newEmployee, err
	}

	return newEmployee, nil
}

func (s *service) DeleteEmployee(identityNumber string) (employee_repository.Employee, error) {
	err := s.repository.SoftDelete(identityNumber)
	if err != nil {
		return employee_repository.Employee{}, err
	}

	employee, err := s.repository.FindByIdentityNumberTrashed(identityNumber)
	if err != nil {
		return employee, err
	}

	return employee, nil
}

func (s *service) FindById(id int) (employee_repository.Employee, error) {
	employee, err := s.repository.FindById(id)
	if err != nil {
		return employee, err
	}

	return employee, nil	
}