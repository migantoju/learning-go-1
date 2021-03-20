package main

import "fmt"

func main() {
	helloList := []string{
		"Hello, world",
		"Hola, Mundo",
		"añsjkdñalskd",
	}

	fmt.Println(len(helloList))
	fmt.Println(helloList[len(helloList)-1])
	fmt.Println(helloList[len(helloList)])
}
