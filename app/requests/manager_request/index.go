package manager_request

type ManagerRequest struct {
	// "email": "name@name.com", // should in email format
	// "name": "", // string | minLength 4 | maxLength 52
	// "userImageUri": "", // string | should be an URI
	// "companyName": "", // string | minLength 4 | maxLength 52
	// "companyImageUri": "", // string | should be an URI
	Email           string `json:"email" form:"email" binding:"required,email"`                    // should in email format
	Name            string `json:"name" form:"name" binding:"required,min=4,max=52"`               // string | minLength 4 | maxLength 52
	UserImageUri    string `json:"userImageUri" form:"userImageUri" binding:"required,uri"`        // string | should be an URI
	CompanyName     string `json:"companyName" form:"companyName" binding:"required,min=4,max=52"` // string | minLength 4 | maxLength 52
	CompanyImageUri string `json:"companyImageUri" form:"companyImageUri" binding:"required,uri"`  // string | should be an URI
}
