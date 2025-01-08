package department_repository

import (
	"time"
	"triesdi/app/requests/department_request"

	"gorm.io/gorm"
)

type Repository interface {
	Save(department Department) (Department, error)
	Update(id int, department Department) (Department, error)
	SoftDelete(id int) error
	FindById(id int) (Department, error)
	GetAll(department_request.DepartmentFilter) ([]Department, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(department Department) (Department, error) {
	err := r.db.Create(&department).Error

	if err != nil {
		return department, err
	}

	return department, nil
}

func (r *repository) Update(id int, department Department) (Department, error) {
	// Update the department
	if err := r.db.Model(&department).Updates(department).Error; err != nil {
		return department, err
	}

	return department, nil
}

func (r *repository) FindById(id int) (Department, error) {
	var department Department

	// Find department where ID matches and deleted_at is NULL (soft delete check)
	if err := r.db.Where("id = ? AND deleted_at IS NULL", id).First(&department).Error; err != nil {
		return department, err
	}

	return department, nil
}


func (r *repository) SoftDelete(id int) error {
	var department Department

	if err := r.db.Model(&department).
	Where("id = ?", id).
	Update("deleted_at", time.Now()).Error; err != nil {
	return err // Error when soft-deleting the department
}

	return nil // Successfully soft-deleted the department
}

func (r *repository) GetAll(filter department_request.DepartmentFilter) ([]Department, error) {
	var departments []Department

	query := r.db.Where("deleted_at IS NULL")

	if filter.Name != "" {
		query = query.Where("name LIKE ?", "%"+filter.Name+"%")
	}

	if err := query.Limit(filter.Limit).Offset(filter.Offset).Find(&departments).Error; err != nil {
		return departments, err
	}

	return departments, nil
}

