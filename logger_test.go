package main

import (
	"testing"
	"os"
	"io"
)

func TestInitLogger(t *testing.T) {
	DebugHandle = os.Stdout
	TraceHandle = os.Stderr
	f, _ := os.Create("gochat_info.log")
	InfoHandle = io.MultiWriter(f)
	errFile, _ := os.Create("gochat_err.log")
	ErrorHandle = io.MultiWriter(errFile)
	InitLogger()
	DebugLogger.Println("debug log")
	TraceLogger.Println("I have something standard to say")
	InfoLogger.Println("Special Information")
	WarningLogger.Println("There is something you need to know about")
	ErrorLogger.Println("Something has failed")
}
