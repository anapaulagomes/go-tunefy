package main

import (
	"fmt"
	"net/http"
	"github.com/anapaulagomes/go-tunefy/api"
)

func main() {

	http.HandleFunc("/", api.Home)
	http.HandleFunc("/health-check", api.HealthCheck)

	fmt.Println("START Tunefy")

	http.ListenAndServe(":8080", nil)
}
