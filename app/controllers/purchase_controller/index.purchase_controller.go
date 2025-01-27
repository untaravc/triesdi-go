package purchase_controller

import (
	"net/http"
	"triesdi/app/repository/product_repository"
	"triesdi/app/repository/purchase_repository"
	"triesdi/app/requests/product_request"
	"triesdi/app/requests/purchase_request"
	"triesdi/app/utils/validator"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	purchase_request := purchase_request.PurchaseRequest{}

	if err := c.ShouldBindJSON(&purchase_request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate type of contact
	if purchase_request.SenderContactType == "email" {
		purchase_request.SenderContactDetailEmail = purchase_request.SenderContactDetail
	} else if purchase_request.SenderContactType == "phone" {
		purchase_request.SenderContactDetailPhone = purchase_request.SenderContactDetail
	}

	if err := validator.ValidateStruct(purchase_request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validator.FormatValidationError(err)})
		return
	}

	// Validate Items
	for _, item := range purchase_request.PurchasedItems {
		if item.Qty < 2 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "qty must be more than 2"})
			return
		}
	}

	// Validate products
	var productIds []string
	for _, item := range purchase_request.PurchasedItems {
		productIds = append(productIds, item.ProductId)
	}

	// Get products
	products, err := product_repository.GetAll(product_request.ProductFilter{ProductIds: productIds})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	purchase_details := []purchase_repository.PurchaseDetail{}
	for _, purchaseItem := range purchase_request.PurchasedItems {
		valid := false
		for _, product := range products {
			if product.ProductId == purchaseItem.ProductId {
				valid = true

				purchase_details = append(purchase_details, purchase_repository.PurchaseDetail{
					ProductId:        purchaseItem.ProductId,
					Qty:              purchaseItem.Qty,
					Name:             product.Name,
					Category:         product.Category,
					Price:            product.Price,
					Sku:              product.Sku,
					FileId:           product.FileId,
					FileUri:          product.FileUri,
					FileThumbnailUri: product.FileThumbnailUri,
				})
				break
			}
		}

		if !valid {
			c.JSON(http.StatusBadRequest, gin.H{"error": "product not found"})
			return
		}
	}

	purchase := purchase_repository.Purchase{
		SenderName:          purchase_request.SenderName,
		SenderContactType:   purchase_request.SenderContactType,
		SenderContactDetail: purchase_request.SenderContactDetail,
	}

	// Create purchase
	rpurchase, err := purchase_repository.Create(purchase, purchase_details)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, rpurchase)
}
