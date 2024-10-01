package server

import (
	"net/http"
	"os"

	"github.com/tealwp/sac/internal/handlers"
	"github.com/tealwp/sac/internal/logger"
	"github.com/tealwp/sac/internal/middleware"
)

func Serve() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	mux := http.NewServeMux()

	// Static file server
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Content handler - This should handle all paths under /science/
	mux.HandleFunc("/science/", handlers.ContentHandler)

	// Home page
	mux.HandleFunc("/", handlers.HomeHandler)

	// Wrap the mux with logging middleware
	loggingMux := middleware.Logging(mux)

	logger.InfoLogger.Printf("Server starting at http://localhost:%s/", port)
	return http.ListenAndServe(":"+port, loggingMux)
}
