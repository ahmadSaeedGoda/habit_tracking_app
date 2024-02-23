package repositories

import (
	"gorm.io/gorm"
	"lasting-dynamics.com/habit_tracking_app/internal/api/models"
	DBManager "lasting-dynamics.com/habit_tracking_app/internal/database/dbmanager"
)

// HabitRepositoryInterface defines the interface for interacting with habit data
type HabitRepositoryInterface interface {
	ListByUser(userID uint) ([]models.Habit, error)
	GetHabitByID(habitID uint) (*models.Habit, error)
	CreateHabit(habit *models.Habit) error
	IsHabitNameUnique(userID uint, habitName string) bool
	IsUserHabit(userID, habitID uint) bool
	UpdateHabit(habit *models.Habit) error
	DeleteHabit(habitID uint) error
}

// HabitRepository handles database operations related to `Habits`
type HabitRepository struct {
	*BaseRepository
	db *gorm.DB
}

// NewHabitRepository creates a new instance of `HabitRepository`
func NewHabitRepository() *HabitRepository {
	return &HabitRepository{db: DBManager.DB}
}

// ListByUser retrieves a list of habits for a specific user by user ID
func (hr *HabitRepository) ListByUser(userID uint) ([]models.Habit, error) {
	var habits []models.Habit
	if err := hr.db.Where("user_id = ?", userID).Find(&habits).Error; err != nil {
		return nil, err
	}
	return habits, nil
}

// GetHabitByID retrieves a habit by its ID
func (hr *HabitRepository) GetHabitByID(habitID uint) (*models.Habit, error) {
	var habit models.Habit
	if err := hr.db.First(&habit, habitID).Error; err != nil {
		return nil, err
	}
	return &habit, nil
}

// CreateHabit creates a new habit
func (hr *HabitRepository) CreateHabit(habit *models.Habit) error {
	if err := hr.db.Create(habit).Error; err != nil {
		return err
	}
	return nil
}

// IsHabitNameUnique checks if the habit name is unique for the user related habits
func (hr *HabitRepository) IsHabitNameUnique(userID uint, habitName string) bool {
	var count int64
	hr.db.Model(&models.Habit{}).Where("user_id = ? AND name = ?", userID, habitName).Count(&count)
	return count == 0
}

// IsUserHabit checks if the habit with the specified ID belongs to the specified user
// Verifies ownership of a habit before update, delete operations etc.
func (hr *HabitRepository) IsUserHabit(userID, habitID uint) bool {
	var count int64
	hr.db.Model(&models.Habit{}).Where("user_id = ? AND id = ?", userID, habitID).Count(&count)
	return count > 0
}

// UpdateHabit updates an existing habit
func (hr *HabitRepository) UpdateHabit(habit *models.Habit) error {
	if err := hr.db.Save(habit).Error; err != nil {
		return err
	}
	return nil
}

// DeleteHabit deletes a habit by its ID
func (hr *HabitRepository) DeleteHabit(habitID uint) error {
	if err := hr.db.Delete(&models.Habit{}, habitID).Error; err != nil {
		return err
	}
	return nil
}
