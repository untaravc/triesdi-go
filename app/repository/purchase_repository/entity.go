package purchase_repository

// {
//   "purchaseId": "", // string | Use any id you want
//   "purchasedItems": [
//     {
//       "productId": "",  // string | Use any id you want
//       "name": "",
//       "category": "",
//       "qty": 1, // qty before bought
//       "price": 100,
//       "sku": "",
//       "fileId": "",
//       "fileUri": "", // related file uri
//       "fileThumbnailUri": "", // related file uri
//       "createdAt": ""
//       "updatedAt": ""
//     }
//   ],
//   "totalPrice": 1, // number | should total of all the products bought
//   "paymentDetails": [{ // collection of the seller bank account information that the user bought | if user bought 3 item from 1 seller, then show 1, but if user bought 3 each from different sellers, then show 3, etc
//     "bankAccountName": "",
//     "bankAccountHolder": "",
//     "bankAccountNumber": "",
//     "totalPrice": 1, // number | should total of the bought product that the seller owns
//   }]
// }
type Purchase struct {
	PurchaseId          int    `json:"purchaseId"`
	UserId              int    `json:"userId"`
	SenderName          string `json:"senderName"`
	SenderContactType   string `json:"senderContactType"`
	SenderContactDetail string `json:"senderContactDetail"`
}

type PurchaseDetail struct {
	UserId           string `json:"userId"`
	ProductId        string `json:"productId"`
	Name             string `json:"name"`
	Category         string `json:"category"`
	Qty              int    `json:"qty"`
	Price            int    `json:"price"`
	Sku              string `json:"sku"`
	FileId           string `json:"fileId"`
	FileUri          string `json:"fileUri"`
	FileThumbnailUri string `json:"fileThumbnailUri"`
	CreatedAt        string `json:"createdAt"`
	UpdatedAt        string `json:"updatedAt"`
}
