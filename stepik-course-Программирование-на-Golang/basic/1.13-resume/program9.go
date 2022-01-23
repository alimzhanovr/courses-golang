package main

func digitalRoot(number int) int {
	if number < 10 {
		return number
	}
	return digitalRoot(number%10 + digitalRoot(number/10))
}
