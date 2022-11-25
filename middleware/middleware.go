package middleware

import (
	"log"
	"net/http"
	"time"
)

// The Middleware handler functions.
// They receive a handler function and return another, to provide
// some pre-processing to the request before actually handling it.
type Middleware func(http.HandlerFunc) http.HandlerFunc

// Logs the request as it starts being processed. This does not log
// any responses, results or statuses. That should be handled by the
// response logger
func RequestLogger() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			func() {
				log.Printf("\033[44m %s \033[0m | PATH: \033[33m\"%s\"\033[0m | DURATION: \033[42m %v \033[0m",
					r.Method, r.URL.Path, time.Since(start),
				)
			}()
			f(w, r)
		}
	}
}
