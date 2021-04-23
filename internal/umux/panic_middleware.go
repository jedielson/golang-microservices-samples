package umux

import (
	"io"
	"net/http"
)

func PanicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				io.WriteString(w, "Something went wrong")
			}
		}()
		next.ServeHTTP(w, r)
	})
}
