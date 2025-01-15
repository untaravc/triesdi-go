package user_repository

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	UpdateUser(id string, user User) (User, error)
	GetUser(id string) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) UpdateUser(id string,user User) (User, error) {
	if err := r.db.Model(&User{}).Where("id = ?", id).Updates(user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) GetUser(id string) (User, error) {
	var user User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, errors.New("user not found")
		}
		return user, err
	}
	return user, nil
}


	