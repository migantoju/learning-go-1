package main

import "fmt"

func main() {
	var x [5]float64
	x[0] = 100
	x[1] = 51
	x[2] = 63
	x[3] = 84
	x[4] = 16

	var total float64 = 0

	for i := 0; i < len(x); i++ {
		total += x[i]
	}

	fmt.Println("Total is: ", total)
	fmt.Println(total / 5)
}
