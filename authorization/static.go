package authorization

import (
	"fmt"
)

// StaticService is a static implementation of the authorization service
// Good for testing, only one token is allowed - "test"
type StaticService struct {
}

func (s StaticService) Verify(username, password string) bool {
	fmt.Println(username)
	return true
}
