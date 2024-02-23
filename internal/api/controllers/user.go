package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"lasting-dynamics.com/habit_tracking_app/internal/api/common"
)

// GetUserByID retrieves a user by its ID
func (uh *UserHandler) GetUserByID(c *gin.Context) {
	// Extract user ID from the context
	userId, exists := c.Get("userId")
	if !exists {
		common.GenerateErrorResponse(c, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	// Parse and validate the user ID
	parsedUserID, ok := userId.(uint)
	if !ok {
		common.GenerateErrorResponse(c, http.StatusInternalServerError, errors.New("failed to convert userId"))
		return
	}

	userDTO, err := uh.userService.GetUserByID(uint(parsedUserID))
	if err != nil {
		common.GenerateErrorResponse(c, http.StatusNotFound, errors.New("user not found"))
		return
	}

	// Send the user DTO in the response
	common.GenerateSuccessResponse(c, http.StatusOK, "Success", userDTO)
}
