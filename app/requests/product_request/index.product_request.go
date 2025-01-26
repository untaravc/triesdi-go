package product_request

type ProductRequest struct {
	Name     string `json:"name" validate:"required,min=4,max=32"`                                    // string | required | minLength: 4 | maxLength: 32
	Category string `json:"category" validate:"required,oneof=Food Beverage Clothes Furniture Tools"` // string | required | should be enum of product category types table
	Qty      int    `json:"qty" validate:"required,min=1"`                                            // number | required | min: 1
	Price    int    `json:"price" validate:"required,min=100"`                                        // number | required | min: 100
	Sku      string `json:"sku" validate:"required,min=0,max=32"`                                     // string | required | minLength: 0 | maxLength: 32
	FileId   string `json:"fileId" validate:"required"`                                               // string | required | should be a valid fileId
}
