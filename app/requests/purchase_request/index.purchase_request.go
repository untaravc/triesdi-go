package purchase_request

// {
//   "purchasedItems": [ // array | minItems: 1
//     {
//       "productId": "", // string | should a valid productId
//       "qty": 1, // number | min: 2
//     },
//   ],
//   "senderName": "", // string | required | minLength: 4 | maxLength: 55
//   "senderContactType": "", // string | required | enum of "email" / "phone"
//   "senderContactDetail": "", // string | required | if "phone" then validates the phone number | if "email" then validates email
// }

type PurchaseRequest struct {
	PurchasedItems           []PurchasedItem `json:"purchasedItems" validate:"required"`                      // array | minItems: 1
	SenderName               string          `json:"senderName" validate:"required,min=4,max=55"`             // string | required | minLength: 4 | maxLength: 55
	SenderContactType        string          `json:"senderContactType" validate:"required,oneof=email phone"` // string | required | enum of "email" / "phone"
	SenderContactDetail      string          `json:"senderContactDetail" validate:"required"`                 // string | required | if "phone" then validates the phone number | if "email" then validates email
	SenderContactDetailEmail string          `json:"senderContactDetailEmail" validate:"omitempty,email"`     // string | required | if "phone" then validates the phone number | if "email" then validates email
	SenderContactDetailPhone string          `json:"senderContactDetailPhone" validate:"omitempty,e164"`      // string | required | if "phone" then validates the phone number | if "email" then validates email
}

type PurchasedItem struct {
	ProductId string `json:"productId" validate:"required"` // string | should a valid productId
	Qty       int    `json:"qty" validate:"required,min=2"` // number | min: 2
}

type PurchaseUpdateRequest struct {
	FileIds []string `json:"fileIds" validate:"required"`
}
