package main

import "fmt"

func main() {
	var amount, total float64

	amount = paintNeeded(4.2, 3.0)
	fmt.Printf("Необходимо %.2f литров.\n", amount)
	total += amount

	amount = paintNeeded(5.2, 3.5)
	fmt.Printf("Необходимо %.2f литров.\n", amount)
	total += amount

	fmt.Printf("Итого нужно %.2f литров.\n", total)
}
func paintNeeded(width float64, height float64) float64 {
	area := width * height
	return area
}
