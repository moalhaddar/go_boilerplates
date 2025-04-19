package main

import (
	"log"
	"net/http"

	"alhaddar.dev/http_middlewares/middleware"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})

	chain := middleware.New([]middleware.Middleware{
		middleware.NewLoggerMiddleware(false),
	}, mux).Build()

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: chain,
	}

	log.Printf("Listening on %s\n", server.Addr)
	server.ListenAndServe()
}
