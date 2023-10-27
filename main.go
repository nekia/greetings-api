package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/rs/cors"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("got / request from %s\n", r.RemoteAddr)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(greeting()))
	})

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://greetings.kylepenfound.com", "http://localhost:8081"},
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})
	handler := c.Handler(mux)
	err := http.ListenAndServe(":8080", handler)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func greeting() string {
	greeting := "Hello Demo!"
	return fmt.Sprintf("{\"greeting\":\"%s\"}", greeting)
}
