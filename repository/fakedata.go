package repository

import (
	"fmt"
	"github.com/Zmey56/openai-api-proxy/log"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func (db *DBImpl) AddTestUsers() error {
	// Generating random data for the "users" table
	for i := 1; i <= 10; i++ {

		login := fmt.Sprintf("login_%d", i)
		firstName := fmt.Sprintf("first_name%d", i)
		lastName := fmt.Sprintf("last_name%d", i)
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(fmt.Sprintf("password%d", i)), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println("Problem with hashedPassword")
			return err
		}
		email := fmt.Sprintf("user%d@example.com", i)
		accessLevel := i
		amountMoney := 100.10 / float64(i)
		tokens := i * 10
		authToken := fmt.Sprintf("test_%d", i)
		createdAt := time.Now()
		updatedAt := time.Now()

		// Inserting the random data into the "users" table

		_, err = db.db.Exec(`INSERT INTO users (login, first_name, last_name, hashed_password, email, access_level,
                   amount_money, tokens, auth_token, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			login, firstName, lastName, hashedPassword, email, accessLevel,
			amountMoney, tokens, authToken, createdAt, updatedAt)
		if err != nil {
			log.Error.Printf("failed to insert data into the 'users' table: %s", err)
		}

		if log.IsDebug() {
			log.Debug.Printf("created user %s, token %s", login, authToken)
		}
	}

	return nil
}
