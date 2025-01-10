package employee_repository

import (
	"time"
	"triesdi/app/requests/employee_request"

	"gorm.io/gorm"
)

type Repository interface {
	GetAll (employeeFilter employee_request.EmployeeFilter) ([]Employee, error)
	Save(employee Employee) (Employee, error)
	Update(identityNumber string, employee Employee) (Employee, error)
	FindById(id int) (Employee, error)
	SoftDelete(identityNumber string) error
	FindByIdentityNumber(identityNumber string) (Employee, error)
	FindByIdentityNumberTrashed(identityNumber string) (Employee, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll (filter employee_request.EmployeeFilter) ([]Employee, error) {
	var employees []Employee
	
	query := r.db.Where("deleted_at IS NULL")

	if filter.Name != "" {
		query = query.Where("name LIKE ?", "%"+filter.Name+"%")
	}

	if filter.Gender != "" {
		query = query.Where("gender = ?", filter.Gender)
	}

	if filter.DepartmentId != 0 {
		query = query.Where("department_id = ?", filter.DepartmentId)
	}

	if err := query.Limit(filter.Limit).Offset(filter.Offset).Find(&employees).Error; err != nil {
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

func (r *repository) Update(identityNumber string, employee Employee) (Employee, error) {
	// Find the employee by identityNumber and update the data
	if err := r.db.Model(&Employee{}).Where("identity_number = ? AND deleted_at IS NULL", identityNumber).Updates(employee).Error; err != nil {
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

func (r *repository) SoftDelete(identityNumber string) error {
	// Soft delete by setting deleted_at timestamp
	if err := r.db.Model(&Employee{}).Where("identity_number = ?", identityNumber).Update("deleted_at", time.Now()).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) FindByIdentityNumber(identityNumber string) (Employee, error) {
	var employee Employee

	if err := r.db.Where("identity_number = ? AND deleted_at IS NULL", identityNumber).First(&employee).Error; err != nil {
		return employee, err
	}

	return employee, nil
}

func (r *repository) FindByIdentityNumberTrashed(identityNumber string) (Employee, error) {
	var employee Employee

	if err := r.db.Where("identity_number = ? AND deleted_at IS NOT NULL", identityNumber).First(&employee).Error; err != nil {
		return employee, err
	}

	return employee, nil
}