package http

import (
	"fmt"
	"net/http"
)

func MiddlewareUser(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// do stuff before the handlers
		h.ServeHTTP(w, r)
		fmt.Print("Middleware Validation \n")
		fmt.Print(r.Header)
		fmt.Print("Middleware Validation \n")
		// do stuff after the hadlers

	})
}
