package server

import (
	"log"
	"net/http"
	"time"
)

// HandlerRegistrar interface for registering handlers
type HandlerRegistrar interface {
	RegisterHandlers(srv *Server)
}

// Server структура HTTP-сервера
type Server struct {
	logger *log.Logger
	httpServer *http.Server
}

// New создает новый экземпляр сервера
func New(logger *log.Logger) *Server {
	// Создаем роутер
	mux := http.NewServeMux()
	
	// Настраиваем HTTP-сервер
	httpServer := &http.Server{
		Addr:         ":8080",              
		Handler:      mux,                  
		ErrorLog:     logger,               
		ReadTimeout:  5 * time.Second,      
		WriteTimeout: 10 * time.Second,     
		IdleTimeout:  15 * time.Second,     
	}

	return &Server{
		logger:     logger,
		httpServer: httpServer,
	}
}

// RegisterHandler регистрирует обработчик для пути
func (s *Server) RegisterHandler(path string, handler http.HandlerFunc) {
	mux := s.httpServer.Handler.(*http.ServeMux)
	mux.HandleFunc(path, handler)
}

// Start запускает сервер
func (s *Server) Start() error {
	s.logger.Printf("Starting server on %s", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

// Stop останавливает сервер
func (s *Server) Stop() error {
	s.logger.Println("Shutting down server")
	return s.httpServer.Close()
}
