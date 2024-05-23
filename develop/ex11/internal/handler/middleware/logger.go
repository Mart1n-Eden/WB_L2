package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logging(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO:
		log.Printf("[%s] %s %s", time.Now().Format("2000-01-01 00:00:00"), r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		//log.Printf("[%s] %s %s", time.Now().Format("2000-01-01 00:00:00"), r.Method, r.URL.Path)
	})
}
