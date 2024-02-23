package seeder

import (
	"fmt"
	"os"

	"gorm.io/gorm"
	"lasting-dynamics.com/habit_tracking_app/internal/api/models"
	"lasting-dynamics.com/habit_tracking_app/internal/database/data"
	"lasting-dynamics.com/habit_tracking_app/internal/util/helpers"
)

// Seeder seeds data into the database.
type Seeder struct {
	Loader      data.Loader
	Transformer data.Transformer
}

// SeedData seeder function
func (s *Seeder) SeedData(db *gorm.DB, data data.JSONData) {
	users, habits := s.Transformer.Transform(data)

	seedData := SeedDataConfig{
		Users:  users,
		Habits: habits,
	}

	// Seed data
	seedData.Seed(db)
}

// SeedDataConfig holds the configuration for seed data.
type SeedDataConfig struct {
	Users  []models.User
	Habits []models.Habit
}

// SeedDataConfig creates a new instance of SeedDataConfig struct.
func NewSeedDataConfig() SeedDataConfig {
	return SeedDataConfig{}
}

// Seed seeds data into the database.
func (s *SeedDataConfig) Seed(db *gorm.DB) {
	for _, user := range s.Users {
		s.seedUser(db, user)
	}

	for _, habit := range s.Habits {
		s.seedHabit(db, habit)
	}
}

func (s *SeedDataConfig) seedUser(db *gorm.DB, user models.User) {
	user.SetPassword(os.Getenv("USER_DEFAULT_PASSWORD"))
	s.seedEntity(db, &user, "email = ?", user.Email)
}

func (s *SeedDataConfig) seedHabit(db *gorm.DB, habit models.Habit) {
	s.seedEntity(db, &habit, "user_id = ? AND name = ?", habit.UserID, habit.Name)
}

func (s *SeedDataConfig) seedEntity(db *gorm.DB, entity interface{}, condition string, args ...interface{}) {
	result := db.Where(condition, args...).First(&entity)
	if result.Error == nil {
		fmt.Printf("%T already exists, skipping seeding.\n", entity)
		return
	}

	result = db.Create(entity)
	if result.Error != nil {
		helpers.PanicIfError(fmt.Errorf("failed to create %T", entity))
	}

	fmt.Printf("%T seeded successfully\n", entity)
}
