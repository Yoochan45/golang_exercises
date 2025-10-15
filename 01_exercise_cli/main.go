package main

import (
	"01_exercise/config"
	"fmt"
)

func main() {
	database := config.ConnectDB()
	defer database.Close()

	if database != nil {
		fmt.Println("Database connection established succesfully!")
	} else {
		fmt.Println("Error: database connection failed (database is nil). Exiting.")
		return
	}
}
