package main

import "fmt"

func main() {

	// define a `map` with a string key and string value
	config := map[string]string{
		"debug":    "1",
		"logLevel": "warn",
		"version":  "1.2.1",
	}

	// Use range to get key and value for an array element and
	// assign the to variables
	for key, value := range config {
		fmt.Println(key, "=", value)
	}

}
