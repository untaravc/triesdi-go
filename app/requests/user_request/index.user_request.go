package user_request

type UserRequest struct {
	FileId            string `json:"fileId" validate:"alphanumunicode"`
	BankAccountName   string `json:"bankAccountName" validate:"required,min=4,max=32"`
	BankAccountHolder string `json:"bankAccountHolder" validate:"required,min=4,max=32"`
	BankAccountNumber string `json:"bankAccountNumber" validate:"required,min=4,max=32"`
}

type LinkPhoneRequest struct {
	Phone string `json:"phone" validate:"required,e164"`
}

type LinkEmailRequest struct {
	Email string `json:"email" validate:"required,email"`
}
