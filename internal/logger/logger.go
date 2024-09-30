package logger

import (
	"log"
	"os"
)

var (
	DebugLogger *log.Logger
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

func init() {
	DebugLogger = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
