package middleware

import (
	"log"
	"net/http"
)

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s - %s  ", r.Method, r.RequestURI)
		next.ServeHTTP(w, r) // call the next handler, which can be another middleware or the final handler
	})
}


func EnableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "http://localhost:3000" || origin == "https://jessandbrent.ca" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}
		next.ServeHTTP(w, r)
	})
}

