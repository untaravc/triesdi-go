package purchase_controller

import (
	"net/http"
	"strconv"
	"sync"
	"triesdi/app/repository/file_repository"
	"triesdi/app/repository/product_repository"
	"triesdi/app/repository/purchase_repository"
	"triesdi/app/repository/user_repository"
	"triesdi/app/requests/product_request"
	"triesdi/app/requests/purchase_request"
	"triesdi/app/responses/purchase_response"
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
			c.JSON(http.StatusBadRequest, gin.H{"error": "qty must be more than 3"})
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
	total_price := 0
	for _, purchaseItem := range purchase_request.PurchasedItems {
		valid := false
		for _, product := range products {
			if product.ProductId == purchaseItem.ProductId {
				valid = true

				purchase_details = append(purchase_details, purchase_repository.PurchaseDetail{
					UserId:           product.UserId,
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

				total_price += product.Price * purchaseItem.Qty
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

	var wg sync.WaitGroup
	wg.Add(2)

	var inserted_id string
	var users []user_repository.User

	go func() {
		defer wg.Done()
		// Create purchase
		id, err := purchase_repository.CreateAndReturn(purchase, purchase_details, total_price)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		inserted_id = strconv.Itoa(id)
	}()

	go func() {
		defer wg.Done()
		// Get users
		var ids []string
		for _, purchase_detail := range purchase_details {
			for _, id := range ids {
				if id == purchase_detail.UserId {
					continue
				}
			}
			ids = append(ids, purchase_detail.UserId)
		}
		users, _ = user_repository.GetUsers(user_repository.UserFilter{Ids: ids})
	}()

	wg.Wait()

	payment_details := []purchase_response.PaymentDetail{}

	for _, user := range users {
		total_price_user := 0

		for _, purchase_detail := range purchase_details {
			if user.Id == purchase_detail.UserId {
				total_price_user += purchase_detail.Price * purchase_detail.Qty

				payment_details = append(payment_details, purchase_response.PaymentDetail{
					BankAccountName:   user.BankAccountName.String,
					BankAccountHolder: user.BankAccountHolder.String,
					BankAccountNumber: user.BankAccountNumber.String,
					TotalPrice:        total_price_user,
				})
			}
		}
	}

	response := purchase_response.PurchaseResponse{
		PurchaseId:     inserted_id,
		PurchasedItems: purchase_details,
		TotalPrice:     total_price,
		PaymentDetails: payment_details,
	}

	c.JSON(http.StatusCreated, response)
}

func Update(c *gin.Context) {
	purchase_update_request := purchase_request.PurchaseUpdateRequest{}
	purchaseId := c.Param("purchaseId")

	if err := c.ShouldBindJSON(&purchase_update_request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validator.ValidateStruct(purchase_update_request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validator.FormatValidationError(err)})
		return
	}

	count := len(purchase_update_request.FileIds)

	// Get file ids
	files, err := file_repository.GetAll(file_repository.FileFilter{FileIds: purchase_update_request.FileIds})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if count != len(files) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file not found"})
		return
	}

	for f, file := range files {
		valid := false

		for _, file_id := range purchase_update_request.FileIds {
			id, _ := strconv.Atoi(file_id)
			if file.FileId == id {
				valid = true
				break
			}
		}

		if !valid {
			c.JSON(http.StatusBadRequest, gin.H{"error": "file [" + strconv.Itoa(f) + "]not valid"})
			return
		}
	}

	id, _ := strconv.Atoi(purchaseId)
	err_update := purchase_repository.UpdateStock(id)

	if err_update != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err_update.Error()})
		return
	}

	c.JSON(http.StatusCreated, files)
}
