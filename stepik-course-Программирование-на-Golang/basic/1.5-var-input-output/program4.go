package main

func FindLastNumber(number int) int {
	//На вход дается натуральное число N, не превосходящее 10000.
	if number > 10000 {
		return -1
	}
	return number % 10
}
