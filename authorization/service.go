package authorization

type Service interface {
	// Verify verifies the token and returns the user
	Verify(username, pass string) bool
}
