package middlewares

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockAuthorizationService struct {
	verifyFunc       func(username, password string) error
	receivedUsername string
	receivedPassword string
}

func (m *MockAuthorizationService) Verify(username, password string) error {
	m.receivedUsername = username
	m.receivedPassword = password
	return m.verifyFunc(username, password)
}

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

func TestAuthorizationMiddleware_CorrectCredentials(t *testing.T) {

	mockService := &MockAuthorizationService{
		verifyFunc: func(username, password string) error {
			if username != "correctUsername" {
				t.Errorf("Expected username 'correctUsername' got %s", username)
			}
			if password != "correctPassword" {
				t.Errorf("Expected username 'correctPassword' got %s", username)
			}
			return nil
		},
	}

	handler := AuthorizationMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}), mockService)

	req := httptest.NewRequest("GET", "/", nil)
	req.SetBasicAuth("correctUsername", "correctPassword")
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rec.Code)
	}

}

func TestAuthorizationMiddleware_IncorrectUsername(t *testing.T) {
	mockService := &MockAuthorizationService{
		verifyFunc: func(username, password string) error {
			if username != "incorrectUsername" {
				t.Errorf("Expected username 'incorrectUsername' got %s", username)
			}
			if password != "somePassword" {
				t.Errorf("Expected username 'somePassword' got %s", username)
			}
			return ErrUserNotFound
		},
	}

	handler := AuthorizationMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}), mockService)

	req := httptest.NewRequest("GET", "/", nil)
	req.SetBasicAuth("incorrectUsername", "somePassword")
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Errorf("Expected status code %d, got %d", http.StatusUnauthorized, rec.Code)
	}
}

func TestAuthorizationMiddleware_IncorrectCredentials(t *testing.T) {
	mockService := &MockAuthorizationService{
		verifyFunc: func(username, password string) error {
			if username != "incorrectUsername" {
				t.Errorf("Expected username 'incorrectUsername' got %s", username)
			}
			if password != "incorrectPassword" {
				t.Errorf("Expected username 'incorrectPassword' got %s", username)
			}
			return ErrInvalidCredentials
		},
	}

	handler := AuthorizationMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}), mockService)

	req := httptest.NewRequest("GET", "/", nil)
	req.SetBasicAuth("incorrectUsername", "incorrectPassword")
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Errorf("Expected status code %d, got %d", http.StatusUnauthorized, rec.Code)
	}
}

func TestAuthorizationMiddleware_NoBasicAuth(t *testing.T) {
	mockService := &MockAuthorizationService{}

	handler := AuthorizationMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}), mockService)

	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Errorf("Expected status code %d, got %d", http.StatusUnauthorized, rec.Code)
	}
}
