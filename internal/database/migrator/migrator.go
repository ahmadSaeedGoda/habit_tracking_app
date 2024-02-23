package migrator

import (
	"gorm.io/gorm"
	"lasting-dynamics.com/habit_tracking_app/internal/api/models"
)

// Migrate performs database migration using `GORM` auto migration technique for simplicity.
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&models.User{}, &models.Habit{})
}
