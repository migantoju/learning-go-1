package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	// seed the random number
	rand.Seed(time.Now().UnixNano())

	// generate a random number between 0 - 4
	r := rand.Intn(5) + 1

	stars := strings.Repeat("*", r)

	fmt.Println(stars)
}
