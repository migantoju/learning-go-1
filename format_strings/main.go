/* Zero values and format strings */

package main

import (
	"fmt"
	"time"
)

func main() {
	// declare and print integer
	var count int
	fmt.Printf("Count : %#v \n", count)

	// declare and print float
	var discount float64
	fmt.Printf("Discount : %#v \n", discount)

	// declare and print boolean
	var debug bool
	fmt.Printf("Debug : %#v \n", debug)

	// declare and print a string
	var message string
	fmt.Printf("Message : %#v \n", message)

	// declare and print a collection of strings
	var emails []string
	fmt.Printf("Emails : %#v \n", emails)

	// declare and print a struct
	var startTime time.Time
	fmt.Printf("Start : %#v \n", startTime)

}
