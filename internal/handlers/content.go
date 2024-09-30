package handlers

import (
	"html/template"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/tealwp/gofileparser"
	"github.com/tealwp/sac/internal/logger"
	"github.com/tealwp/sac/internal/parser"
)

var contentTmpl *template.Template

func init() {
	var err error
	contentTmpl, err = template.ParseFiles("templates/base.html", "templates/content.html", "templates/import.html", "templates/constant.html", "templates/variable.html", "templates/type.html", "templates/function.html", "templates/method.html", "templates/interface.html", "templates/comment.html")
	if err != nil {
		logger.ErrorLogger.Fatalf("Failed to parse content template: %v", err)
	}
	logger.DebugLogger.Println("Content template parsed successfully")
}

func ContentHandler(w http.ResponseWriter, r *http.Request) {
	logger.DebugLogger.Printf("ContentHandler called with path: %s", r.URL.Path)

	// Remove the "/science/" prefix
	urlPath := strings.TrimPrefix(r.URL.Path, "/science/")

	// Construct the file path using forward slashes
	filePath := path.Join("static", "content", "science", urlPath+".go")

	// Convert to OS-specific path
	osFilePath := filepath.FromSlash(filePath)

	logger.DebugLogger.Printf("Attempting to read file: %s", osFilePath)

	// Check if the file exists
	if _, err := os.Stat(osFilePath); os.IsNotExist(err) {
		logger.ErrorLogger.Printf("File not found: %s", osFilePath)
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	// Extract the nested path
	pathParts := strings.Split(urlPath, "/")
	nestedPath := strings.Join(reverseSlice(pathParts[:len(pathParts)-1]), " < ")

	parsedFile, err := parser.ParseGoFile(osFilePath)
	if err != nil {
		logger.ErrorLogger.Printf("Error parsing file: %v", err)
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	logger.InfoLogger.Printf("File parsed successfully: %s", osFilePath)

	// Create a struct to hold both the parsed file and the nested path
	data := struct {
		File       *gofileparser.GFP_GoFile
		NestedPath string
		Package    string
	}{
		File:       parsedFile,
		NestedPath: nestedPath,
		Package:    parsedFile.Package, // Assuming the package name is stored in a 'Package' field
	}

	w.Header().Set("Content-Type", "text/html")
	if err := contentTmpl.ExecuteTemplate(w, "base", data); err != nil {
		logger.ErrorLogger.Printf("Error executing template: %v", err)
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	logger.InfoLogger.Printf("Template executed successfully")
}

// Helper function to reverse a slice of strings
func reverseSlice(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	s = append(s, "science")
	return s
}
