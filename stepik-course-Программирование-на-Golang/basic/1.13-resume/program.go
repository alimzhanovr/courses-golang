package main

func SumOfNumbers(number int) int {
	if number < 10 {
		return number
	}
	return number%10 + SumOfNumbers(number/10)
}
