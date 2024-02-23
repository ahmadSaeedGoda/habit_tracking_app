package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"lasting-dynamics.com/habit_tracking_app/internal/api/common"
	dto "lasting-dynamics.com/habit_tracking_app/internal/api/entities"
	"lasting-dynamics.com/habit_tracking_app/internal/api/services"
)

// HabitHandler handles HTTP requests related to habits
type HabitHandler struct {
	*BaseHandler
	habitService services.HabitServiceInterface
}

// NewHabitHandler creates a new instance of HabitHandler
func NewHabitHandler() *HabitHandler {
	return &HabitHandler{
		habitService: services.NewHabitService(),
	}
}

// ListHabits retrieves a list of habits for the authenticated user
func (hh *HabitHandler) ListHabits(c *gin.Context) {
	// Extract the user ID from the JWT token
	userId, exists := c.Get("userId")
	if !exists {
		common.GenerateErrorResponse(c, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	// Direct type assertion assuming userId is of type uint
	uintUserID, ok := userId.(uint)
	if !ok {
		common.GenerateErrorResponse(c, http.StatusInternalServerError, errors.New("failed to Parse userId to Proper type"))
		return
	}

	habitDTOs, err := hh.habitService.ListByUser(uintUserID)
	if err != nil {
		common.GenerateErrorResponse(c, http.StatusNotFound, errors.New("no Habits found"))
		return
	}

	// Send the list of habit DTOs in the response
	common.GenerateSuccessResponse(c, http.StatusOK, "Success", habitDTOs)
}

// GetHabitByID retrieves a habit by its ID
func (hh *HabitHandler) GetHabitByID(c *gin.Context) {
	// Extract user ID from the context
	userId, exists := c.Get("userId")
	if !exists {
		common.GenerateErrorResponse(c, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	habitID, err := strconv.ParseUint(c.Param("habitId"), 10, 64)
	if err != nil {
		common.GenerateErrorResponse(c, http.StatusBadRequest, errors.New("invalid habit ID"))
		return
	}

	// Check if the habit belongs to the user
	habitDTO, err := hh.habitService.GetHabitByID(uint(habitID))
	if err != nil {
		common.GenerateErrorResponse(c, http.StatusNotFound, errors.New("habit not found"))
		return
	}

	// Verify ownership by comparing the user ID in the context and the habit's user ID
	if userId != habitDTO.UserId {
		common.GenerateErrorResponse(c, http.StatusForbidden, errors.New("unauthorized to access this habit"))
		return
	}

	// Send the habit DTO in the response
	common.GenerateSuccessResponse(c, http.StatusOK, "Success", habitDTO)
}

// CreateHabit creates a new habit
func (hh *HabitHandler) CreateHabit(c *gin.Context) {
	var habitRequest struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&habitRequest); err != nil {
		common.GenerateErrorResponse(c, http.StatusBadRequest, errors.New("invalid input"))
		return
	}

	// Extract user ID from the context
	userId, exists := c.Get("userId")
	if !exists {
		common.GenerateErrorResponse(c, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	// Parse and validate the user ID
	parsedUserID, ok := userId.(uint)
	if !ok {
		common.GenerateErrorResponse(c, http.StatusInternalServerError, errors.New("failed to convert userId to uint"))
		return
	}

	// Check if the habit name is unique for the user
	if !hh.habitService.IsHabitNameUnique(parsedUserID, habitRequest.Name) {
		common.GenerateErrorResponse(c, http.StatusConflict, errors.New("a habit with this name already exists"))
		return
	}

	// Create the habit
	if err := hh.habitService.CreateHabit(parsedUserID, &dto.HabitDTO{
		Name:        habitRequest.Name,
		Description: habitRequest.Description,
		UserId:      parsedUserID,
	}); err != nil {
		common.GenerateErrorResponse(c, http.StatusInternalServerError, errors.New("failed to create habit"))
		return
	}

	common.GenerateSuccessResponse(c, http.StatusCreated, "Habit created successfully", nil)
}

// UpdateHabit updates an existing habit
func (hh *HabitHandler) UpdateHabit(c *gin.Context) {
	habitID, err := strconv.ParseUint(c.Param("habitId"), 10, 64)
	if err != nil {
		common.GenerateErrorResponse(c, http.StatusBadRequest, errors.New("invalid habit ID"))
		return
	}

	// Extract user ID from the context
	userID, exists := c.Get("userId")
	if !exists {
		common.GenerateErrorResponse(c, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	// Parse and validate the user ID
	parsedUserID, ok := userID.(uint)
	if !ok {
		common.GenerateErrorResponse(c, http.StatusInternalServerError, errors.New("failed to convert userId to uint"))
		return
	}

	// Check if the habit exists
	habitDTO, err := hh.habitService.GetHabitByID(uint(habitID))
	if err != nil {
		common.GenerateErrorResponse(c, http.StatusNotFound, errors.New("habit not found"))
		return
	}

	// Verify ownership by comparing the user ID in the context and the habit's user ID
	if parsedUserID != habitDTO.UserId {
		common.GenerateErrorResponse(c, http.StatusForbidden, errors.New("unauthorized to access this habit"))
		return
	}

	var habitRequest struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&habitRequest); err != nil {
		common.GenerateErrorResponse(c, http.StatusBadRequest, errors.New("invalid input"))
		return
	}

	// Ensure uniqueness of the habit name across user habits
	if habitRequest.Name != "" && !hh.habitService.IsHabitNameUnique(parsedUserID, habitRequest.Name) {
		common.GenerateErrorResponse(c, http.StatusConflict, errors.New("a habit with this name already exists. let your habits be unique"))
		return
	}

	// Update the habit
	if err := hh.habitService.UpdateHabit(parsedUserID, uint(habitID), &dto.HabitDTO{
		Name:        habitRequest.Name,
		Description: habitRequest.Description,
		UserId:      parsedUserID,
	}); err != nil {
		common.GenerateErrorResponse(c, http.StatusInternalServerError, errors.New("failed to update habit"))
		return
	}

	common.GenerateSuccessResponse(c, http.StatusOK, "Habit updated successfully", nil)
}

// DeleteHabit deletes a habit by its ID
func (hh *HabitHandler) DeleteHabit(c *gin.Context) {
	// Extract user ID from the context
	userId, exists := c.Get("userId")
	if !exists {
		common.GenerateErrorResponse(c, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	// Parse and validate the user ID
	parsedUserID, ok := userId.(uint)
	if !ok {
		common.GenerateErrorResponse(c, http.StatusInternalServerError, errors.New("failed to convert userId to uint"))
		return
	}

	habitID, err := strconv.ParseUint(c.Param("habitId"), 10, 64)
	if err != nil {
		common.GenerateErrorResponse(c, http.StatusBadRequest, errors.New("invalid habit ID"))
		return
	}

	// Check if the habit exists
	habitDTO, err := hh.habitService.GetHabitByID(uint(habitID))
	if err != nil {
		common.GenerateErrorResponse(c, http.StatusNotFound, errors.New("habit not found"))
		return
	}

	// Verify ownership by comparing the user ID in the context and the habit's user ID
	if parsedUserID != habitDTO.UserId {
		common.GenerateErrorResponse(c, http.StatusForbidden, errors.New("unauthorized to access this habit"))
		return
	}

	// Delete the habit
	if err := hh.habitService.DeleteHabit(uint(parsedUserID), uint(habitID)); err != nil {
		common.GenerateErrorResponse(c, http.StatusInternalServerError, errors.New("failed to delete habit"))
		return
	}

	common.GenerateSuccessResponse(c, http.StatusOK, "Habit deleted successfully", nil)
}
