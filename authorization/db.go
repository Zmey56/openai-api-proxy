package authorization

import (
	"github.com/Zmey56/openai-api-proxy/repository"
)

type DatabaseService struct {
	database *repository.DBImpl
}

func NewDatabaseService(db *repository.DBImpl) *DatabaseService {
	return &DatabaseService{
		database: db,
	}
}

func (s DatabaseService) Verify(username, password string) error {
	err := s.database.VerifyUserPass(username, password)

	if err != nil {
		return err
	}
	return nil
}
