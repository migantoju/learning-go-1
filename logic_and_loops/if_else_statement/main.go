package main

import "fmt"

func main() {
	// define an `int` variable
	input := 4

	// create an `if` statement
	if input%2 == 0 {
		fmt.Println(input, " is even")
	} else {
		fmt.Println(input, " is odd")
	}
}
