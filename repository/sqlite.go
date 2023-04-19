package repository

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

// connecting to DB
func ConnectToDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", GetCurrentDirectory()+"/openaiapiproxi.db")
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

// current directory to DB
func GetCurrentDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

// CheckPassword check the password with the database password
func CheckPassword(login, password string) (bool, error) {
	db, err := ConnectToDatabase()
	if err != nil {
		return false, err
	}
	log.Println("OK with DB", login)
	row := db.QueryRow("SELECT hashed_password FROM users WHERE login=?", login)
	var passwordDB string
	err = row.Scan(&passwordDB)
	if err != nil {
		return false, err
	}
	log.Println(password, passwordDB)

	if password == passwordDB {
		return true, nil
	} else {
		return false, nil
	}
}
