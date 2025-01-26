package product_controller

import (
	"net/http"
	"triesdi/app/repository/file_repository"
	"triesdi/app/repository/product_repository"
	"triesdi/app/requests/product_request"
	"triesdi/app/utils/common"
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

	common.ConsoleLog(auth_user)

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
