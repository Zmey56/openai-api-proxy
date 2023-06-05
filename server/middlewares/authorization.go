package middlewares

import (
	"github.com/Zmey56/openai-api-proxy/authorization"
	"github.com/Zmey56/openai-api-proxy/log"
	"io"
	"net/http"
)

// AuthorizationMiddleware is a middleware that checks if the user is authorized to use the proxy
// and injects the user into in the header Openai-Api-Proxy-User
func AuthorizationMiddleware(next http.Handler, service authorization.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		username, pass, ok := r.BasicAuth()

		if !ok {
			log.Warning.Printf("request does not have basic auth")

			_, _ = io.Copy(io.Discard, r.Body)
			_ = r.Body.Close()

			http.Error(w, "Request does not have basic auth", http.StatusUnauthorized)
			return
		}

		err := service.Verify(username, pass)

		if err != nil {
			log.Warning.Printf("authorization failed for user %s. %s", username, err)

			_, _ = io.Copy(io.Discard, r.Body)
			_ = r.Body.Close()

			http.Error(w, "User was not found or password did not match", http.StatusUnauthorized)
			return
		}

		// set our own user to the header
		r.Header.Set("openai-api-proxy-user", username)

		next.ServeHTTP(w, r)

	})
}
