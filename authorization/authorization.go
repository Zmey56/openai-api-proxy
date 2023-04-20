package authorization

import (
	"fmt"
	"github.com/Zmey56/openai-api-proxy/repository"
	"log"
	"net/http"
	"time"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Token struct {
	Value     string    `json:"value"`
	ExpiresAt time.Time `json:"expiresAt"`
	Username  string    `json:"username"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

func Authorization(tokens map[string]Token) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//  If the request is by the GET method, then we show the form
		if r.Method != "POST" {
			http.NotFound(w, r)
		}

		// If the request is made by the POST method, then we check the data
		login := r.FormValue("login")
		password := r.FormValue("password")

		//If  he has a different password, we return the error and return to the authentication page

		checkPassword, err := repository.CheckPassword(login, password)
		log.Println(checkPassword)
		if err != nil || !checkPassword {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		tokenValue := generateToken(login)
		expiresAt := time.Now().Add(time.Hour)
		token := Token{Value: tokenValue, ExpiresAt: expiresAt, Username: login}
		tokens[tokenValue] = token

		w.Header().Set("Authorization", "Bearer "+tokenValue)
		http.Redirect(w, r, "/openai", http.StatusFound)
	}
}

func generateToken(username string) string {
	tokenValue := fmt.Sprintf("%d_%s", time.Now().Add(time.Hour), username)
	return tokenValue
}
