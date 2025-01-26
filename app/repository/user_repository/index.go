package user_repository

import (
	"fmt"
	"triesdi/app/utils/common"
	"triesdi/app/utils/database"
)

const DB_NAME = "users"

func GetUsers(filter UserFilter) ([]User, error) {
	selector := "id, phone, email, password"
	query := fmt.Sprintf("SELECT %s FROM %s", selector, DB_NAME)

	if filter.Email != "" {
		query += fmt.Sprintf(" WHERE email = '%s'", filter.Email)
	}

	if filter.Phone != "" {
		query += fmt.Sprintf(" WHERE phone = '%s'", filter.Phone)
	}

	if filter.Id != "" {
		query += fmt.Sprintf(" WHERE id = '%s'", filter.Id)
	}

	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		var id int
		var phone, email, password string
		if err := rows.Scan(&id, &phone, &email, &password); err != nil {
			return nil, err
		}

		users = append(users, User{Id: id, Phone: phone, Email: email, Password: password})
	}

	return users, nil
}

func CreateUser(user User) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (phone, email, password) VALUES ($1, $2, $3) RETURNING id", DB_NAME)

	// hash password
	hashedPassword, err_hash := common.HashingPassword(user.Password)
	if err_hash != nil {
		return 0, err_hash
	}

	var insertedID int
	err := database.DB.QueryRow(query, user.Phone, user.Email, hashedPassword).Scan(&insertedID)
	if err != nil {
		return 0, err
	}

	return insertedID, nil
}

func UpdateUser(user User, id string) (ruser User, err error) {
	query := fmt.Sprintf("UPDATE %s SET ", DB_NAME)

	if user.FileId != "" {
		query += fmt.Sprintf("file_id = %s, ", user.FileId)
	}

	if user.BankAccountName != "" {
		query += fmt.Sprintf("bank_account_name = '%s', ", user.BankAccountName)
	}

	if user.BankAccountHolder != "" {
		query += fmt.Sprintf("bank_account_holder = '%s', ", user.BankAccountHolder)
	}

	if user.BankAccountNumber != "" {
		query += fmt.Sprintf("bank_account_number = '%s', ", user.BankAccountNumber)
	}

	if user.FileUri != "" {
		query += fmt.Sprintf("file_uri = '%s', ", user.FileUri)
	}

	if user.FileThumbnailUri != "" {
		query += fmt.Sprintf("file_thumbnail_uri = '%s', ", user.FileThumbnailUri)
	}

	query += "updated_at = now() WHERE id = $1"
	query += " RETURNING id, phone, email, file_id, file_uri, file_thumbnail_uri, bank_account_name, bank_account_holder, bank_account_number"

	err_query := database.DB.QueryRow(query, id).
		Scan(&user.Id, &user.Phone, &user.Email, &user.FileId, &user.FileUri, &user.FileThumbnailUri, &user.BankAccountName, &user.BankAccountHolder, &user.BankAccountNumber)
	if err_query != nil {
		return user, err_query
	}

	return user, nil
}
