package handler

import "github.com/gin-gonic/gin"

type successBody struct {
	Success string `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type errorBody struct {
	Success string `json:"success"`
	Message string `json:"message"`
}

// success response
func successResponse(c *gin.Context, status int, message string, data any) {
	c.JSON(status, successBody{
		Success: "OK",
		Message: message,
		Data:    data,
	})
}

// error response
func errorResponse(c *gin.Context, status int, message string) {
	c.JSON(status, errorBody{
		Success: "NG",
		Message: message,
	})
}
