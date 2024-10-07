package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
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
	tmpl := template.New("base").Funcs(template.FuncMap{
		"formatCodeBlock": FormatCodeBlock,
	})
	contentTmpl, err = tmpl.ParseFiles("templates/base.html", "templates/content.html", "templates/imports.html", "templates/constant.html", "templates/variable.html", "templates/type.html", "templates/function.html", "templates/method.html", "templates/interface.html", "templates/comments.html", "templates/parameter.html")
	if err != nil {
		logger.ErrorLogger.Fatalf("Failed to parse content template: %v", err)
	}
	logger.DebugLogger.Println("Content template parsed successfully")
}

func ContentHandler(w http.ResponseWriter, r *http.Request) {
	logger.DebugLogger.Printf("ContentHandler called with path: %s", r.URL.Path)

	// Construct the file path using forward slashes
	filePath := path.Join("static", "content", r.URL.Path)

	// Convert to OS-specific path
	osBaseDir := filepath.FromSlash(getBasePath(filePath))
	osDirPath := filepath.FromSlash(filePath)
	goOSFilePath := filepath.FromSlash(filePath + ".go")

	_, fiErr := os.Stat(goOSFilePath)
	_, dirErr := os.Stat(osDirPath)
	isFile := !os.IsNotExist(fiErr)
	isDir := !os.IsNotExist(dirErr)

	var nestedDirs, nestedFiles []string
	var err error

	
	if !isFile && !isDir { // no file, no dir ==> error
		logger.ErrorLogger.Printf("could not match path: %s", osDirPath)
		http.Error(w, "Path Not Found", http.StatusNotFound)
		return
	} else if isDir && isFile { // is file, is dir ==> add nested content to render
		// get other paths
		nestedDirs, nestedFiles, err = getNestedContent(osDirPath)
		if err != nil {
			logger.ErrorLogger.Printf("could not get nested content from: %s", osDirPath)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	} else if isDir && !isFile { // no file, is dir ==> render other nav template with nested content
		// get other paths
		nestedDirs, nestedFiles, err = getNestedContent(osDirPath)
		if err != nil {
			logger.ErrorLogger.Printf("could not get nested content from: %s", osDirPath)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		// render directory nav template
	} else if !isDir && isFile { // is file, no dir ==> get all same-level content and directories
		// get other paths
		nestedDirs, nestedFiles, err = getNestedContent(osBaseDir)
		if err != nil {
			logger.ErrorLogger.Printf("could not get nested content from: %s", osBaseDir)
		}
	}

	// Extract the nested path
	pathParts := strings.Split(r.URL.Path, "/")
	nestedPath := strings.Join(reverseSlice(pathParts[:len(pathParts)-1]), " < ")

	parsedFile, err := parser.ParseGoFile(goOSFilePath)
	if err != nil {
		logger.ErrorLogger.Printf("Error parsing file: %v", err)
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	logger.InfoLogger.Printf("File parsed successfully: %s", goOSFilePath)

	// Create a struct to hold both the parsed file and the nested path
	data := struct {
		File       *gofileparser.GFPGoFile
		FullPath string
		Title      string
		NestedDirs []string
		NestedContent []string
	}{
		File:       parsedFile,
		FullPath: nestedPath,
		Title:      parsedFile.Package,
		NestedDirs: nestedDirs,
		NestedContent: nestedFiles,
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
	return s
}

func getBasePath(filePath string) string {
	u, err := url.Parse(filePath)
	if err != nil {
			return ""
	}

	return filepath.Dir(u.Path)
}

func getNestedContent(dirPath string) ([]string, []string, error) {
	dirs := []string{}
	files := []string{}
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return nil, nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			dirs = append(dirs, entry.Name())
		} else {
			files = append(files, entry.Name())
		}
	}

	return dirs, files, nil
}
