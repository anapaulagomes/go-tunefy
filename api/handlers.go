package api

import (
	"io"
	"net/http"
)

const (
	contentType string = "application/json"
)

//Home returns a beautiful Hello World
func Home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", contentType)

	io.WriteString(w, `{"message": "Hello World"}`)
}

//HealthCheck returns 200 to serve the smoke test
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
