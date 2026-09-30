package middleware

import (
	//"hotelapi/handlers"
	"net/http"
)

func HotelMiddleware(key string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			requestkey := r.Header.Get("API-KEY")
			if requestkey != key {
				http.Error(w, "Wrong Key", http.StatusForbidden)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
