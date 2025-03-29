package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func main() {
	// Get the directory where main.go is located
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal("Could not determine executable path:", err)
	}
	baseDir := filepath.Dir(exePath)

	// For development (when using go run)
	if filepath.Base(exePath) == "main" {
		baseDir = filepath.Join(baseDir, "../../") //
	}

	// Create service
	morseService := service.NewMorseService()

	// Create server
	mux := http.NewServeMux()
	handlers.RegisterAll(mux, morseService)

	log.Println("Starting server on :8080")
	log.Println("Serving from directory:", baseDir)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
