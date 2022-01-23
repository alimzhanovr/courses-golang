package main

func findNumberOfTens(number int) int {
	//На вход дается натуральное число, не превосходящее 10000.
	if number > 10000 {
		return -1
	}
	return number % 100 / 10
}
