package main

import (
	"fmt"
	"net/http"
	"tunefy/api"
)

func main() {

	http.HandleFunc("/", api.Home)
	http.HandleFunc("/health-check", api.HealthCheck)

	fmt.Println("START Tunefy")

	http.ListenAndServe(":8181", nil)
}
