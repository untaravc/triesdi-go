package user_service

import (
	"triesdi/app/repository/user_repository"
	"triesdi/app/requests/user_request"
)

type Service interface {
	 UpdateUser (id string, user user_request.UserRequest) error
     GetUser (id string) error
}

type service struct {
    repository user_repository.Repository
}

func NewService(repository user_repository.Repository) *service {
    return &service{repository}
}

func (s *service) UpdateUser(id string, user user_request.UserRequest) (user_repository.User, error) {
    updatedUser := user_repository.User{
        Preference: user.Preference,
        WeightUnit: user.WeightUnit,
        HeightUnit: user.HeightUnit,
        Weight: user.Weight,
        Height: user.Height,
        Name: user.Name,
        ImageUri: user.ImageUri,
    }
    updatedUser, err := s.repository.UpdateUser(id, updatedUser)
    if err != nil {
        return user_repository.User{}, err
    }
    return updatedUser, nil
}

func (s *service) GetUser(id string) (user_repository.User, error) {
    user, err := s.repository.GetUser(id)
    if err != nil {
        return user_repository.User{}, err
    }
    return user, nil
}
