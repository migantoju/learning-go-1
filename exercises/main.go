/*
The following code doesn't work. The person who wrote it can't fix it,
and they've asked you to help them. Can you get it to work?
*/
package main

import "fmt"

func main() {
	count := 5
	var message string

	if count > 5 {
		message = "Greater than 5"
	} else {
		message = "Not greater than 5"
	}
	fmt.Println(message)
	Activity2()
}

func Activity2() {
	count := 0

	if count < 5 {
		count = 10
		count++
	}
	fmt.Println(count == 11)
}
