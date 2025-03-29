package server

import (
	"log"
	"net/http"
)

// Server structure for the HTTP server
type Server struct {
	Logger  *log.Logger
	server  *http.Server
	Mux     *http.ServeMux // Expose Mux so handlers can be registered externally
}

// New creates a new Server instance
func New(logger *log.Logger) *Server {
	mux := http.NewServeMux()
	return &Server{
		Logger: logger,
		Mux:    mux,
		server: &http.Server{
			Handler: mux,
		},
	}
}

// Start starts the HTTP server on the specified address
func (s *Server) Start(addr string) error {
	s.server.Addr = addr
	s.Logger.Printf("Starting server on %s", addr)
	return s.server.ListenAndServe()
}

// Stop gracefully shuts down the server
func (s *Server) Stop() error {
	s.Logger.Println("Shutting down server")
	return s.server.Close()
}
