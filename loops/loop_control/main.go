package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// create an empty for loop
	for {
		// use Intn to pick a random number between 0 and 8
		r := rand.Intn(8)

		// if the random number is divisible by 3, print "Skip"
		// and skip the rest of the loop using `continue`
		// if the random number is divisible by 2, the print "Stop"
		// and stop the loop using break.
		if r%3 == 0 {
			fmt.Println("Skip")
			continue
		} else if r%2 == 0 {
			fmt.Println("Stop")
			break
		}
		// if the number is neither of those things, then print the number
		fmt.Println(r)
	}
}
