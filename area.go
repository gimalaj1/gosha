package main

import "fmt"

func main() {
	var amount, total float64

	amount, err := paintNeeded(4.2, 3.0)
	fmt.Println(err)
	fmt.Printf("Необходимо %.2f литров.\n", amount)
	total += amount

	amount, err = paintNeeded(5.2, 3.5)
	fmt.Println(err)
	fmt.Printf("Необходимо %.2f литров.\n", amount)
	total += amount

	fmt.Printf("Итого нужно %.2f литров.\n", total)
}

func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("Неверно указано значение ширины: %.2f", width)
	}

	if height < 0 {
		return 0, fmt.Errorf("Неверно указано значение высоты: %.2f", height)
	}

	area := width * height
	return area, nil
}
