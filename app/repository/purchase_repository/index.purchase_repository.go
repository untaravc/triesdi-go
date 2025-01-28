package purchase_repository

import (
	"fmt"
	"triesdi/app/utils/common"
	"triesdi/app/utils/database"
)

const DB_NAME = "purchases"
const DB_NAME_DETAIL = "purchase_details"

func Create(purchase Purchase, purchaseDetails []PurchaseDetail) (Purchase, error) {
	// 	DO $$
	// DECLARE
	//     _purchase_id BIGINT; -- Declare the variable to hold the purchase_id
	// BEGIN
	//     -- Insert into purchases and get the generated purchase_id
	//     INSERT INTO purchases (sender_name, sender_contact_type, sender_contact_detail, total_price)
	//     VALUES ('Wahyi', 'phone', '+2929222222', 100)
	//     RETURNING purchase_id INTO _purchase_id;

	//     -- Insert into purchase_details using the retrieved purchase_id
	//     INSERT INTO purchase_details (purchase_id, user_id, product_id, name, category, quantity, price, sku, file_id, file_uri, file_thumbnail_uri)
	//     VALUES (_purchase_id, '2', '1', 'Ciduk', 'Food', '2', '100', '123123', '1',
	//             'https://projectsprint-bucket-public-read.s3.ap-southeast-1.amazonaws.com/uploads/fVcg5F6PeEsxEr0OYOOgtGX2ADRjW0YpKhXvl6FubJUSrf90T600uuwhtkDD.png',
	//             'https://projectsprint-bucket-public-read.s3.ap-southeast-1.amazonaws.com/uploads/thumbnail/eIilUx8I93K9B8JPyqbZZwPIhCCDUR8RVV21UZcQKXzixDDhKwFOmzi1Czkw.png');
	// END $$;

	totalPrice := 0
	for _, purchaseDetail := range purchaseDetails {
		totalPrice += purchaseDetail.Price
	}

	query := `DO $$ DECLARE _purchase_id BIGINT; BEGIN `
	query += fmt.Sprintf("INSERT INTO %s", DB_NAME)
	query += ` (sender_name, sender_contact_type, sender_contact_detail, total_price) `
	query += fmt.Sprintf("VALUES ('%s', '%s', '%s', %d) ", purchase.SenderName, purchase.SenderContactType, purchase.SenderContactDetail, totalPrice)

	query += "RETURNING purchase_id INTO _purchase_id; "

	query += fmt.Sprintf("INSERT INTO %s", DB_NAME_DETAIL)
	query += ` (purchase_id, user_id, product_id, name, category,quantity, price, sku, file_id, file_uri, file_thumbnail_uri) VALUES `

	for i, pd := range purchaseDetails {
		query += fmt.Sprintf("(_purchase_id, '%s','%s', '%s', '%s', '%d', '%d', '%s', '%s', '%s', '%s')", pd.UserId, pd.ProductId, pd.Name, pd.Category, pd.Qty, pd.Price, pd.Sku, pd.FileId, pd.FileUri, pd.FileThumbnailUri)

		if i < len(purchaseDetails)-1 {
			query += ", "
		}
	}
	query += "; COMMIT; END $$;"
	common.ConsoleLog(query)
	_, err := database.DB.Exec(query)
	if err != nil {
		return purchase, err
	}
	return purchase, nil
}
