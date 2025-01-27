package product_repository

type Product struct {
	ProductId        string `json:"product_id"`
	UserId           string `json:"user_id"`
	Name             string `json:"name"`
	Category         string `json:"category"`
	Qty              int    `json:"qty"`
	Price            int    `json:"price"`
	Sku              string `json:"sku"`
	FileId           string `json:"file_id"`
	FileUri          string `json:"file_uri"`
	FileThumbnailUri string `json:"file_thumbnail_uri"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
}
