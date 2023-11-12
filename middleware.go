package main

import "net/http"

const (
	USERNAME = "fanialfi"
	PASSWORD = "saichiopy"
)

func MiddlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			http.Error(w, "something went wrong", http.StatusUnauthorized)
			return
		}

		isValid := (username == USERNAME) && (password == PASSWORD)
		if !isValid {
			http.Error(w, "wrong username/password", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func MiddlewareAllowOnlyGet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Only Accept GET method", http.StatusMethodNotAllowed)
			return
		}
		next.ServeHTTP(w, r)
	})
}
