package purchase_repository

import (
	"fmt"
	"strings"
	"triesdi/app/utils/database"
)

const DB_NAME = "purchases"
const DB_NAME_DETAIL = "purchase_details"
const DB_PRODUCT = "products"

func CreateAndReturn(purchase Purchase, purchaseDetails []PurchaseDetail, totalPrice int) (int, error) {
	// 	WITH new_purchase AS (
	//     INSERT INTO purchases (sender_name, sender_contact_type, sender_contact_detail, total_price)
	//         VALUES ('wahyu', 'email', 'hehe', 100)
	//         RETURNING purchase_id
	// )
	// INSERT INTO purchase_details (
	//     user_id,
	//     purchase_id,
	//     product_id,
	//     name,
	//     category,
	//     quantity,
	//     price,
	//     sku,
	//     file_id,
	//     file_uri,
	//     file_thumbnail_uri
	// )
	// VALUES
	//     (1, (SELECT purchase_id FROM new_purchase), 1, 'name', 'category', 6, 120, 'sku', 'file_id', 'file_uri', 'file_thumbnail_uri'),
	//     (1, (SELECT purchase_id FROM new_purchase), 2, 'name2', 'category2', 8, 222, 'sku2', 'file_id2', 'file_uri2', 'file_thumbnail_uri2')
	// RETURNING purchase_id;

	query := `WITH new_purchase AS (`
	query += fmt.Sprintf("INSERT INTO %s", DB_NAME)
	query += ` (sender_name, sender_contact_type, sender_contact_detail, total_price) `
	query += fmt.Sprintf("VALUES ('%s', '%s', '%s', %d) ", purchase.SenderName, purchase.SenderContactType, purchase.SenderContactDetail, totalPrice)

	query += "RETURNING purchase_id) "

	query += fmt.Sprintf("INSERT INTO %s", DB_NAME_DETAIL)
	query += ` (user_id, purchase_id, product_id, name, category, qty, price, sku, file_id, file_uri, file_thumbnail_uri) VALUES `

	detail_query := []string{}
	for _, pd := range purchaseDetails {
		detail_query = append(detail_query, fmt.Sprintf("('%s', (SELECT purchase_id FROM new_purchase) ,'%s', '%s', '%s', %d, '%d', '%s', '%s', '%s', '%s')", pd.UserId, pd.ProductId, pd.Name, pd.Category, pd.Qty, pd.Price, pd.Sku, pd.FileId, pd.FileUri, pd.FileThumbnailUri))
	}

	query += strings.Join(detail_query, ", ")
	query += " RETURNING purchase_id"

	var purchase_id int
	err := database.DB.QueryRow(query).Scan(&purchase_id)
	if err != nil {
		return 0, err
	}

	return purchase_id, nil
}

func UpdateStock(purchase_id int) error {
	// 	WITH purchase_data AS (
	//     SELECT product_id, qty
	//     FROM purchase_details
	//     WHERE purchase_id = 1
	// )
	// UPDATE products p
	// SET qty = p.qty - pd.qty
	// FROM purchase_data pd
	// WHERE p.product_id = pd.product_id;

	query := `WITH purchase_data AS (`
	query += fmt.Sprintf("SELECT product_id, qty FROM %s WHERE purchase_id = %d", DB_NAME_DETAIL, purchase_id)
	query += `) `
	query += fmt.Sprintf("UPDATE %s p SET qty = p.qty - pd.qty FROM purchase_data pd WHERE p.product_id = pd.product_id", DB_PRODUCT)

	_, err := database.DB.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
