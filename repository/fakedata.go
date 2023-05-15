package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

func AddTestUsers() {
	//path to DB
	currentWorkingDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	pathDB := fmt.Sprintf("%s/openaiapiproxi.db", currentWorkingDirectory)
	log.Println(pathDB)

	//Opening a database connection
	db, err := sql.Open("sqlite3", pathDB)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// Generating random data for the "users" table
	for i := 1; i <= 10; i++ {

		login := fmt.Sprintf("login_%d", i)
		firstName := fmt.Sprintf("first_name%d", i)
		lastName := fmt.Sprintf("last_name%d", i)
		hashedPassword := fmt.Sprintf("password%d", i)
		email := fmt.Sprintf("user%d@example.com", i)
		accessLevel := i
		amountMoney := 100.10 / float64(i)
		tokens := i * 10
		authToken := fmt.Sprintf("test_%d", i)
		createdAt := time.Now()
		updatedAt := time.Now()

		// Inserting the random data into the "users" table
		_, err = db.Exec(`INSERT INTO users (first_name, last_name, hashed_password, email, access_level,
                   amount_money, tokens, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			login, firstName, lastName, hashedPassword, email, accessLevel,
			amountMoney, tokens, authToken, createdAt, updatedAt)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Data has been generated and inserted into the 'users' table")
}