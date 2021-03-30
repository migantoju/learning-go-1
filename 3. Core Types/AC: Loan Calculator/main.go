package main

func taxCalculator(amount, taxRate float64) float64 {
	return amount * taxRate / 100
}

func getInterest(amount, interestRate, loanTerm float64) float64 {
	return amount * interestRate * loanTerm
}

func getLoanTerm(period int) bool {
	if !(period%12 == 0) {
		panic("No se puede.")
	}
	return true
}

func getMonthlyPayment(amount, rate float64) float64 {
	return amount * rate / 100
}

func main() {

	goodScore := 450
	goodRate := 15
	badRate := 20

	loanTerm := 24

}
