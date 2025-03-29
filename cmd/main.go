package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func main() {
	logger := log.New(os.Stdout, "SERVER:", log.LstdFlags)
	// Get the directory where main.go is located
	exePath, err := os.Executable()
	if err != nil {
		logger.Fatal("Could not determine executable path:", err)
	}
	baseDir := filepath.Dir(exePath)
	if filepath.Base(exePath) == "main" {
		baseDir = filepath.Join(baseDir, "../../")
	}

	// Create service
	morseService := service.NewMorseService()

	// Creating handlers
	handler := handlers.NewMorseHandlers(morseService, baseDir)

	// Create server
	mux := server.New(logger)
	handler.RegisterHandlers(mux)

	if err := mux.Start(); err != nil {
		logger.Fatal(err)
	}
}
