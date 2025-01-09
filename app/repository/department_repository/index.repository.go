package department_repository

import "gorm.io/gorm"

type Repository interface {
	Save(department Department) (Department, error)
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