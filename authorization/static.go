package authorization

import "errors"

// StaticService is a static implementation of the authorization service
// Good for testing, only one token is allowed - "test"
type StaticService struct {
}

func (s StaticService) Verify(token string) (string, error) {
	if token == "test" {
		return "test", nil
	}

	return "", errors.New("unknown token")
}
