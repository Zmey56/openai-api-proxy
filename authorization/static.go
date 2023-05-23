package authorization

import (
	"errors"
	"github.com/Zmey56/openai-api-proxy/repository"
	"strings"
)

// StaticService is a static implementation of the authorization service
// Good for testing, only one token is allowed - "test"
type StaticService struct {
}

func (s StaticService) Verify(token string) (string, error) {
	// TO DO connect to DB and get auth token

	sepToken := strings.Split(token, "|")

	if len(sepToken) == 2 {
		verifyToken, err := repository.VerifyTokenSQL(sepToken)
		if err != nil {
			return "", err
		}
		if verifyToken {
			return strings.Join(sepToken, "|"), nil
		} else {
			return "", err
		}

	}

	return "", errors.New("unknown token")
}
