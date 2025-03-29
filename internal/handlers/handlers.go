package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// MorseService interface to use from Service package
type MorseService interface {
	AutoConvert(input string) (string, error)
}

// MorseHandlers struct to use it in Handlers
type MorseHandlers struct {
	service MorseService
	baseDir string
}

// NewMorseHandlers create new handlers structure
func NewMorseHandlers(ms MorseService) *MorseHandlers {
	return &MorseHandlers{service: ms}
}

// RegisterAll registring all handlers
func RegisterAll(mux *http.ServeMux, ms MorseService) {
	h := NewMorseHandlers(ms)

	mux.HandleFunc("/", h.serveStaticPage)
	mux.HandleFunc("/upload", h.handleUpload)
}

func (h *MorseHandlers) serveStaticPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Construct absolute path to index.html
	htmlPath := filepath.Join(h.baseDir, "index.html")

	// Check if file exists
	if _, err := os.Stat(htmlPath); os.IsNotExist(err) {
		http.Error(w, "index.html not found at: "+htmlPath, http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, htmlPath)
}

func (h *MorseHandlers) handleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "No file uploaded", http.StatusBadRequest)
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}

	result, err := h.service.AutoConvert(string(content))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(result))

	// Write to file
	err = os.WriteFile("output.txt", []byte(result), 0o644)
	if err != nil {
		log.Printf("Failed to write file: %v", err)
	}
}
