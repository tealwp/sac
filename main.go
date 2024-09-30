package main

import (
	"github.com/tealwp/sac/cmd/server"
	"github.com/tealwp/sac/internal/logger"
)

func main() {
	if err := server.Serve(); err != nil {
		logger.ErrorLogger.Fatalf("Server error: %v", err)
	}
}
