package main

import (
	"fmt"
)

// Declare package-level variables (outside any function)
var globalCounter int           // zero value: 0
var appName string = "GoMaster" // explicitly initialized
var debugMode bool              // zero value: false

// Variable block - cleaner for multiple declarations
var (
	version   = "1.0.0"
	buildTime string // will be empty (zero value) unless set
	author    = "Anonymous"
)

// Constants - compile-time bound
const (
	AppName    = "100 Days of Go"
	MaxRetries = 3
	Debug      = false
)

func main() {
	// --- 1. VAR DECLARATIONS ---
	var message string // zero value: ""
	message = "Good morning!"
	fmt.Println("Message:", message)
}
