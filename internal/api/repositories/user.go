package repositories

import (
	"lasting-dynamics.com/habit_tracking_app/internal/api/models"
)

// UserRepositoryInterface defines the interface for interacting with user data
type UserRepositoryInterface interface {
	GetUserByID(userID uint) (*models.User, error)
}

// GetUserByID retrieves a user by its ID
func (ur *UserRepository) GetUserByID(userID uint) (*models.User, error) {
	var user models.User
	if err := ur.db.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
