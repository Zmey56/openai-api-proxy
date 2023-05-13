package middlewares

import (
	"github.com/Zmey56/openai-api-proxy/authorization"
	"io"
	"log"
	"net/http"
)

// AuthorizationMiddleware is a middleware that checks if the user is authorized to use the proxy
// and injects the user into in the header Openai-Api-Proxy-User
func AuthorizationMiddleware(next http.Handler, service authorization.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// trying to identify the user
		user, err := service.Verify(r.Header.Get("Authorization"))
		log.Println("test user", user)
		if err != nil {
			_, _ = io.Copy(io.Discard, r.Body)
			_ = r.Body.Close()

			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// set our own user to the header
		r.Header.Set("openai-api-proxy-user", user)

		next.ServeHTTP(w, r)
	})
}
