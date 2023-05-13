package middlewares

import (
	"net/http"
)

// RemovePathPrefixMiddleware Removes the prefix from the request path
func RemovePathPrefixMiddleware(next http.Handler, prefix string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = r.URL.Path[len(prefix):]
		next.ServeHTTP(w, r)
	})
}
