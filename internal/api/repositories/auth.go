package repositories

import (
	"gorm.io/gorm"
	"lasting-dynamics.com/habit_tracking_app/internal/api/models"
	DBManager "lasting-dynamics.com/habit_tracking_app/internal/database/dbmanager"
)

// UserRepository handles database operations related to users
type UserRepository struct {
	*BaseRepository
	db *gorm.DB
}

// NewUserRepository creates a new instance of UserRepository
// Uses DBManager.DB for connection to data store
func NewUserRepository() *UserRepository {
	return &UserRepository{db: DBManager.DB}
}

// Create is for new user record persistence
func (ur *UserRepository) Create(user *models.User) error {
	return ur.db.Create(user).Error
}

// FindByEmail is finding the first user record by email for auth purposes
func (ur *UserRepository) FindByEmail(email string) (*models.User, error) {
	user := &models.User{}
	err := ur.db.Where("email = ?", email).First(user).Error
	return user, err
}

// EmailExists verifies whether a user record exists having the specified email address.
func (ur *UserRepository) EmailExists(email string) bool {
	var count int64
	ur.db.Model(&models.User{}).Where("email = ?", email).Count(&count)
	return count > 0
}
