package main

import "fmt"

func taxCalculate(amount, taxRate float64) float64 {
	return amount * taxRate / 100
}

func main() {
	cake := 0.99
	milk := 2.75
	butter := 0.87

	var result float64

	cakeTax := taxCalculate(cake, 7.5)
	milkTax := taxCalculate(milk, 1.5)
	butterTax := taxCalculate(butter, 2)

	result = cakeTax + milkTax + butterTax

	fmt.Println(result)
}
