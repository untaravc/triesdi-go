package user_repository

import (
	"database/sql"
	"fmt"
	"strings"
	"triesdi/app/utils/common"
	"triesdi/app/utils/database"
)

const DB_NAME = "users"

func GetUsers(filter UserFilter) ([]User, error) {
	selector := "id, phone, email, password, file_id, file_uri, file_thumbnail_uri, bank_account_name, bank_account_holder, bank_account_number"
	query := fmt.Sprintf("SELECT %s FROM %s", selector, DB_NAME)

	conditions := make([]string, 0)
	if filter.Email != "" {
		conditions = append(conditions, fmt.Sprintf("email = '%s'", filter.Email))
	}

	if filter.Phone != "" {
		conditions = append(conditions, fmt.Sprintf("phone = '%s'", filter.Phone))
	}

	if filter.Id != "" {
		conditions = append(conditions, fmt.Sprintf("id = '%s'", filter.Id))
	}
	if len(filter.Ids) > 0 {
		conditions = append(conditions, fmt.Sprintf("id IN (%s)", strings.Join(filter.Ids, ",")))
	}

	// Add conditions to the query
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		var id, password string
		var phone, email, file_id, file_uri, file_thumbnail_uri, bank_account_name, bank_account_holder, bank_account_number sql.NullString

		if err := rows.Scan(&id, &phone, &email, &password, &file_id, &file_uri, &file_thumbnail_uri, &bank_account_name, &bank_account_holder, &bank_account_number); err != nil {
			return nil, err
		}

		users = append(users, User{
			Id:                id,
			Phone:             sql.NullString{String: phone.String, Valid: true},
			Email:             sql.NullString{String: email.String, Valid: true},
			Password:          password,
			FileId:            sql.NullString{String: file_id.String, Valid: true},
			FileUri:           sql.NullString{String: file_uri.String, Valid: true},
			FileThumbnailUri:  sql.NullString{String: file_thumbnail_uri.String, Valid: true},
			BankAccountName:   sql.NullString{String: bank_account_name.String, Valid: true},
			BankAccountHolder: sql.NullString{String: bank_account_holder.String, Valid: true},
			BankAccountNumber: sql.NullString{String: bank_account_number.String, Valid: true},
		})
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

func UpdateUser(user User, user_id string) (ruser User, err error) {
	query := fmt.Sprintf("UPDATE %s SET ", DB_NAME)

	if user.Email.Valid {
		query += fmt.Sprintf("email = '%s', ", user.Email.String)
	}

	if user.Phone.Valid {
		query += fmt.Sprintf("phone = '%s', ", user.Phone.String)
	}

	if user.FileId.Valid {
		query += fmt.Sprintf("file_id = '%s', ", user.FileId.String)
	}

	if user.BankAccountName.Valid {
		query += fmt.Sprintf("bank_account_name = '%s', ", user.BankAccountName.String)
	}

	if user.BankAccountHolder.Valid {
		query += fmt.Sprintf("bank_account_holder = '%s', ", user.BankAccountHolder.String)
	}

	if user.BankAccountNumber.Valid {
		query += fmt.Sprintf("bank_account_number = '%s', ", user.BankAccountNumber.String)
	}

	if user.FileUri.Valid {
		query += fmt.Sprintf("file_uri = '%s', ", user.FileUri.String)
	}

	if user.FileThumbnailUri.Valid {
		query += fmt.Sprintf("file_thumbnail_uri = '%s', ", user.FileThumbnailUri.String)
	}

	query += "updated_at = now() WHERE id = $1"
	query += " RETURNING id, email, phone, file_id, bank_account_name, bank_account_holder, bank_account_number, file_uri, file_thumbnail_uri"

	err_query := database.DB.QueryRow(query, user_id).
		Scan(&user.Id, &user.Email, &user.Phone, &user.FileId, &user.BankAccountName, &user.BankAccountHolder, &user.BankAccountNumber, &user.FileUri, &user.FileThumbnailUri)

	if err_query != nil {
		return user, err_query
	}

	return user, nil
}

func UniqueUser(user User) (ruser User, err error) {
	query := fmt.Sprintf("SELECT id, phone, email FROM %s", DB_NAME)

	conditions := []string{}

	if user.Phone.String != "" {
		conditions = append(conditions, fmt.Sprintf("phone = '%s'", user.Phone.String))
	}

	if user.Email.String != "" {
		conditions = append(conditions, fmt.Sprintf("email = '%s'", user.Email.String))
	}

	if user.Id != "" {
		conditions = append(conditions, fmt.Sprintf("id != '%s'", user.Id))
	}

	// Add conditions to the query
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	err_query := database.DB.QueryRow(query).Scan(&user.Id, &user.Phone, &user.Email)
	if err_query != nil {
		return User{}, err_query
	}

	return user, nil
}
