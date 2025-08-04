package middleware

import (
	"log"
	"net/http"
	"runtime/debug"
	"subscriptions/internal/shared/lib/res"
)

func Recoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("PANIC recovered: %v\n%s", err, debug.Stack())

				msg := "panic occurred while processing the request"
				res.SendError(w, http.StatusInternalServerError, msg, res.ServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
