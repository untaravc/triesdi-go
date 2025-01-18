package user_response

import "triesdi/app/repository/user_repository"

type GetUserResponse struct {
	Preference string `json:"preference"`
	WeightUnit string `json:"weightUnit"`
	HeightUnit string `json:"heightUnit"`
	Weight     *int    `json:"weight"`
	Height     *int    `json:"height"`
	Email	   string `json:"email"`
	Name       string `json:"name"`
	ImageUri   string `json:"imageUri"`
}

type UpdateUserResponse struct {
	Preference string `json:"preference"`
	WeightUnit string `json:"weightUnit"`
	HeightUnit string `json:"heightUnit"`
	Weight     *int    `json:"weight"`
	Height     *int    `json:"height"`
	Name       string `json:"name"`
	ImageUri   string `json:"imageUri"`
}

func FormatGetUserResponse(user user_repository.User) GetUserResponse {

	// If user Weight null change to nil
	if user.Weight == nil {
		user.Weight = nil
	}

	// If user Height null change to nil
	if user.Height == nil {
		user.Height = nil
	}

	formatter := GetUserResponse{
		Preference: user.Preference,
		WeightUnit: user.WeightUnit,
		HeightUnit: user.HeightUnit,
		Weight:     user.Weight,
		Height:     user.Height,
		Email:      user.Email,
		Name:       user.Name,
		ImageUri:   user.ImageUri,
	}

	return formatter
}

func FormatUpdateUserResponse(user user_repository.User) UpdateUserResponse {
	formatter := UpdateUserResponse{
		Preference: user.Preference,
		WeightUnit: user.WeightUnit,
		HeightUnit: user.HeightUnit,
		Weight:     user.Weight,
		Height:     user.Height,
		Name:       user.Name,
		ImageUri:   user.ImageUri,
	}

	return formatter
}