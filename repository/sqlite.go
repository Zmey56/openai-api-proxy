package repository

import (
	"database/sql"
	"github.com/Zmey56/openai-api-proxy/log"
	_ "github.com/mattn/go-sqlite3"
)

// CreatedTableUsers create new table for users TO DO add colums with token and amount money
func CreatedTableUsers(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
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

	log.Debug.Println("table users has been created")

}

func FindUser() {

}
