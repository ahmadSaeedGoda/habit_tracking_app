package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"lasting-dynamics.com/habit_tracking_app/internal/api/auth"
	"lasting-dynamics.com/habit_tracking_app/internal/api/common"
	"lasting-dynamics.com/habit_tracking_app/internal/api/services"
)

// UserHandler handles HTTP requests related to users
type UserHandler struct {
	*BaseHandler
	userService *services.UserService
}

// NewUserHandler creates a new instance of UserHandler
func NewUserHandler() *UserHandler {
	return &UserHandler{userService: services.NewUserService()}
}

// Register handles user registration request
func (uh *UserHandler) Register(c *gin.Context) {
	var request struct {
		Name     string `json:"name" binding:"required" validate:"required"`
		Email    string `json:"email" binding:"required,email" validate:"required,email"`
		Password string `json:"password" binding:"required,min=8" validate:"required,min=8"`
	}

	// TODO:: Introduce a generic validation function to be used across the app handlers
	// for proper handling, separation of concerns, avoiding code smells like long functions
	// and violation of single responsibility, reusability, and maintainability.
	if err := c.ShouldBindJSON(&request); err != nil { // Handle validation error
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			// Custom error messages based on validation errors
			errorMessages := make(map[string]string)
			for _, e := range validationErrors {
				switch e.Tag() {
				case "required":
					errorMessages[e.Field()] = e.Field() + " is required"
				case "email":
					errorMessages[e.Field()] = "Invalid email format"
				case "min":
					errorMessages[e.Field()] = e.Field() + " must be at least " + e.Param() + " characters"
				default:
					errorMessages[e.Field()] = e.Field() + " is invalid"
				}
			}
			c.JSON(http.StatusBadRequest, gin.H{"error": errorMessages})
			return
			// common.GenerateErrorResponse(c, http.StatusBadRequest, errors.New(errorMessages))
			// return
		}

	}

	if emailExists := uh.userService.EmailExists(request.Email); emailExists {
		common.GenerateErrorResponse(c, http.StatusConflict, errors.New("email address is already registered"))
		return
	}

	// Register user
	if err := uh.userService.Register(request.Name, request.Email, request.Password); err != nil {
		common.GenerateErrorResponse(c, http.StatusConflict, errors.New("registration failed. please try again later"))
		return
	}

	c.JSON(201, gin.H{"message": "User registered successfully"})
}

// Login handles user login request
func (uh *UserHandler) Login(c *gin.Context) {
	var request struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		common.GenerateErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	user, err := uh.userService.AuthenticateUser(request.Email, request.Password)
	if err != nil {
		common.GenerateErrorResponse(c, http.StatusUnauthorized, errors.New("incorrect credentials"))
		return
	}

	// Generate JWT token (you need to implement this part)
	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		common.GenerateErrorResponse(c, http.StatusInternalServerError, errors.New("something went wrong while processing your request"))
		return
	}

	common.GenerateSuccessResponse(c, http.StatusOK, "Success", gin.H{"token": token})
}
