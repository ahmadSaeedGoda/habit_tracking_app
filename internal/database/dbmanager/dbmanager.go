package DBManager

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"lasting-dynamics.com/habit_tracking_app/internal/database/data"
	"lasting-dynamics.com/habit_tracking_app/internal/database/migrator"
	"lasting-dynamics.com/habit_tracking_app/internal/database/seeder"
	"lasting-dynamics.com/habit_tracking_app/internal/util/helpers"
)

// DB is a global variable for other modules/packages
var DB *gorm.DB

// DBManager handles database operations.
type DBManager struct {
	db *gorm.DB
}

func createDSN() string {
	// Fetch environment variables
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	port := os.Getenv("POSTGRES_PORT")
	sslMode := os.Getenv("SSL_MODE")

	// Construct the DSN string
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, password, dbName, port, sslMode)
}

// NewDBManager creates a new instance of DBManager.
func NewDBManager() (*DBManager, error) {
	// Construct the DSN string
	dsn := createDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &DBManager{db: db}, nil
}

// GetConnection returns a reference to the current DB conn.
// TODO:: Can introduce connection pooling in the future.
func (m *DBManager) GetConnection() *gorm.DB {
	return m.db
}

// Migrate performs database migration.
func (dbm *DBManager) Migrate() error {
	return migrator.Migrate(dbm.db)
}

// Close closes the database connection.
func (dbm *DBManager) CloseConnection() error {
	sqlDB, err := dbm.db.DB()
	helpers.PanicIfError(err)
	return sqlDB.Close()
}

// Init is the entrypoint for the DB conn.
// It initializes a dbManager obj, sets it in a global variable for other modules/packages to access,
// then migrate the DB up if any, and finally seeds the data into it.
func Init() (*DBManager, error) {
	dbManager, err := NewDBManager()
	if err != nil {
		return nil, err
	}

	DB = dbManager.db

	// Migrate the database
	if err := dbManager.Migrate(); err != nil {
		return nil, err
	}

	// Seed the database
	loader := data.NewFileLoader()
	transformer := data.NewTransformer()

	seeder := &seeder.Seeder{Loader: loader, Transformer: transformer}

	data, err := seeder.Loader.LoadData()
	if err != nil {
		log.Fatal(err)
	}

	seeder.SeedData(DB, data)

	return dbManager, nil
}
