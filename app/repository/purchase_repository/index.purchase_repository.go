package purchase_repository

import (
	"fmt"
	"triesdi/app/utils/database"
)

const DB_NAME = "purchases"
const DB_NAME_DETAIL = "purchase_details"

func Create(purchase Purchase, purchaseDetails []PurchaseDetail) (Purchase, error) {
	// 	-- Start the transaction
	// START TRANSACTION;

	// -- Step 1: Insert into the purchases table
	// INSERT INTO purchases (buyer_name, purchase_date, total_amount)
	// VALUES ('John Doe', '2025-01-27', 200.00);

	// -- Step 2: Retrieve the last inserted purchase_id
	// -- (This assumes `purchase_id` is an auto-incremented primary key)
	// SET @purchase_id = LAST_INSERT_ID();

	// -- Step 3: Insert into the purchase_details table using the retrieved purchase_id
	// INSERT INTO purchase_details (purchase_id, product_id, quantity, price)
	// VALUES
	//     (@purchase_id, 1, 2, 50.00),
	//     (@purchase_id, 2, 1, 100.00);

	// -- Step 4: Commit the transaction
	// COMMIT;
	totalPrice := 0
	for _, purchaseDetail := range purchaseDetails {
		totalPrice += purchaseDetail.Price
	}

	query := `START TRANSACTION; `
	query += fmt.Sprintf("INSERT INTO %s", DB_NAME)
	query += ` (sender_name, sender_contact_type, sender_contact_detail, total_price) `
	query += fmt.Sprintf("VALUES ('%s', '%s', '%s', %d);", purchase.SenderName, purchase.SenderContactType, purchase.SenderContactDetail, totalPrice)

	query += "SET @purchase_id = LAST_INSERT_ID();"

	query += fmt.Sprintf("INSERT INTO %s", DB_NAME_DETAIL)
	query += ` (purchase_id, user_id, product_id, name, category,quantity, price, sku, file_id, file_uri, file_thumbnail_uri) VALUES `

	for i, pd := range purchaseDetails {
		query += fmt.Sprintf("(@purchase_id, %s,'%s', '%s', '%s', '%d', '%d', '%s', '%s', '%s', '%s')", pd.UserId, pd.ProductId, pd.Name, pd.Category, pd.Qty, pd.Price, pd.Sku, pd.FileId, pd.FileUri, pd.FileThumbnailUri)

		if i < len(purchaseDetails)-1 {
			query += ", "
		}
	}
	query += "; COMMIT;"

	_, err := database.DB.Exec(query)
	if err != nil {
		return purchase, err
	}
	return purchase, nil
}
