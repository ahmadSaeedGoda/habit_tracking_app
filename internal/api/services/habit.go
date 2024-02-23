package services

import (
	"errors"
	"strings"

	dto "lasting-dynamics.com/habit_tracking_app/internal/api/entities"
	"lasting-dynamics.com/habit_tracking_app/internal/api/models"
	"lasting-dynamics.com/habit_tracking_app/internal/api/repositories"
)

// HabitServiceInterface defines the interface for interacting with habits
type HabitServiceInterface interface {
	ListByUser(userID uint) ([]*dto.HabitDTO, error)
	GetHabitByID(habitID uint) (*dto.HabitDTO, error)
	CreateHabit(userId uint, habit *dto.HabitDTO) error
	IsHabitNameUnique(userID uint, habitName string) bool
	IsUserHabit(userID uint, habitID uint) bool
	UpdateHabit(userId uint, habitID uint, habit *dto.HabitDTO) error
	DeleteHabit(userId uint, habitID uint) error
}

// HabitService provides methods to interact with habits
// Layer for business logic to separate data access from request handling
type HabitService struct {
	userRepo  repositories.UserRepositoryInterface
	habitRepo repositories.HabitRepositoryInterface
}

// NewHabitService creates a new instance of HabitService
func NewHabitService() *HabitService {
	return &HabitService{
		userRepo:  repositories.NewUserRepository(),
		habitRepo: repositories.NewHabitRepository(),
	}
}

// ListByUser retrieves a list of habits for a specific user by user ID
func (hs *HabitService) ListByUser(userID uint) ([]*dto.HabitDTO, error) {
	habits, err := hs.habitRepo.ListByUser(userID)
	if err != nil {
		return nil, err
	}

	// Map habits to DTOs
	habitDTOs := make([]*dto.HabitDTO, len(habits))
	for i, h := range habits {
		habitDTOs[i] = &dto.HabitDTO{
			ID:          h.ID,
			Name:        h.Name,
			Description: h.Description,
		}
	}

	return habitDTOs, nil
}

// GetHabitByID retrieves a habit by its ID
func (hs *HabitService) GetHabitByID(habitID uint) (*dto.HabitDTO, error) {
	habit, err := hs.habitRepo.GetHabitByID(habitID)
	if err != nil {
		return nil, err
	}

	habitDTO := &dto.HabitDTO{
		ID:          habit.ID,
		Name:        habit.Name,
		Description: habit.Description,
		UserId:      habit.UserID,
	}

	return habitDTO, nil
}

// CreateHabit creates a new habit for a specific user
func (hs *HabitService) CreateHabit(userId uint, habit *dto.HabitDTO) error {
	// Check if the user exists
	_, err := hs.userRepo.GetUserByID(habit.UserId)
	if err != nil {
		return err
	}

	// Map DTO to model
	newHabit := &models.Habit{
		Name:        habit.Name,
		Description: habit.Description,
		UserID:      habit.UserId,
	}

	// Create habit
	if err := hs.habitRepo.CreateHabit(newHabit); err != nil {
		return err
	}

	return nil
}

// IsUserHabit checks if the habit belongs to the user
func (hs *HabitService) IsUserHabit(userID uint, habitID uint) bool {
	return hs.habitRepo.IsUserHabit(userID, habitID)
}

// IsHabitNameUnique checks if the habit name is unique for the user who own it
func (hs *HabitService) IsHabitNameUnique(userID uint, habitName string) bool {
	return hs.habitRepo.IsHabitNameUnique(userID, habitName)
}

// UpdateHabit updates an existing habit, ensuring it belongs to the specified user
func (hs *HabitService) UpdateHabit(userID, habitID uint, habitDTO *dto.HabitDTO) error {
	// Retrieve the habit by ID
	habit, err := hs.habitRepo.GetHabitByID(habitID)
	if err != nil {
		return err
	}

	// Validate that the habit belongs to the specified user
	if habit.UserID != userID {
		return errors.New("Unauthorized")
	}
	// TODO:: perform validation in separate place for each entity across all its operations
	// Update the habit fields
	habit.Name = func() string {
		if strings.TrimSpace(habitDTO.Name) == "" {
			return habit.Name
		}
		return habitDTO.Name
	}()
	habit.Description = func() string {
		if strings.TrimSpace(habitDTO.Description) == "" {
			return habit.Description
		}
		return habitDTO.Description
	}()

	// Call the habit repository to update the habit
	return hs.habitRepo.UpdateHabit(habit)
}

// DeleteHabit deletes a habit by its ID, ensuring it belongs to the specified user
func (hs *HabitService) DeleteHabit(userID uint, habitID uint) error {
	// Retrieve the habit by ID
	habit, err := hs.habitRepo.GetHabitByID(habitID)
	if err != nil {
		return err // Handle the error accordingly (not found, database error, etc.)
	}

	// Validate that the habit belongs to the specified user
	if habit.UserID != userID {
		return errors.New("habit does not belong to the specified user")
	}

	// Call the habit repository to delete the habit
	return hs.habitRepo.DeleteHabit(habitID)
}
