package authorization

import (
	"fmt"
	"github.com/Zmey56/openai-api-proxy/repository"
)

// StaticService is a static implementation of the authorization service
// Good for testing, only one token is allowed - "test"
type StaticService struct {
	DataBase *repository.DBImpl
}

func (s StaticService) Verify(username, password string) error {
	err := s.DataBase.VerifyToken(username, password)
	if err != nil {
		return err
	} else {
		fmt.Println(username)
	}
	return nil
}
