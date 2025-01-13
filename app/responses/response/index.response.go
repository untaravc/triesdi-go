package response

import (
	"github.com/gin-gonic/gin"
)

func BaseResponse(ctx *gin.Context, httpStatus int, success bool, message string, result interface{}) {
	// ctx.AbortWithStatusJSON(httpStatus, TypeResponse{
	// 	Success: success,
	// 	Message: message,
	// 	Result:  result,
	// })
	ctx.AbortWithStatusJSON(httpStatus, result)
}

func UploadResponse(ctx *gin.Context, httpStatus int, result string) {
	ctx.AbortWithStatusJSON(httpStatus, TypeUploadResponse{
		Result: result,
	})
}

type TypeResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

type TypeUploadResponse struct {
	Result string `json:"uri"`
}

type TypePagination struct {
	Total   int `json:"total"`
	From    int `json:"from"`
	To      int `json:"to"`
	PerPage int `json:"per_page"`
	Page    int `json:"page"`
}

type ErrorMessage struct {
	Field string `json:"field"`
	Msg   string `json:"msg"`
}

type AuthResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}
