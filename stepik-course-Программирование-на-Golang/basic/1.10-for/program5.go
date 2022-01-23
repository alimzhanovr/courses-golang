package main

func multipleCAndnotMultipleD(number, c, d int) bool {
	if (number%c == 0) && (number%d != 0) {
		return true
	}
	return false
}
func numbersLessThan10K(numbers ...int) bool {
	for _, number := range numbers {
		if number > 10e3 {
			return false
		}
	}
	return true
}

func FirstNumberFrom1ToNumberMultCAndNotMultD(number, c, d int) int {
	if !numbersLessThan10K(number, c, d) {
		return -1
	}
	for i := 1; i <= number; i++ {
		if multipleCAndnotMultipleD(i, c, d) {
			return i
		}
	}
	return -1
}
