package handlers

import (
	"github.com/gin-gonic/gin"
)

type ApiResponse struct {
	Success bool                   `json:"success"`
	Code    string                 `json:"code"`
	Status  int                    `json:"status"`
	Result  map[string]interface{} `json:"result"`
}

func GenerateResponse(ctx *gin.Context, data interface{}, code string, status int) {
	response := ApiResponse{
		Success: status >= 200 && status < 300,
		Code:    code,
		Status:  status,
		Result:  map[string]interface{}{"data": data},
	}
	ctx.JSON(status, response)
}
