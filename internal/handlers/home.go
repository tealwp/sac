package handlers

import (
	"html/template"
	"net/http"

	"github.com/tealwp/sac/internal/logger"
)

var homeTmpl *template.Template

func init() {
	var err error
	homeTmpl, err = template.ParseFiles("templates/base.html", "templates/home.html")
	if err != nil {
		logger.ErrorLogger.Fatalf("Failed to parse home template: %v", err)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	logger.DebugLogger.Printf("HomeHandler called with path: %s", r.URL.Path)

	w.Header().Set("Content-Type", "text/html")
	if err := homeTmpl.ExecuteTemplate(w, "base", nil); err != nil {
		logger.ErrorLogger.Printf("Error executing template: %v", err)
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
	}
}
