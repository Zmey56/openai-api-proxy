package authorization

type Service interface {
	// Verify verifies the token and returns the user
	Verify(token string) (string, error)
}
