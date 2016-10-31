// See: https://elithrar.github.io/article/http-handler-error-handling-revisited/

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ErrorResponse represents an error sent to the user
type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// Error represents a handler error. It provides methods for a HTTP status
// code and embeds the built-in error interface.
type Error interface {
	error
	Status() int
}

// StatusError represents an error with an associated HTTP status code.
type StatusError struct {
	Code int
	Err  error
}

// Allows StatusError to satisfy the error interface.
func (se StatusError) Error() string {
	return se.Err.Error()

}

// Status Returns our HTTP status code.
func (se StatusError) Status() int {
	return se.Code

}

// The Handler struct that wraps our handlers
type Handler struct {
	H func(w http.ResponseWriter, r *http.Request) error
}

// ServeHTTP allows our Handler type to satisfy http.Handler.
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.H(w, r)

	if err != nil {
		var response ErrorResponse

		switch e := err.(type) {
		case Error:
			response = ErrorResponse{e.Status(), e.Error()}
		default:
			response = ErrorResponse{http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)}
		}

		w.WriteHeader(response.Status)
		json.NewEncoder(w).Encode(response)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) error {
	fmt.Fprint(w, "OK")
	return nil
}

func errorHandler(w http.ResponseWriter, r *http.Request) error {
	err := json.NewEncoder(w).Encode(fmt.Println)

	if err != nil {
		return StatusError{http.StatusInternalServerError, err}

	}

	return nil
}

func main() {
	http.Handle("/error", Handler{errorHandler})
	http.Handle("/", Handler{indexHandler})

	http.ListenAndServe(":8080", nil)
}
