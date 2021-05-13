package middlewares

import (
	"fmt"
	"net/http"
)

// LogRequest logs all requests
func LogRequest (next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("\n %s %s %s ", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}