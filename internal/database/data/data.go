package data

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"lasting-dynamics.com/habit_tracking_app/internal/api/models"
)

const (
	dataDir      = "internal/database/data"
	dataFileName = "seed_data.json"
)

type JSONData struct {
	Habits []HabitJSON `json:"habits"`
	Users  []UserJSON  `json:"users"`
	// Append more fields as we evolve...
}

type HabitJSON struct {
	Name        string `json:"Name"`
	Description string `json:"Description"`
	UserID      uint   `json:"UserID"`
}

type UserJSON struct {
	Name  string `json:"Name"`
	Email string `json:"Email"`
}

// NewTransformer creates a new instance of transformer.
// A Transformer is responsible for transforming the data from the json format to be returned as slices
// Meant to be used by the seeder, for initial data seeding purposes.
func NewTransformer() Transformer {
	return &transformer{}
}

// Transform is responsible for transforming the data from the `json` format to be returned as slices.
// The slices are more convenient for now for the seeder to seed them into DB.
func (t *transformer) Transform(jsonData JSONData) ([]models.User, []models.Habit) {
	var users []models.User
	var habits []models.Habit

	for _, userJSON := range jsonData.Users {
		users = append(users, models.User{
			Name:  userJSON.Name,
			Email: userJSON.Email,
		})
	}

	for _, habitJSON := range jsonData.Habits {
		habits = append(habits, models.Habit{
			Name:        habitJSON.Name,
			Description: habitJSON.Description,
			UserID:      habitJSON.UserID,
		})
	}

	return users, habits
}

// Loader provides a method to load data.
// Meant to be used by seeder object along with the transformer.
// This is just a contract for the behavior to be replaced by implementations.
type Loader interface {
	LoadData() (JSONData, error)
}

// fileLoader implements Loader.
type fileLoader struct{}

// Transformer provides a method to Transform data from `JSON` to `Models`.
type Transformer interface {
	Transform(jsonData JSONData) ([]models.User, []models.Habit)
}

// transformer implements `Transformer` interface.
type transformer struct{}

// getDataFilePath is a util func for crafting the path to the file containing the data to seed in DB.
// Extracted for Single responsibility.
func (fl *fileLoader) getDataFilePath() (string, error) {
	currentPath, err := os.Getwd()
	if err != nil {
		// Log the error and exit the program, as the working directory is critical.
		log.Fatalf("Error getting current working directory: %v", err)
	}

	dataFilePath := filepath.Join(currentPath, dataDir, dataFileName)
	return dataFilePath, nil
}

// readSeedData is a util func for fetching the data from the json file containing it to be seed in DB.
// Extracted for Single responsibility.
func (fl *fileLoader) readSeedData(filePath string) (JSONData, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return JSONData{}, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var data JSONData
	err = decoder.Decode(&data)
	if err != nil {
		return JSONData{}, err
	}

	return data, nil
}

// NewFileLoader creates a new instance of fileLoader.
func NewFileLoader() Loader {
	return &fileLoader{}
}

// LoadData loads the seed data from the file.
func (fl *fileLoader) LoadData() (JSONData, error) {
	dataFilePath, err := fl.getDataFilePath()
	if err != nil {
		return JSONData{}, err
	}

	// Read data from file
	data, err := fl.readSeedData(dataFilePath)
	if err != nil {
		return JSONData{}, err
	}

	return data, nil
}
