package repository

import (
	"database/sql"
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
