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

	db, err := sql.Open("sqlite3", getPathDB())
	if err != nil {
		panic(err)
	}
	defer db.Close()

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

}

func VerifyTokenSQL(usertoken []string) (bool, error) {
	db, err := sql.Open("sqlite3", getPathDB())
	if err != nil {
		return false, err
	}
	defer db.Close()

	query := `SELECT auth_token FROM users WHERE login = ?`
	rows, err := db.Query(query, usertoken[0])
	if err != nil {
		return false, err
	}
	defer rows.Close()

	if rows.Next() {
		var tokenDB string
		err = rows.Scan(&tokenDB)
		if err != nil {
			return false, err
		}

		if usertoken[1] == tokenDB {
			return true, nil
		}
	}

	return false, nil
}

func getPathDB() string {
	currentWorkingDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	pathDB := fmt.Sprintf("%s/openaiapiproxi.db", currentWorkingDirectory)
	return pathDB
}
