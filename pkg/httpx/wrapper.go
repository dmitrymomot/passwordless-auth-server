package httpx

import (
	"net/http"

	"github.com/dmitrymomot/go-pkg/response"
)

// Endpoint is a function to wrap http.HandleFunc to return error.
type Endpoint func(w http.ResponseWriter, r *http.Request) error

// Wrap wraps Endpoint to return http.HandlerFunc.
func HandlerFunc(e Endpoint) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := e(w, r); err != nil {
			if err := response.JSON(w, NewError(err)); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	}
}
