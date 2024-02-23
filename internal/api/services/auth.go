package services

import (
	"lasting-dynamics.com/habit_tracking_app/internal/api/models"
	"lasting-dynamics.com/habit_tracking_app/internal/api/repositories"
)

// UserService handles business logic related to users
type UserService struct {
	*BaseService
	userRepo *repositories.UserRepository
}

// NewUserService creates a new instance of UserService
func NewUserService() *UserService {
	return &UserService{userRepo: repositories.NewUserRepository()}
}

// EmailExists verifies whether a user record exists with the specified email address
// Used for authentication
func (us *UserService) EmailExists(email string) bool {
	return us.userRepo.EmailExists(email)
}

// Register registers a new user record
func (us *UserService) Register(name string, email string, password string) error {
	// Hash password
	user := &models.User{Name: name, Email: email}
	if err := user.SetPassword(password); err != nil {
		return err
	}

	// Create user in the database
	return us.userRepo.Create(user)
}

// FindByEmail is meant for finding the first user record by email for auth purposes
func (us *UserService) FindByEmail(email string) (*models.User, error) {
	return us.userRepo.FindByEmail(email)
}

// AuthenticateUser authenticates a user based on email and password
func (us *UserService) AuthenticateUser(email, password string) (*models.User, error) {
	// Retrieve user by email from the database
	user, err := us.userRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	// Check if the provided password matches the hashed password
	if err := user.CheckPassword(password); err != nil {
		return nil, err
	}

	return user, nil
}
