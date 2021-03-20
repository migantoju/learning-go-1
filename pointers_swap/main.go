/*
Unfinished code to complete with pointers exercise.
*/
package main

import "fmt"

func main() {
	a, b := 5, 10

	// call swap here
	swap(&a, &b)

	// print to the console the boolean values both may be true.
	fmt.Println(a == 10, b == 5)
}

func swap(a, b *int) {
	// swap the values here
	*a, *b = *b, *a
}
