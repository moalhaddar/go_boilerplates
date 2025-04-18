package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// struct embedding
// https://gobyexample.com/struct-embedding
type Server struct {
	http.Server
}

func NewServer(addr string, mux http.Handler) *Server {
	return &Server{
		Server: http.Server{
			Addr:    addr,
			Handler: mux,
		},
	}
}

func (server *Server) StartServer() {
	gracefulShutdownChan := make(chan bool, 1)
	go func() {
		log.Printf("Listening on %s\n", server.Addr)
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server error: %v\n", err)
		}

		log.Println("Gracefully shutting down the server.")
		gracefulShutdownChan <- true
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("HTTP close error: %v\n", err)
	}
	<-gracefulShutdownChan
	log.Println("Graceful shutdown complete.")
}
