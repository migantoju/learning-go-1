package main

import (
	"fmt"
	"time"
)

func main() {
	// using pointers

	// declare a ´pointer´ using var statement.
	var count1 *int

	// create a variable using ´new´
	count2 := new(int)

	// create a temporary variable to hold a number
	countTemp := 5

	// using ´&´, create a pointer from existing variable
	count3 := &countTemp

	t := &time.Time{}

	// print each using format strings
	fmt.Printf("Count1: %#v\n", count1)
	fmt.Printf("Count2: %#v\n", count2)
	fmt.Printf("Count3: %#v\n", count3)
	fmt.Printf("Time: %#v\n", t)

	// getting value from pointers
	if count1 != nil {
		fmt.Printf("Count1: %#v\n", *count1)
	}
	if count2 != nil {
		fmt.Printf("Count2: %#v\n", *count2)
	}
	if count3 != nil {
		fmt.Printf("Count3: %#v\n", *count3)
	}
	if t != nil {
		fmt.Printf("Time: %#v\n", *t)
		fmt.Printf("Time: %#v\n", t.String())
	}
}
