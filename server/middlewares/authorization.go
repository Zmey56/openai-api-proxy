package middlewares

import (
	"github.com/Zmey56/openai-api-proxy/authorization"
	"io"
	"net/http"
)

// AuthorizationMiddleware is a middleware that checks if the user is authorized to use the proxy
// and injects the user into in the header Openai-Api-Proxy-User
func AuthorizationMiddleware(next http.Handler, service authorization.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// trying to identify the user

		username, pass, ok := r.BasicAuth()

		if !ok {
			_, _ = io.Copy(io.Discard, r.Body)
			_ = r.Body.Close()

			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}

		err := service.Verify(username, pass)

		if err != nil {
			_, _ = io.Copy(io.Discard, r.Body)
			_ = r.Body.Close()

			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}

		// set our own user to the header
		r.Header.Set("openai-api-proxy-user", username)

		next.ServeHTTP(w, r)

	})
}
