package api

import (
	"fmt"
	"net/http"
)

// ShortenURLHandler  Shortens the given long form of the URL
// service services.ShortenURLService
func ShortenURLHandler(s string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("The URL was hit")
	}
}
