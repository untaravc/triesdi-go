package user_repository

type User struct {
	Id                int    `json:"id"`
	Email             string `json:"email"`
	Phone             string `json:"phone"`
	Password          string `json:"password"`
	Token             string `json:"token"`
	FileId            string `json:"file_id"`
	BankAccountName   string `json:"bank_account_name"`
	BankAccountHolder string `json:"bank_account_holder"`
	BankAccountNumber string `json:"bank_account_number"`
	FileUri           string `json:"file_uri"`
	FileThumbnailUri  string `json:"file_thumbnail_uri"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
}

type UserFilter struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}
