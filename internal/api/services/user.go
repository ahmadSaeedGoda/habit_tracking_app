package services

import (
	dto "lasting-dynamics.com/habit_tracking_app/internal/api/entities"
)

// GetUserByID retrieves a user by its ID
func (us *UserService) GetUserByID(userID uint) (*dto.UserDTO, error) {
	user, err := us.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	// Map user to DTO
	userDTO := &dto.UserDTO{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	return userDTO, nil
}
