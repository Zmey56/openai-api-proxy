package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

// CreatedTableUsers create new table for users TO DO add colums with token and amount money
func CreatedTableUsers() {
	currentWorkingDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	pathDB := fmt.Sprintf("%s/openaiapiproxi.db", currentWorkingDirectory)

	db, err := sql.Open("sqlite3", pathDB)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	log.Println(db)

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
                        id INTEGER PRIMARY KEY,
                        login NOT NULL UNIQUE,
                        first_name TEXT NOT NULL,
                        last_name TEXT NOT NULL,
                        hashed_password TEXT NOT NULL,
                        email TEXT NOT NULL UNIQUE,
					 	access_level INT NOT NULL,
					 	amount_money REAL NOT NULL,
					 	tokens INT NOT NULL,
					 	auth_token TEXT,  
					 	created_at DATETIME NOT NULL,
					 	updated_at DATETIME NOT NULL
                    )`)

	if err != nil {
		panic(err)
	}

	log.Println("Table for users have created!!!")

}

func FindUser() {

}
