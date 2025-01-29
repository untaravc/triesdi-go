package product_repository

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
	"triesdi/app/requests/product_request"
	"triesdi/app/utils/database"
)

const DB_NAME = "products"

func Create(product Product) (Product, error) {
	query := fmt.Sprintf("INSERT INTO %s (user_id, name, category, qty, price, sku, file_id, file_uri, file_thumbnail_uri) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) ", DB_NAME)
	query += "RETURNING product_id, name, category, qty, price, sku, file_id, file_uri, file_thumbnail_uri, created_at, updated_at"

	err := database.DB.QueryRow(query, product.UserId, product.Name, product.Category, product.Qty, product.Price, product.Sku, product.FileId, product.FileUri, product.FileThumbnailUri).
		Scan(
			&product.ProductId,
			&product.Name,
			&product.Category,
			&product.Qty,
			&product.Price,
			&product.Sku,
			&product.FileId,
			&product.FileUri,
			&product.FileThumbnailUri,
			&product.CreatedAt,
			&product.UpdatedAt,
		)

	if err != nil {
		return product, err
	}
	return product, nil
}

func GetAll(filter product_request.ProductFilter) ([]Product, error) {
	selector := "product_id, user_id, name, category, qty, price, sku, file_id, file_uri, file_thumbnail_uri, created_at, updated_at"
	query := fmt.Sprintf("SELECT %s FROM %s", selector, DB_NAME)

	conditions := make([]string, 0)
	if filter.ProductId != "" {
		conditions = append(conditions, fmt.Sprintf("product_id = '%s'", filter.ProductId))
	}

	if filter.Sku != "" {
		conditions = append(conditions, fmt.Sprintf("sku = '%s'", filter.Sku))
	}

	if filter.Category != "" {
		conditions = append(conditions, fmt.Sprintf("category = '%s'", filter.Category))
	}

	if len(filter.ProductIds) > 0 {
		conditions = append(conditions, fmt.Sprintf("product_id IN ('%s')", strings.Join(filter.ProductIds, "','")))
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	if filter.SortBy != "" {
		switch filter.SortBy {
		case "cheapest":
			query += fmt.Sprintf(" ORDER BY %s", "price ASC")
		case "newest":
			query += fmt.Sprintf(" ORDER BY %s", "created_at DESC")
		default:
			pattern := `^sold-(\d+)$`
			re := regexp.MustCompile(pattern)
			matches := re.FindStringSubmatch(filter.SortBy)
			if len(matches) > 1 {
				second, e := strconv.Atoi(matches[1])
				if e == nil {
					query += fmt.Sprintf(" WHERE updated_at > '%s' ORDER BY %s", time.Now().Add(-time.Second*time.Duration(second)).Format("2006-01-02 15:04:05"), "updated_at DESC")
				}
			}
		}
	}

	fmt.Println(query)

	if filter.Limit != 0 {
		query += fmt.Sprintf(" LIMIT %d", filter.Limit)
	}

	if filter.Offset != 0 {
		query += fmt.Sprintf(" OFFSET %d", filter.Offset)
	}

	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]Product, 0)
	for rows.Next() {
		var qty, price int
		var product_id, user_id, name, category, sku, file_id, file_uri, file_thumbnail_uri, created_at, updated_at string
		if err := rows.Scan(&product_id, &user_id, &name, &category, &qty, &price, &sku, &file_id, &file_uri, &file_thumbnail_uri, &created_at, &updated_at); err != nil {
			return nil, err
		}

		products = append(products, Product{
			ProductId:        product_id,
			UserId:           user_id,
			Name:             name,
			Category:         category,
			Qty:              qty,
			Price:            price,
			Sku:              sku,
			FileId:           file_id,
			FileUri:          file_uri,
			FileThumbnailUri: file_thumbnail_uri,
			CreatedAt:        created_at,
			UpdatedAt:        updated_at,
		})
	}

	return products, nil
}

func Update(product Product) (Product, error) {
	query := fmt.Sprintf("UPDATE %s SET", DB_NAME)

	conditions := make([]string, 0)
	if product.Name != "" {
		conditions = append(conditions, fmt.Sprintf("name = '%s'", product.Name))
	}

	if product.Category != "" {
		conditions = append(conditions, fmt.Sprintf("category = '%s'", product.Category))
	}

	if product.Qty != 0 {
		conditions = append(conditions, fmt.Sprintf("qty = %d", product.Qty))
	}

	if product.Price != 0 {
		conditions = append(conditions, fmt.Sprintf("price = %d", product.Price))
	}

	if product.Sku != "" {
		conditions = append(conditions, fmt.Sprintf("sku = '%s'", product.Sku))
	}

	if product.FileId != "" {
		conditions = append(conditions, fmt.Sprintf("file_id = '%s'", product.FileId))
	}

	if product.FileUri != "" {
		conditions = append(conditions, fmt.Sprintf("file_uri = '%s'", product.FileUri))
	}

	if product.FileThumbnailUri != "" {
		conditions = append(conditions, fmt.Sprintf("file_thumbnail_uri = '%s'", product.FileThumbnailUri))
	}

	if len(conditions) > 0 {
		query += " " + strings.Join(conditions, ", ")
	}

	query += fmt.Sprintf(" WHERE product_id = '%s'", product.ProductId)
	query += " RETURNING product_id, name, category, qty, price, sku, file_id, file_uri, file_thumbnail_uri"

	err := database.DB.QueryRow(query).Scan(&product.ProductId, &product.Name, &product.Category, &product.Qty, &product.Price, &product.Sku, &product.FileId, &product.FileUri, &product.FileThumbnailUri)
	if err != nil {
		return product, err
	}

	return product, nil
}

func Delete(product_id string) (bool, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE product_id = $1", DB_NAME)
	result, err := database.DB.Exec(query, product_id)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	if rowsAffected == 0 {
		return false, nil
	}

	return true, nil
}
