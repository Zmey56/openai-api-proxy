package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

// CreatedTableUsers create new table for users TO DO add colums with token and amount money

func CreatedTableUsers() {

	pathDB, err := getPathDB()
	if err != nil {
		fmt.Println(err)
		return
	}

	db, err := sql.Open("sqlite3", pathDB)
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

func VerifyTokenSQL(user, pass string) (bool, error) {
	pathDB, err := getPathDB()
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	db, err := sql.Open("sqlite3", pathDB)
	if err != nil {
		return false, err
	}
	defer db.Close()

	query := `SELECT hashed_password FROM users WHERE login = ?`
	rows, err := db.Query(query, user)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	if rows.Next() {
		var hashed_password string
		err = rows.Scan(&hashed_password)
		if err != nil {
			return false, err
		}

		if hashed_password == pass {
			return true, nil
		}
	}

	return false, nil
}

func getPathDB() (string, error) {
	currentWorkingDirectory, err := os.Getwd()
	if err != nil {
		return "", err
	}
	pathDB := fmt.Sprintf("%s/openaiapiproxi.db", currentWorkingDirectory)
	return pathDB, nil
}
