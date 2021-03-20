// In this exercise, we're going to print out a message that tells us whether the day
// someone was born was a weekday or the weekend. We only need two cases as each case
// can support checking multiple values:
package main

import (
	"fmt"
	"time"
)

func main() {
	dayBorn := time.Sunday

	switch dayBorn {
	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		fmt.Println("Born on a weekday")
	case time.Saturday, time.Sunday:
		fmt.Println("Born on a weekend")
	default:
		fmt.Println("Error, day born not valid.")
	}
}
