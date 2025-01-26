package product_repository

import (
	"fmt"
	"triesdi/app/utils/database"
)

const DB_NAME = "products"

func Create(product Product) (Product, error) {
	query := fmt.Sprintf("INSERT INTO %s (user_id, name, category, qty, price, sku, file_id, file_uri, file_thumbnail_uri) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING product_id, name, category, qty, price, sku, file_id, file_uri, file_thumbnail_uri", DB_NAME)

	err := database.DB.QueryRow(query, product.UserId, product.Name, product.Category, product.Qty, product.Price, product.Sku, product.FileId, product.FileUri, product.FileThumbnailUri).
		Scan(&product.ProductId, &product.Name, &product.Category, &product.Qty, &product.Price, &product.Sku, &product.FileId, &product.FileUri, &product.FileThumbnailUri)

	if err != nil {
		return product, err
	}
	return product, nil
}
