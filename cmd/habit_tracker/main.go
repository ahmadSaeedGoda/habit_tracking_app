package main

import (
	"fmt"

	"lasting-dynamics.com/habit_tracking_app/internal/api/routes"
	DBManager "lasting-dynamics.com/habit_tracking_app/internal/database/dbmanager"
	"lasting-dynamics.com/habit_tracking_app/internal/util/helpers"
)

// ANSI color escape codes
const (
	Reset = "\033[0m"
	// Red    = "\033[31m"
	Green = "\033[32m"
	// Yellow = "\033[33m"
)

func main() {
	dbm, err := DBManager.Init()
	helpers.PanicIfError(err)

	defer dbm.CloseConnection()

	fmt.Println(Green + "Database connection, migration, and seeding completed successfully." + Reset)

	routes.SetupRouter().Run()
}
