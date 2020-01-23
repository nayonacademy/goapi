package middlewares

import "net/http"

// SetMiddlewareJSON - middleware
func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		next(w, r)
	}
}
