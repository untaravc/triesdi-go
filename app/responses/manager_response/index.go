package manager_response

import model "triesdi/app/models"

type ManagerResponse struct {
	// email
	// name
	// userImageUri
	// companyName
	// companyImageUri

	Email           string `json:"email"`
	Name            string `json:"name"`
	UserImageUri    string `json:"userImageUri"`
	CompanyName     string `json:"companyName"`
	CompanyImageUri string `json:"companyImageUri"`
}

func NewManagerResponse(manager model.Manager) ManagerResponse {
	return ManagerResponse{
		Email:           manager.Email,
		Name:            manager.Name,
		UserImageUri:    manager.UserImageUri,
		CompanyName:     manager.CompanyName,
		CompanyImageUri: manager.CompanyImageUri,
	}
}
