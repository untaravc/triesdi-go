package auth_service

import (
	"errors"
	"net/http"
	"triesdi/app/repository/auth_repository"
	"triesdi/app/responses/response"
	"triesdi/app/utils"
	"triesdi/app/utils/common"
)

type Service interface {
	CreateUser(email, password string) (response.AuthResponse, int, error)
	Login(email, password string) (response.AuthResponse, int, error)
}

type service struct {
	repository auth_repository.Repository
}

func NewService(repository auth_repository.Repository) *service {
	return &service{repository}
}

func (s *service) CreateUser(email, password string) (response.AuthResponse, int, error) {
	// Check Email Exist
	_, err := s.repository.GetUserByEmail(email)
	if err == nil {
		return response.AuthResponse{}, http.StatusConflict, errors.New("email already exist")
	}


	user, err := s.repository.CreateUser(email, password)
	if err != nil {
		return response.AuthResponse{}, http.StatusInternalServerError , err
	}

	// Generate Token
	token, err := utils.GenerateToken(user.Email)
	if err != nil {
		return response.AuthResponse{}, http.StatusInternalServerError, err
	}

	// Make Response Email & Token
	authResponse := response.AuthResponse{
		Email: user.Email,
		Token: token,
	}

	return authResponse, http.StatusCreated, nil
}

func (s *service) Login(email, password string) (response.AuthResponse, int, error) {
	user := auth_repository.Auth{Email: email, Password: password}

	// Get User By Email
	user, err := s.repository.GetUserByEmail(user.Email)
	if err != nil {
		return response.AuthResponse{}, http.StatusNotFound,err
	}

	// Check Password
	if !common.CheckPasswordHash(password, user.Password) {
		return response.AuthResponse{}, http.StatusUnauthorized, errors.New("incorrect password")
	}

	// Generate Token
	token, err := utils.GenerateToken(user.Email)
	if err != nil {
		return response.AuthResponse{}, http.StatusInternalServerError, err
	}

	// Make Response Email & Token
	authResponse := response.AuthResponse{
		Email: user.Email,
		Token: token,
	}

	return authResponse, http.StatusOK, nil
}


