package authorization

import (
	"github.com/Zmey56/openai-api-proxy/repository"
)

// StaticService is a static implementation of the authorization service
// Good for testing, only one token is allowed - "test"
type StaticService struct {
}

func (s StaticService) Verify(username, password string) bool {
	verifyToken, err := repository.VerifyTokenSQL(username, password)
	if err != nil {
		return false
	}

	if verifyToken {
		return true
	} else {
		return false
	}

}
