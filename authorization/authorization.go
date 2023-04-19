package authorization

import (
	"bytes"
	"github.com/Zmey56/openai-api-proxy/repository"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"os"
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

func Authorization(w http.ResponseWriter, r *http.Request) {
	jwtSecretKey := []byte(os.Getenv("jwtSecretKey"))

	if bytes.Equal(jwtSecretKey, []byte{}) {
		jwtSecretKey = []byte("jwtSecretKey")
	}

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

	//Generating useful data that will be stored in the token
	payload := jwt.MapClaims{
		"sub": login,
		"exp": time.Now().Add(time.Minute * 72).Unix(),
	}

	// Create a new JWT token and sign it using the HS256 algorithm
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	t, err := token.SignedString(jwtSecretKey)
	if err != nil {
		log.Println("JWT token signing")
		w.WriteHeader(http.StatusInternalServerError)
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	w.Header().Set("Authorization", "Bearer "+t)
	http.Redirect(w, r, "/openai", http.StatusFound)
}
