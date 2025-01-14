package model

type Manager struct {
	// 	id
	// email
	// password
	// name
	// user_image_uri
	// company_name
	// company_image_uri
	// created_at
	// updated_at
	// deleted_at
	Id              int    `json:"id"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	Name            string `json:"name"`
	UserImageUri    string `json:"user_image_uri"`
	CompanyName     string `json:"company_name"`
	CompanyImageUri string `json:"company_image_uri"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	DeletedAt       string `json:"deleted_at"`
}
