package repository

import (
	"database/sql"
	"os"
	"path/filepath"
	"testing"
)

type testDBImpl struct {
	db *sql.DB
}

type testDBImplClose struct {
	CloseFunc func() error
}

func (m *testDBImplClose) Close() error {
	if m.CloseFunc != nil {
		return m.CloseFunc()
	}
	return nil
}

func (m *testDBImpl) CreatedTableUsers() error {
	return nil
}

func (m *testDBImpl) VerifyUserPass(user, pass string) error {
	return nil
}

func TestSQLite(t *testing.T) {
	dbPath, err := os.MkdirTemp("", "sqlite*")
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		err := os.RemoveAll(dbPath)
		if err != nil {
			t.Fatal(err)
		}
	}()

	db, err := NewSQLite(filepath.Join(dbPath, "db.sqlite"))
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		err := db.Close()
		if err != nil {
			t.Fatal(err)
		}
	}()

	err = db.CreatedTableUsers()
	if err != nil {
		t.Fatal(err)
	}

	err = db.AddTestUsers()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("correct password", func(t *testing.T) {
		err = db.VerifyUserPass("login_1", "password1")
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("incorrect password", func(t *testing.T) {
		err = db.VerifyUserPass("login_2", "password_incorrect")
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})

	t.Run("incorrect username", func(t *testing.T) {
		err = db.VerifyUserPass("jdfhgjkadshgfkjsdhgkjdsf", "password_incorrect")
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})
}
