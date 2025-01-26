package user_response

import "triesdi/app/repository/user_repository"

type AuthResponse struct {
	Token string `json:"token"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type UserResponse struct {
	Email             string `json:"email"`
	Phone             string `json:"phone"`
	FileId            string `json:"fileId"`
	FileUri           string `json:"fileUri"`
	FileThumbnailUri  string `json:"fileThumbnailUri"`
	BankAccountName   string `json:"bankAccountName"`
	BankAccountHolder string `json:"bankAccountHolder"`
	BankAccountNumber string `json:"bankAccountNumber"`
}

func UserToUserResponse(user user_repository.User) UserResponse {
	return UserResponse{
		Email:             user.Email,
		Phone:             user.Phone,
		FileId:            user.FileId,
		FileUri:           user.FileUri,
		FileThumbnailUri:  user.FileThumbnailUri,
		BankAccountName:   user.BankAccountName,
		BankAccountHolder: user.BankAccountHolder,
		BankAccountNumber: user.BankAccountNumber,
	}
}
