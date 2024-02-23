package common

import (
	"github.com/gin-gonic/gin"
)

// JSONResponse represents a standardized JSON response format
type JSONResponse struct {
	Message string      `json:"message"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
}

// GenerateErrorResponse sends an error response with the given status code and error message
func GenerateErrorResponse(c *gin.Context, statusCode int, err error) {
	response := JSONResponse{
		Status:  "error",
		Message: err.Error(),
		Data:    nil,
	}
	GenerateJSONResponse(c, statusCode, response)
}

// GenerateSuccessResponse sends a success response with the given status code, message, and data
func GenerateSuccessResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	response := JSONResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	}
	GenerateJSONResponse(c, statusCode, response)
}

// GenerateJSONResponse generates a JSON response with the given status code and data
func GenerateJSONResponse(c *gin.Context, statusCode int, response JSONResponse) {
	c.Header("Content-Type", "application/json")
	c.JSON(statusCode, response)
}
