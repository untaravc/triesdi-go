package auth_repository

type Auth struct {
	Email string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
	// Password string `json:"password" gorm:"not null"`
}