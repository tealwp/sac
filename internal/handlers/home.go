package handlers

import (
	"html/template"
	"net/http"

	"github.com/tealwp/sac/internal/logger"
)

var homeTmpl *template.Template

func init() {
	var err error
	homeTmpl, err = template.ParseFiles("templates/home.html")
	if err != nil {
		logger.ErrorLogger.Fatalf("Failed to parse template: %v", err)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	if err := homeTmpl.Execute(w, nil); err != nil {
		logger.ErrorLogger.Printf("Error executing template: %v", err)
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
	}
}
