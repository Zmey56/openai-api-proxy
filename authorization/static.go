package authorization

import (
	"errors"
	"strings"
)

var (
	InvalidUsernameOrPassword = errors.New("user does not exist or passwords do not match")
)

// StaticService is a static implementation of the authorization service
// Good for testing, only one token is allowed - "test"
type StaticService struct {
	username string
	password string
}

func NewStaticService(u, p string) *StaticService {
	return &StaticService{username: u, password: p}
}

func (s StaticService) Verify(u, p string) error {

	if strings.ToLower(s.username) == strings.ToLower(u) && strings.ToLower(s.password) == strings.ToLower(p) {
		return nil
	}

	return InvalidUsernameOrPassword
}
