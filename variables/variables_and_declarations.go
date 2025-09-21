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

	// --- 2. SHORT VARIABLE DECLARATION (:=) - inside functions only
	name := "Alice" // type inferred as string
	age := 30       // type inferred as int
	isStudent := true
	fmt.Printf("%s is %d years old. Student? %t\n", name, age, isStudent)

	// --- 3. MULTIPLE ASSIGNMENTS & SWAP ---
	a, b := 10, 20
	fmt.Println("Before swap: a =", a, ", b =", b)
	a, b = b, a // Go magic: no temp variable needed
	fmt.Println("After swap: a =", a, "b =", b)

	// --- 4. BLANK IDENTIFIER (ignore values) ---
	_, remainder := divideWithRemainder(17, 5)
	fmt.Println("Remainder:", remainder)

	// --- 5. TYPE CONVERSIONS (explicit in Go) ---
	var floatValue float64 = 3.14
	var intValue int = int(floatValue) // truncates, does not round
	fmt.Printf("Float: %f -> Int: %d\n", floatValue, intValue)

	// --- 6. PRINT ALL PACKAGE-LEVEL VARS ---
	fmt.Println("\n=== Package-Level State ===")
	fmt.Println("Global Counter:", globalCounter)
	fmt.Println("App Name:", appName)
	fmt.Println("Debug Mode:", debugMode)
	fmt.Println("Version:", version)
	fmt.Println("Build Time:", buildTime)
	fmt.Println("Author:", author)
}

// Helper function: returns quotient and remainder
func divideWithRemainder(dividend, divisor int) (quotient, remainder int) {
	quotient = dividend / divisor
	remainder = dividend % divisor
	return // naked return (uses named return values)
}
