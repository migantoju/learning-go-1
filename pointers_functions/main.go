/*
we'll create two functions: one that accepts a number by value,
and another function that accepts a number as a pointer.
*/
package main

import "fmt"

// create a function that takes an int as an argument
func add5Value(count int) {
	// add 5 to the passed number
	count += 5

	// print the updated number
	fmt.Println("add5Value	:", count)
}

// create another function that takes an ´int´ pointer
func add5Point(count *int) {
	// deference the value and add 5 to it
	*count += 5

	// print out the updated value of ´count´ and deference it.
	fmt.Println("add5Point	:", *count)
}

func main() {
	// declare an ´int´ variable
	var count int

	// call the first function with the variable
	add5Value(count)

	// print the current value of the variable
	fmt.Println("add5Value post: ", count)

	// call the second function and pass the pointer.
	add5Point(&count)

	// print the current value of the variable
	fmt.Println("add5Point post: ", count)
}
