package auth_request

type AuthUser struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type LoginEmail struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,max=32,min=8"`
}

type LoginPhone struct {
	Phone    string `json:"phone" validate:"required,e164"`
	Password string `json:"password" validate:"required,max=32,min=8"`
}

type RegisterEmailRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,max=32,min=8"`
}

type RegisterPhoneRequest struct {
	Phone    string `json:"phone" validate:"required,e164"`
	Password string `json:"password" validate:"required,max=32,min=8"`
}
