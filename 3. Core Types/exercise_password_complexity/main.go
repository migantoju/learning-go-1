package main

import (
	"fmt"
	"unicode"
)

// create a function that takes a string argument and returns a bool
func passwordChecker(password string) bool {
	// convert the password string into `rune`, which is a safe type
	// for multi-byte (UTF-8) chars.
	passwordRune := []rune(password)

	// count the number of the password
	if len(passwordRune) < 8 {
		return false
	}

	// define some bool variables
	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSymbol := false

	// loop over the multi-type charc one at time.
	for _, v := range passwordRune {
		if unicode.IsUpper(v) {
			hasUpper = true
		}
		if unicode.IsLower(v) {
			hasLower = true
		}
		if unicode.IsNumber(v) {
			hasNumber = true
		}
		if unicode.IsPunct(v) || unicode.IsSymbol(v) {
			hasSymbol = true
		}
	}

	return hasUpper && hasLower && hasNumber && hasSymbol
}

func main() {
	if passwordChecker("This&@a87") {
		fmt.Println("Password Good")
	} else {
		fmt.Println("Password Bad.")
	}
}
