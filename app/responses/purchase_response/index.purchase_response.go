package purchase_response

import "triesdi/app/repository/purchase_repository"

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

type PurchaseResponse struct {
	PurchaseId     string                               `json:"purchaseId"`
	PurchasedItems []purchase_repository.PurchaseDetail `json:"purchasedItems"`
	TotalPrice     int                                  `json:"totalPrice"`
	PaymentDetails []PaymentDetail                      `json:"paymentDetails"`
}

type PaymentDetail struct {
	BankAccountName   string `json:"bankAccountName"`
	BankAccountHolder string `json:"bankAccountHolder"`
	BankAccountNumber string `json:"bankAccountNumber"`
	TotalPrice        int    `json:"totalPrice"`
}
