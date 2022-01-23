package main

func PositiveOrNegative(number int) string {
	if number > 0 {
		return "Число положительное"
	}
	if number < 0 {
		return "Число отрицательное"
	}
	return "Ноль"
}
