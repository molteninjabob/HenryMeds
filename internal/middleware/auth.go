package middleware

import (
	"net/http"
	"strings"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
        if tokenString == "" {
            w.WriteHeader(http.StatusUnauthorized)
            return
        }

        tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		if tokenString == "" {
            w.WriteHeader(http.StatusUnauthorized)
            return
        }
	
		next.ServeHTTP(w, r)
	})
}