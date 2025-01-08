package employee_repository

import (
	"time"
	"triesdi/app/requests/employee_request"

	"gorm.io/gorm"
)

type Repository interface {
	GetAll (employeeFilter employee_request.EmployeeFilter) ([]Employee, error)
	Save(employee Employee) (Employee, error)
	Update(id int, employee Employee) (Employee, error)
	FindById(id int) (Employee, error)
	SoftDelete(id int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll (filter employee_request.EmployeeFilter) ([]Employee, error) {
	var employees []Employee

	if err := r.db.Find(&employees).Error; err != nil {
		return employees, err
	}

	return employees, nil
}

func (r *repository) Save(employee Employee) (Employee, error) {
	err := r.db.Create(&employee).Error

	if err != nil {
		return employee, err
	}

	return employee, nil
}

func (r *repository) Update(id int, employee Employee) (Employee, error) {
	// Update the employee
	if err := r.db.Model(&employee).Updates(employee).Error; err != nil {
		return employee, err
	}

	return employee, nil
}

func (r *repository) FindById(id int) (Employee, error) {
	var employee Employee

	// Find employee where ID matches and deleted_at is NULL (soft delete check)
	if err := r.db.Where("id = ? AND deleted_at IS NULL", id).First(&employee).Error; err != nil {
		return employee, err
	}

	return employee, nil
}

func (r *repository) SoftDelete(id int) error {
	var employee Employee

	if err := r.db.Model(&employee).
		Where("id = ?", id).
		Update("deleted_at", time.Now()).Error; err != nil {
		return err // Error when soft-deleting the employee
	}

	return nil
}