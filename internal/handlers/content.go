package handlers

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/tealwp/sac/internal/logger"
	"github.com/tealwp/sac/internal/parser"
)

var templates *template.Template

func init() {
	var err error
	templates, err = template.ParseGlob("templates/*.html")
	if err != nil {
		logger.ErrorLogger.Fatalf("Failed to parse templates: %v", err)
	}
}

func ContentHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/science/")
	filePath := filepath.Join("static", "content", "science", path)

	fileInfo, err := os.Stat(filePath)
	if err == nil && fileInfo.IsDir() {
		filePath = filepath.Join(filePath, "index.go")
	} else {
		filePath += ".go"
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		logger.ErrorLogger.Printf("File not found: %s", filePath)
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	parsedFile, err := parser.ParseGoFile(filePath)
	if err != nil {
		logger.ErrorLogger.Printf("Error parsing file: %v", err)
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Log the parsed data
	logger.InfoLogger.Printf("Parsed file data: %+v", parsedFile)

	w.Header().Set("Content-Type", "text/html")
	if err := templates.ExecuteTemplate(w, "content.html", parsedFile); err != nil {
		logger.ErrorLogger.Printf("Error executing template: %v", err)
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Log after template execution
	logger.InfoLogger.Println("Template executed successfully")
}
