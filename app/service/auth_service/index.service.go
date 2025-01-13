package auth_service

import (
	"errors"
	"triesdi/app/repository/auth_repository"
	"triesdi/app/responses/response"
	"triesdi/app/utils"
	"triesdi/app/utils/common"
)

type Service interface {
	CreateUser(email, password string) (response.AuthResponse, error)
	Login(email, password string) (response.AuthResponse, error)
}

type service struct {
	repository auth_repository.Repository
}

func NewService(repository auth_repository.Repository) *service {
	return &service{repository}
}

func (s *service) CreateUser(email, password string) (response.AuthResponse, error) {
	user, err := s.repository.CreateUser(email, password)
	if err != nil {
		return response.AuthResponse{}, err
	}

	// Generate Token
	token, err := utils.GenerateToken(user.Email)
	if err != nil {
		return response.AuthResponse{}, err
	}

	// Make Response Email & Token
	authResponse := response.AuthResponse{
		Email: user.Email,
		Token: token,
	}

	return authResponse, nil
}

func (s *service) Login(email, password string) (response.AuthResponse, error) {
	user := auth_repository.Auth{Email: email, Password: password}

	// Get User By Email
	user, err := s.repository.GetUserByEmail(user.Email)
	if err != nil {
		return response.AuthResponse{}, err
	}

	// Check Password
	if !common.CheckPasswordHash(password, user.Password) {
		return response.AuthResponse{}, errors.New("incorrect password")
	}

	// Generate Token
	token, err := utils.GenerateToken(user.Email)
	if err != nil {
		return response.AuthResponse{}, err
	}

	// Make Response Email & Token
	authResponse := response.AuthResponse{
		Email: user.Email,
		Token: token,
	}

	return authResponse, nil
}


