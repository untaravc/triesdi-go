package product_request

import (
	"triesdi/app/utils/converter"

	"github.com/gin-gonic/gin"
)

type ProductRequest struct {
	Name     string `json:"name" validate:"required,min=4,max=32"`
	Category string `json:"category" validate:"required,oneof=Food Beverage Clothes Furniture Tools"`
	Qty      int    `json:"qty" validate:"required,min=1"`
	Price    int    `json:"price" validate:"required,min=100"`
	Sku      string `json:"sku" validate:"required,min=0,max=32"`
	FileId   string `json:"fileId" validate:"required"`
}

type ProductUpdateRequest struct {
	Name     string `json:"name" validate:"omitempty,min=4,max=32"`
	Category string `json:"category" validate:"omitempty,oneof=Food Beverage Clothes Furniture Tools"`
	Qty      int    `json:"qty" validate:"omitempty,min=1"`
	Price    int    `json:"price" validate:"omitempty,min=100"`
	Sku      string `json:"sku" validate:"omitempty,min=0,max=32"`
	FileId   string `json:"fileId" validate:"omitempty"`
}

type ProductFilter struct {
	Limit      int      `json:"limit"`
	Offset     int      `json:"offset"`
	ProductId  string   `json:"productId"`
	Sku        string   `json:"sku"`
	Category   string   `json:"category"`
	SortBy     string   `json:"sortBy"`
	ProductIds []string `json:"productIds"`
}

func FilterToProductFilter(c *gin.Context) ProductFilter {
	product_filter := ProductFilter{}

	product_filter.Limit = 10
	product_filter.Offset = 0

	if c.Query("limit") != "" {
		product_filter.Limit = converter.StringToInt(c.Query("limit"))
	}
	if c.Query("offset") != "" {
		product_filter.Offset = converter.StringToInt(c.Query("offset"))
	}
	if c.Query("productId") != "" {
		product_filter.ProductId = c.Query("productId")
	}
	if c.Query("sku") != "" {
		product_filter.Sku = c.Query("sku")
	}
	if c.Query("category") != "" {
		product_filter.Category = c.Query("category")
	}
	if c.Query("sortBy") != "" {
		product_filter.SortBy = c.Query("sortBy")
	}

	return product_filter
}
