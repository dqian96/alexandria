package server

import (
	"log"
	"net/http"

	d "github.com/dqian96/alexandria/director"
	"github.com/gorilla/mux"
)

// Server is the http server for handling client (external) requests
type Server interface {
	Serve(string) error
}

type server struct {
	router *mux.Router
}

// Takes a port and starts the server on said port; returns an error
func (s *server) Serve(port string) error {
	log.Printf("Attempting to start HTTP on port %s", port)

	if err := http.ListenAndServe(port, s.router); err != nil {
		log.Printf("Failed to start HTTP server: %v", err)
		return err
	}
	return nil
}

// NewServer creates a new Server based on a given Archive
func NewServer(d d.Director) Server {
	r := NewRouter(d)
	s := &server{
		router: r,
	}
	return s
}
