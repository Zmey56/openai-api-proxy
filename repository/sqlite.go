package repository

import (
	"database/sql"
	"errors"
	"github.com/Zmey56/openai-api-proxy/log"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type DBImpl struct {
	db *sql.DB
}

func NewSQLite(path string) (*DBImpl, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &DBImpl{db: db}, nil
}

func (db *DBImpl) Close() error {
	return db.db.Close()
}

func (db *DBImpl) CreatedTableUsers() error {
	_, err := db.db.Exec(`CREATE TABLE IF NOT EXISTS users (
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

	return err
}

func (db *DBImpl) VerifyUserPass(user, pass string) error {
	query := `SELECT hashed_password FROM users WHERE login = ?`
	rows, err := db.db.Query(query, strings.ToLower(user))
	if err != nil {
		return err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Debug.Printf("failed to close rows: %s", err)
		}
	}()

	if rows.Next() {
		var hashedPassword []byte
		err = rows.Scan(&hashedPassword)
		if err != nil {
			log.Error.Print("Can't Scan password")
			return err
		}

		err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(pass))
		if err != nil {
			log.Error.Print("Password not found")
			return err
		} else {
			return nil
		}
	}

	return errors.New("user not found")
}
