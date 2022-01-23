package main

func firstNumber(number int) int {
	if 0 > number || number > 10000 {
		return -1
	}
	for {
		if number < 10 {
			return number
		}
		number /= 10
	}
}
