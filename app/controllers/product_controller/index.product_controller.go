package product_controller

import (
	"net/http"
	"triesdi/app/repository/file_repository"
	"triesdi/app/repository/product_repository"
	"triesdi/app/requests/product_request"
	"triesdi/app/utils/jwt"
	"triesdi/app/utils/validator"

	"github.com/gin-gonic/gin"
)

func Store(c *gin.Context) {
	auth_user := jwt.GetAuth(c)
	product_request := product_request.ProductRequest{}

	if err := c.ShouldBindJSON(&product_request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validator.ValidateStruct(product_request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validator.FormatValidationError(err)})
		return
	}

	file, err := file_repository.GetById(product_request.FileId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := product_repository.Product{
		Name:             product_request.Name,
		UserId:           auth_user.Id,
		Category:         product_request.Category,
		Qty:              product_request.Qty,
		Price:            product_request.Price,
		Sku:              product_request.Sku,
		FileId:           product_request.FileId,
		FileUri:          file.FileUri,
		FileThumbnailUri: file.FileThumbnailUri,
	}

	rproduct, err_product := product_repository.Create(product)

	if err_product != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err_product.Error()})
		return
	}

	c.JSON(http.StatusOK, rproduct)
}

func GetAll(c *gin.Context) {
	filter := product_request.FilterToProductFilter(c)

	products, err := product_repository.GetAll(filter)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

func Update(c *gin.Context) {
	product_update_request := product_request.ProductUpdateRequest{}

	if err := c.ShouldBindJSON(&product_update_request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validator.ValidateStruct(product_update_request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validator.FormatValidationError(err)})
		return
	}

	product := product_repository.Product{
		ProductId: c.Param("product_id"),
		Name:      product_update_request.Name,
		Category:  product_update_request.Category,
		Qty:       product_update_request.Qty,
		Price:     product_update_request.Price,
		Sku:       product_update_request.Sku,
	}

	if product_update_request.FileId != "" {
		file, err := file_repository.GetById(product_update_request.FileId)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		product.FileId = product_update_request.FileId
		product.FileUri = file.FileUri
		product.FileThumbnailUri = file.FileThumbnailUri
	}

	rproduct, err_product := product_repository.Update(product)

	if err_product != nil {
		if err_product.Error() == "sql: no rows in result set" {
			c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err_product.Error()})
		return
	}

	c.JSON(http.StatusOK, rproduct)
}

func Delete(c *gin.Context) {
	productId := c.Param("product_id")

	status, err := product_repository.Delete(productId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !status {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product deleted"})
}
