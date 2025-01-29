package user_repository

import (
	"database/sql"
	"time"
)

type User struct {
	Id                string         `json:"id"`
	Email             sql.NullString `json:"email"`
	Phone             sql.NullString `json:"phone"`
	Password          string         `json:"password"`
	Token             string         `json:"token"`
	FileId            sql.NullString `json:"file_id"`
	BankAccountName   sql.NullString `json:"bank_account_name"`
	BankAccountHolder sql.NullString `json:"bank_account_holder"`
	BankAccountNumber sql.NullString `json:"bank_account_number"`
	FileUri           sql.NullString `json:"file_uri"`
	FileThumbnailUri  sql.NullString `json:"file_thumbnail_uri"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
}

type UserFilter struct {
	Id    string   `json:"id"`
	Ids   []string `json:"ids"`
	Email string   `json:"email"`
	Phone string   `json:"phone"`
}
