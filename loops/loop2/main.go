package main

import "fmt"

func main() {
	// define a variable which is a slice of strings.
	names := []string{"Jim", "Jane", "Joe", "June"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
}
