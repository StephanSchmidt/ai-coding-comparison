package main

import (
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDb() *gorm.DB {
	dbName := "todo.db"

	// Check if the database file exists
	_, err := os.Stat(dbName)
	dbExists := !os.IsNotExist(err)

	if !dbExists {
		fmt.Println("Database does not exist. Creating database and tables.")

		// Create the database file
		file, err := os.Create(dbName)
		if err != nil {
			panic(err)
		}
		file.Close()
	} else {
		fmt.Println("Database exists.")
	}

	// Open the database using GORM
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// If the database did not exist, execute SQL statements to create tables
	if !dbExists {
		createTableSQL := `
		CREATE TABLE todos (
			id SERIAL PRIMARY KEY,
			title VARCHAR(255) NOT NULL
		);
		`
		db.Exec(createTableSQL)
	}

	// Continue with your application logic...
	return db
}
