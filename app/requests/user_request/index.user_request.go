package user_request

type UserRequest struct {
	// "fileId": "", // string | not required | should be a valid fileId
	// "bankAccountName": "", // string | required | minLength: 4 | maxLength | 32
	// "bankAccountHolder": "", // string | required | minLength: 4 | maxLength | 32
	// "bankAccountNumber": "", // string | required | minLength: 4 | maxLength | 32
	FileId            string `json:"fileId" validate:"alphanumunicode"`
	BankAccountName   string `json:"bankAccountName" validate:"required,min=4,max=32"`
	BankAccountHolder string `json:"bankAccountHolder" validate:"required,min=4,max=32"`
	BankAccountNumber string `json:"bankAccountNumber" validate:"required,min=4,max=32"`
}
