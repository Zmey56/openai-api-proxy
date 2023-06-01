package authorization

import (
	"github.com/Zmey56/openai-api-proxy/repository"
)

type DBAuth struct {
	Database *repository.DBImpl
}

func (s DBAuth) VerifyDB(username, password string) error {
	err := s.Database.VerifyUserPass(username, password)
	if err != nil {
		return err
	}
	return nil
}
