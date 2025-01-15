package auth_repository

import (
	"errors"
	"triesdi/app/utils/common"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

func (Auth) TableName() string {
	return "users"
}

type Repository interface {
	CreateUser(email, password string) (Auth, error)
	GetUserByEmail(email string) (Auth, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateUser(email, password string) (Auth, error) {
	// hash password
	hashedPassword, err := common.HashingPassword(password)
	if err != nil {
		return Auth{}, err
	}

	user := Auth{ID: uuid.New(), Email: email, Password: hashedPassword}
	if err := r.db.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) GetUserByEmail(email string) (Auth, error) {
	var user Auth
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, errors.New("user not found")
		}
		return user, err
	}
	return user, nil
}
